// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"golang.org/x/sys/unix"
)

const (
	zerotierPath    = "/var/lib/zerotier-one"
	identityPath    = "/var/lib/zerotier-one/identity.secret"
	planetPath      = "/var/lib/zerotier-one/planet"
	identityPubPath = "/var/lib/zerotier-one/identity.public"
	zerotierBinPath = "/usr/local/bin/zerotier-one"
)

func main() {
	log.Printf("zerotier-wrapper: initializing...")

	// Ensure the ZeroTier state directory exists.
	if err := os.MkdirAll(zerotierPath, 0o755); err != nil {
		log.Fatalf("failed to create state directory: %v", err)
	}

	// Ensure identity configuration.
	identitySource, err := ensureIdentity()
	if err != nil {
		log.Fatalf("identity configuration failed: %v", err)
	}
	log.Printf("identity configured (source: %s)", identitySource)

	// If ZEROTIER_PLANET env var is set, set the planet file.
	if planet := os.Getenv("ZEROTIER_PLANET"); planet != "" {
		planet, err := base64.StdEncoding.DecodeString(planet)
		if err != nil {
			log.Fatalf("failed to decode base64 planet from environment: %v", err)
		}
		if err := os.WriteFile(planetPath, planet, 0o644); err != nil {
			log.Fatalf("failed to write planet file: %v", err)
		}
		log.Printf("custom planet file loaded")
	}

	// If ZEROTIER_NETWORK env var is set, join the network.
	if network := os.Getenv("ZEROTIER_NETWORK"); network != "" {
		for _, network := range strings.Split(network, ",") {
			log.Printf("joining network %s", network)
			if err := joinNetwork(network); err != nil {
				log.Fatalf("failed to join network: %v", err)
			}
			log.Printf("joined network %s", network)
		}
	}

	// Start zerotier-one process.
	if err := unix.Exec(zerotierBinPath, []string{zerotierBinPath, "-U", zerotierPath}, os.Environ()); err != nil {
		log.Fatalf("error executing zerotier-one: %v", err)
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
	if err := os.WriteFile(identityPath, []byte(identity), 0o600); err != nil {
		return fmt.Errorf("failed to write secret identity: %w", err)
	}
	log.Printf("wrote secret identity to %s", identityPath)

	// Write the public identity file with only the first 3 parts.
	public := strings.Join(parts[:3], ":")
	if err := os.WriteFile(identityPubPath, []byte(public), 0o644); err != nil {
		return fmt.Errorf("failed to write public identity: %w", err)
	}
	log.Printf("wrote public identity to %s", identityPubPath)

	return nil
}

// joinNetwork creates a config file for the relevant network if it doesn't already exist.
// This is typically done while the service is running via `zerotier-one -q join <network>`,
// however this just creates an empty file with the network name, so we do that instead.
func joinNetwork(network string) error {
	networkConfDir := filepath.Join(zerotierPath, "networks.d")
	if err := os.MkdirAll(networkConfDir, 0o755); err != nil {
		return fmt.Errorf("failed to create networks.d directory: %w", err)
	}
	networkConfFile := filepath.Join(networkConfDir, network+".conf")

	file, err := os.OpenFile(networkConfFile, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
	if err != nil {
		if os.IsExist(err) {
			log.Printf("network configuration file %s already exists", networkConfFile)
			return nil
		}
		return fmt.Errorf("failed to create network conf file: %w", err)
	}
	defer file.Close()

	log.Printf("created network configuration file %s", networkConfFile)
	return nil
}
