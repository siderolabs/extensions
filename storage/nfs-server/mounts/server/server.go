// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package server supervises the host processes that comprise the NFS server.
package server

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	exportsPath           = "/etc/exports"
	defaultReloadInterval = time.Second
	defaultCommandTimeout = 5 * time.Second
	defaultCleanupTimeout = 8 * time.Second
)

// ErrChildExited indicates that a supervised daemon exited unexpectedly.
var ErrChildExited = errors.New("NFS server child process exited")

// Process is a supervised operating-system process.
type Process interface {
	PID() int
	Done() <-chan struct{}
	Err() error
	Stop(context.Context) error
}

// Runtime provides the host operations required by Run.
type Runtime interface {
	EnsureMounts() error
	Start(name string, args ...string) (Process, error)
	WaitReady(context.Context, string, Process) error
	Run(context.Context, string, ...string) error
	ReadExports() ([]byte, error)
	IsMountPoint(string) (bool, error)
}

// Config controls server supervision.
type Config struct {
	ReloadInterval time.Duration
	CommandTimeout time.Duration
	CleanupTimeout time.Duration
}

// DefaultConfig returns production supervision timeouts. Cleanup deliberately completes before
// Talos's ten-second host-process termination grace period.
func DefaultConfig() Config {
	return Config{
		ReloadInterval: defaultReloadInterval,
		CommandTimeout: defaultCommandTimeout,
		CleanupTimeout: defaultCleanupTimeout,
	}
}

func normalizeConfig(config Config) Config {
	defaults := DefaultConfig()
	if config.ReloadInterval <= 0 {
		config.ReloadInterval = defaults.ReloadInterval
	}
	if config.CommandTimeout <= 0 {
		config.CommandTimeout = defaults.CommandTimeout
	}
	if config.CleanupTimeout <= 0 {
		config.CleanupTimeout = defaults.CleanupTimeout
	}

	return config
}

// Run starts the NFS server components in dependency order and supervises them.
func Run(ctx context.Context, runtime Runtime, config Config) error {
	config = normalizeConfig(config)

	if err := runtime.EnsureMounts(); err != nil {
		return fmt.Errorf("prepare NFS filesystems: %w", err)
	}

	if err := waitForBackingMounts(ctx, runtime, config.ReloadInterval); err != nil {
		if ctx.Err() != nil {
			return nil
		}

		return fmt.Errorf("wait for NFS backing mounts: %w", err)
	}

	mountd, err := runtime.Start("rpc.mountd", "--foreground", "--port", "20048")
	if err != nil {
		return fmt.Errorf("start rpc.mountd: %w", err)
	}

	var nfsdcld Process
	cleanup := func() error {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), config.CleanupTimeout)
		defer cancel()

		operationTimeout := min(config.CommandTimeout, config.CleanupTimeout/4)
		var cleanupErrors []error
		if err := runCommand(cleanupCtx, runtime, operationTimeout, "rpc.nfsd", "0"); err != nil {
			cleanupErrors = append(cleanupErrors, fmt.Errorf("stop kernel NFS server: %w", err))
		}
		if err := runCommand(cleanupCtx, runtime, operationTimeout, "exportfs", "-a", "-u"); err != nil {
			cleanupErrors = append(cleanupErrors, fmt.Errorf("unload NFS exports: %w", err))
		}

		var waitGroup sync.WaitGroup
		var errorsMu sync.Mutex
		stop := func(name string, process Process) {
			defer waitGroup.Done()
			if err := process.Stop(cleanupCtx); err != nil {
				errorsMu.Lock()
				cleanupErrors = append(cleanupErrors, fmt.Errorf("stop %s: %w", name, err))
				errorsMu.Unlock()
			}
		}

		waitGroup.Add(1)
		go stop("rpc.mountd", mountd)
		if nfsdcld != nil {
			waitGroup.Add(1)
			go stop("nfsdcld", nfsdcld)
		}
		waitGroup.Wait()

		return errors.Join(cleanupErrors...)
	}

	fail := func(runErr error) error {
		return errors.Join(runErr, cleanup())
	}

	if err = waitReady(ctx, runtime, config.CommandTimeout, "rpc.mountd", mountd); err != nil {
		return fail(fmt.Errorf("wait for rpc.mountd: %w", err))
	}

	nfsdcld, err = runtime.Start("nfsdcld", "-F")
	if err != nil {
		return fail(fmt.Errorf("start nfsdcld: %w", err))
	}

	if err = waitReady(ctx, runtime, config.CommandTimeout, "nfsdcld", nfsdcld); err != nil {
		return fail(fmt.Errorf("wait for nfsdcld: %w", err))
	}

	exports, err := runtime.ReadExports()
	if err != nil {
		return fail(fmt.Errorf("read %s: %w", exportsPath, err))
	}

	ready, err := backingMountsReady(runtime, exports)
	if err != nil {
		return fail(fmt.Errorf("check NFS backing mounts: %w", err))
	}
	if !ready {
		return fail(errors.New("NFS backing mount disappeared or configuration changed during startup"))
	}

	if err = runCommand(ctx, runtime, config.CommandTimeout, "exportfs", "-r", "-v"); err != nil {
		return fail(fmt.Errorf("load NFS exports: %w", err))
	}

	if err = runCommand(ctx, runtime, config.CommandTimeout, "rpc.nfsd", "--port", "2049", "8"); err != nil {
		return fail(fmt.Errorf("start kernel NFS server: %w", err))
	}

	if err = waitReady(ctx, runtime, config.CommandTimeout, "nfsdcld-attached", nfsdcld); err != nil {
		return fail(err)
	}

	ticker := time.NewTicker(config.ReloadInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return cleanup()
		case <-mountd.Done():
			return fail(childExitError("rpc.mountd", mountd.Err()))
		case <-nfsdcld.Done():
			return fail(childExitError("nfsdcld", nfsdcld.Err()))
		case <-ticker.C:
			currentExports, readErr := runtime.ReadExports()
			if readErr != nil || bytes.Equal(exports, currentExports) {
				continue
			}

			ready, mountErr := backingMountsReady(runtime, currentExports)
			if mountErr != nil || !ready {
				continue
			}

			if reloadErr := runCommand(ctx, runtime, config.CommandTimeout, "exportfs", "-r", "-v"); reloadErr != nil {
				continue
			}

			exports = currentExports
		}
	}
}

