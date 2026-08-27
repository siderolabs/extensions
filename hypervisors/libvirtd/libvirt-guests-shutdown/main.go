// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

const (
	connectionURI      = "qemu+unix:///system?socket=/run/libvirt/virtqemud-sock"
	managedSaveTimeout = 300 * time.Second
	virshPath          = "/usr/local/bin/virsh"
)

func main() {
	log.SetFlags(0)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := shutdownGuests(ctx); err != nil {
		log.Fatal(err)
	}
}

type virshCommand func(context.Context, ...string) (string, error)

func shutdownGuests(ctx context.Context) error {
	managedSaveCtx, cancel := context.WithTimeout(ctx, managedSaveTimeout)
	defer cancel()

	return shutdownGuestsWith(ctx, managedSaveCtx, runVirsh)
}

func shutdownGuestsWith(ctx, managedSaveCtx context.Context, run virshCommand) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	domains, err := activeDomains(managedSaveCtx, run)
	if err != nil {
		log.Printf("failed to enumerate domains for managed save: %v", err)
	}

	for _, domain := range domains {
		if managedSaveCtx.Err() != nil {
			break
		}

		log.Printf("saving domain %s", domain)

		if _, err = run(managedSaveCtx, "managedsave", "--running", domain); err != nil {
			log.Printf("managed save failed for domain %s, will destroy it: %v", domain, err)
		}
	}

	if err = ctx.Err(); err != nil {
		return err
	}

	remaining, err := activeDomains(ctx, run)
	if err != nil {
		return fmt.Errorf("failed to enumerate domains before forced shutdown: %w", err)
	}

	var errs []error

	for _, domain := range remaining {
		log.Printf("destroying domain %s", domain)

		if _, destroyErr := run(ctx, "destroy", domain); destroyErr != nil {
			errs = append(errs, fmt.Errorf("failed to stop domain %s: %w", domain, destroyErr))
		}
	}

	remaining, err = activeDomains(ctx, run)
	if err != nil {
		errs = append(errs, fmt.Errorf("failed to verify domain shutdown: %w", err))
	} else if len(remaining) > 0 {
		errs = append(errs, fmt.Errorf("domains still active after shutdown: %s", strings.Join(remaining, ", ")))
	}

	return errors.Join(errs...)
}

func activeDomains(ctx context.Context, run virshCommand) ([]string, error) {
	output, err := run(ctx, "list", "--uuid")
	if err != nil {
		return nil, err
	}

	var domains []string

	scanner := bufio.NewScanner(strings.NewReader(output))
	for scanner.Scan() {
		if domain := strings.TrimSpace(scanner.Text()); domain != "" {
			domains = append(domains, domain)
		}
	}

	if err = scanner.Err(); err != nil {
		return nil, fmt.Errorf("failed to read active domains: %w", err)
	}

	return domains, nil
}

func runVirsh(ctx context.Context, args ...string) (string, error) {
	commandArgs := append([]string{"--connect", connectionURI}, args...)
	output, err := exec.CommandContext(ctx, virshPath, commandArgs...).CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("virsh %s failed: %w: %s", strings.Join(args, " "), err, strings.TrimSpace(string(output)))
	}

	return string(output), nil
}
