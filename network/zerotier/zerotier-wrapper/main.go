package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"time"

	"golang.org/x/sys/unix"
)

const (
	zerotierPath    = "/var/lib/zerotier-one"
	identityPath    = "/var/lib/zerotier-one/identity.secret"
	identityPubPath = "/var/lib/zerotier-one/identity.public"
	pidFile         = "/var/lib/zerotier-one/zerotier-one.pid"
	zerotierBinPath = "/usr/local/bin/zerotier-one"
)

func main() {
	log.Printf("zerotier-wrapper: initializing...")

	// Ensure the ZeroTier state directory exists.
	if err := os.MkdirAll(zerotierPath, 0755); err != nil {
		log.Fatalf("failed to create state directory: %v", err)
	}

	// Ensure identity configuration.
	identitySource, err := ensureIdentity()
	if err != nil {
		log.Fatalf("identity configuration failed: %v", err)
	}
	log.Printf("identity configured (source: %s)", identitySource)

	// Cleanup any existing zerotier-one process.
	if err := cleanupProcess(); err != nil {
		log.Fatalf("process cleanup failed: %v", err)
	}

	// If ZEROTIER_NETWORK env var is set, join network using zerotier-one -q.
	if network := os.Getenv("ZEROTIER_NETWORK"); network != "" {
		log.Printf("will join network %s after startup", network)
		go func() {
			time.Sleep(2 * time.Second)
			if err := joinNetwork(network); err != nil {
				log.Printf("failed to join network: %v", err)
			} else {
				log.Printf("joined network %s", network)
			}
		}()
	}

	// Start zerotier-one process.
	cmd := exec.Command(zerotierBinPath, "-U", zerotierPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		log.Fatalf("error starting zerotier-one: %v", err)
	}

	// Write the PID file.
	pidStr := strconv.Itoa(cmd.Process.Pid)
	if err := os.WriteFile(pidFile, []byte(pidStr), 0644); err != nil {
		log.Printf("failed to write PID file: %v", err)
	}

	// Forward termination signals to the zerotier-one process.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, unix.SIGINT, unix.SIGTERM)
	sig := <-ch
	log.Printf("received signal %v, forwarding to zerotier-one process", sig)
	if err := cmd.Process.Signal(sig); err != nil {
		log.Fatalf("error sending signal to zerotier-one: %v", err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatalf("zerotier-one exited with error: %v", err)
	}
}

// ensureIdentity checks for an existing identity file, validates it if found,
// or else uses the identity from the ZEROTIER_IDENTITY_SECRET environment variable (after validation)
// or generates a new one using "zerotier-one -i generate".
func ensureIdentity() (string, error) {
	// If the identity file exists, validate its contents.
	if _, err := os.Stat(identityPath); err == nil {
		data, err := os.ReadFile(identityPath)
		if err != nil {
			return "", fmt.Errorf("failed to read existing identity: %w", err)
		}
		identity := strings.TrimSpace(string(data))
		log.Printf("found existing identity at %s, validating...", identityPath)
		if err := validateIdentity(identity); err != nil {
			return "", fmt.Errorf("existing identity failed validation: %w", err)
		}
		log.Printf("existing identity validated")
		return "existing", nil
	} else if !errors.Is(err, os.ErrNotExist) {
		return "", fmt.Errorf("failed to stat identity file: %w", err)
	}

	// Check for identity in environment.
	if identity := os.Getenv("ZEROTIER_IDENTITY_SECRET"); identity != "" {
		log.Printf("found identity in ZEROTIER_IDENTITY_SECRET environment variable, validating...")
		if err := validateIdentity(identity); err != nil {
			return "", fmt.Errorf("environment identity invalid: %w", err)
		}
		log.Printf("environment identity validated")
		if err := writeIdentity(identity); err != nil {
			return "", fmt.Errorf("failed to write identity from environment: %w", err)
		}
		return "environment", nil
	}

	// Generate a new identity using "zerotier-one -i generate".
	log.Printf("generating new identity using zerotier-one -i generate")
	cmd := exec.Command(zerotierBinPath, "-i", "generate")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to generate identity: %w", err)
	}
	identity := strings.TrimSpace(out.String())
	if err := validateIdentity(identity); err != nil {
		return "", fmt.Errorf("generated identity failed validation: %w", err)
	}
	if err := writeIdentity(identity); err != nil {
		return "", fmt.Errorf("failed to write generated identity: %w", err)
	}
	return "generated", nil
}

// validateIdentity runs "zerotier-one -i validate <identity>" to ensure the identity is valid.
func validateIdentity(identity string) error {
	cmd := exec.Command(zerotierBinPath, "-i", "validate", identity)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("identity validation failed: %w", err)
	}
	return nil
}

// writeIdentity writes the complete identity string (all four parts) to identity.secret,
// while writing only the first three parts (separated by ':') to identity.public.
func writeIdentity(identity string) error {
	parts := strings.Split(identity, ":")
	if len(parts) != 4 {
		return fmt.Errorf("invalid identity format: expected 4 parts, got %d", len(parts))
	}

	// Write the secret identity file with the full identity.
	if err := os.WriteFile(identityPath, []byte(identity), 0600); err != nil {
		return fmt.Errorf("failed to write secret identity: %w", err)
	}
	log.Printf("wrote secret identity to %s", identityPath)

	// Write the public identity file with only the first 3 parts.
	public := strings.Join(parts[:3], ":")
	if err := os.WriteFile(identityPubPath, []byte(public), 0644); err != nil {
		return fmt.Errorf("failed to write public identity: %w", err)
	}
	log.Printf("wrote public identity to %s", identityPubPath)

	return nil
}

// cleanupProcess checks for an existing PID file; if found, it kills the process and removes the file.
func cleanupProcess() error {
	if _, err := os.Stat(pidFile); err == nil {
		pid, err := getProcessId()
		if err != nil {
			return fmt.Errorf("error reading pid file: %w", err)
		}
		if err := killProcess(pid); err != nil {
			return fmt.Errorf("error killing process: %w", err)
		}
		if err := os.Remove(pidFile); err != nil {
			return fmt.Errorf("error removing pid file: %w", err)
		}
		log.Printf("cleaned up existing process (PID %d)", pid)
	} else if !errors.Is(err, os.ErrNotExist) {
		return fmt.Errorf("failed to stat pid file: %w", err)
	} else {
		log.Printf("no PID file found, no existing process to clean up")
	}
	return nil
}

func getProcessId() (int, error) {
	pidData, err := os.ReadFile(pidFile)
	if err != nil {
		return 0, err
	}
	pidData = bytes.TrimRight(pidData, "\n")
	pid, err := strconv.Atoi(string(pidData))
	if err != nil {
		return 0, err
	}
	return pid, nil
}

func killProcess(pid int) error {
	p, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	if err := p.Kill(); err != nil && !errors.Is(err, os.ErrProcessDone) {
		return err
	}
	return nil
}

// joinNetwork uses "zerotier-one -q join <network>" to join the specified network.
func joinNetwork(network string) error {
	log.Printf("attempting to join network %s", network)
	cmd := exec.Command(zerotierBinPath, "-q", "join", network)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("join network failed: %w", err)
	}
	return nil
}