func waitForBackingMounts(ctx context.Context, runtime Runtime, interval time.Duration) error {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	waiting := false
	for {
		if err := ctx.Err(); err != nil {
			return err
		}

		exports, err := runtime.ReadExports()
		if err != nil {
			return fmt.Errorf("read %s: %w", exportsPath, err)
		}

		ready, err := backingMountsReady(runtime, exports)
		if err != nil {
			return err
		}
		if ready {
			if waiting {
				log.Print("NFS export mounts are ready.")
			}
			return nil
		}
		if !waiting {
			log.Printf("Waiting for NFS export mounts. Check the mountpoint options in %s and the backing volume configuration.", exportsPath)
			waiting = true
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
		}
	}
}

// backingMountsReady only gates exports that explicitly declare the NFS mountpoint
// option. A plain directory cannot be distinguished from an unmounted volume.
func backingMountsReady(runtime Runtime, exports []byte) (bool, error) {
	contents := strings.NewReplacer("\\\r\n", "", "\\\n", "").Replace(string(exports))
	for line := range strings.SplitSeq(contents, "\n") {
		fields, err := exportsFields(line)
		if err != nil {
			return false, err
		}
		if len(fields) < 2 {
			continue
		}
		for _, client := range fields[1:] {
			var options string
			if defaults, ok := strings.CutPrefix(client, "-"); ok {
				options = defaults
			} else {
				start := strings.IndexByte(client, '(')
				end := strings.LastIndexByte(client, ')')
				if start < 0 || end <= start {
					continue
				}
				options = client[start+1 : end]
			}
			for option := range strings.SplitSeq(options, ",") {
				name, path, _ := strings.Cut(option, "=")
				if name != "mountpoint" && name != "mp" {
					continue
				}
				if path == "" {
					path = fields[0]
				}
				if !filepath.IsAbs(path) {
					return false, fmt.Errorf("invalid NFS mountpoint %q", path)
				}
				mounted, err := runtime.IsMountPoint(filepath.Clean(path))
				if err != nil {
					return false, fmt.Errorf("check NFS mountpoint %q: %w", path, err)
				}
				if !mounted {
					return false, nil
				}
			}
		}
	}

	return true, nil
}

// exportsFields preserves quoted paths and decodes exports(5) octal escapes.
// Parsing failures stop startup rather than silently bypassing a mount requirement.
func exportsFields(line string) ([]string, error) {
	var fields []string
	var field strings.Builder
	quoted := false
	flush := func() {
		if field.Len() > 0 {
			fields = append(fields, field.String())
			field.Reset()
		}
	}

	for i := 0; i < len(line); i++ {
		switch ch := line[i]; {
		case ch == '"':
			quoted = !quoted
		case ch == '\\':
			if i+1 == len(line) {
				return nil, errors.New("incomplete escape in NFS exports")
			}
			if i+3 < len(line) {
				if value, err := strconv.ParseUint(line[i+1:i+4], 8, 8); err == nil {
					field.WriteByte(byte(value))
					i += 3
					continue
				}
			}
			i++
			field.WriteByte(line[i])
		case !quoted && ch == '#':
			flush()
			return fields, nil
		case !quoted && (ch == ' ' || ch == '\t' || ch == '\r'):
			flush()
		default:
			field.WriteByte(ch)
		}
	}
	if quoted {
		return nil, errors.New("unterminated quoted path in NFS exports")
	}
	flush()

	return fields, nil
}

func runCommand(ctx context.Context, runtime Runtime, timeout time.Duration, name string, args ...string) error {
	commandCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return runtime.Run(commandCtx, name, args...)
}

func waitReady(ctx context.Context, runtime Runtime, timeout time.Duration, name string, process Process) error {
	readyCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return runtime.WaitReady(readyCtx, name, process)
}

func childExitError(name string, err error) error {
	if err == nil {
		err = errors.New("exited without an error")
	}

	return fmt.Errorf("%w: %s: %w", ErrChildExited, name, err)
}
