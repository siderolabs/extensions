// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package server

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"time"
)

const (
	// DefaultStateDir holds the drain request.
	DefaultStateDir = "/run/mergerfs"

	// DefaultDrainMaxAge is how long a drain request stays in force. Talos
	// stops the services within seconds of the pre-shutdown hooks; a request
	// still around after this belongs to a shutdown that was aborted, and the
	// pools come back on their own.
	DefaultDrainMaxAge = 5 * time.Minute

	drainFile = "drain"
)

// Drain is the drain request file the pre-shutdown hook leaves for the
// supervisor. Its modification time is the time of the request.
type Drain struct {
	Dir    string
	MaxAge time.Duration
}

// Request writes the drain request.
func (d Drain) Request() error {
	if err := os.MkdirAll(d.Dir, 0o755); err != nil {
		return err
	}

	path := filepath.Join(d.Dir, drainFile)

	if err := os.WriteFile(path, nil, 0o644); err != nil {
		return err
	}

	now := time.Now()

	return os.Chtimes(path, now, now)
}

// Active reports whether a drain request is in force, removing one that has expired.
func (d Drain) Active(now time.Time) (bool, error) {
	path := filepath.Join(d.Dir, drainFile)

	info, err := os.Stat(path)
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	if now.Sub(info.ModTime()) <= d.MaxAge {
		return true, nil
	}

	if err = os.Remove(path); err != nil && !errors.Is(err, fs.ErrNotExist) {
		return false, err
	}

	return false, nil
}

// PreShutdown is the pre-shutdown hook: it requests a drain and waits until no
// configured mountpoint carries a mergerfs mount any more. Mounts still there
// after HookWait, for example because the supervisor is not running, are
// lazily detached here.
//
// The returned error describes what went wrong; the caller decides whether it
// fails the hook. Talos aborts the whole reboot on a failed hook, which is
// worse than a late unmount, so the command never does.
func PreShutdown(ctx context.Context, runtime Runtime, config Config) error {
	config = normalizeConfig(config)

	if err := runtime.RequestDrain(); err != nil {
		return fmt.Errorf("request drain: %w", err)
	}

	configs, err := runtime.ReadConfigs(config.Dir)
	if err != nil {
		return fmt.Errorf("read %s: %w", config.Dir, err)
	}

	mountpoints := map[string]string{} // by config name

	for name, content := range configs {
		if mountpoint := Mountpoint(content); mountpoint != "" {
			mountpoints[name] = mountpoint
		}
	}

	started := time.Now()
	deadline := time.NewTimer(config.HookWait)
	ticker := time.NewTicker(config.HookPoll)

	defer deadline.Stop()
	defer ticker.Stop()

	for {
		// A mountpoint that cannot be checked counts as still mounted; the
		// error is reported once, with the final result.
		var (
			remaining []string
			errs      []error
		)

		for _, name := range slices.Sorted(maps.Keys(mountpoints)) {
			mounted, err := runtime.IsMounted(mountpoints[name])
			if err != nil {
				errs = append(errs, fmt.Errorf("%s: %w", name, err))
			}

			if mounted || err != nil {
				remaining = append(remaining, name)
			}
		}

		if len(remaining) == 0 {
			config.Logger.Printf("drained %d mount(s) in %s", len(mountpoints), time.Since(started).Round(time.Millisecond))

			return nil
		}

		select {
		case <-ticker.C:
			continue
		case <-ctx.Done():
		case <-deadline.C:
		}

		for _, name := range remaining {
			detached, err := runtime.DetachStale(mountpoints[name])
			if detached > 0 {
				config.Logger.Printf("%s: detached %d leftover mount(s) at %s", name, detached, mountpoints[name])
			}

			if err != nil {
				errs = append(errs, fmt.Errorf("%s: %w", name, err))
			}
		}

		return errors.Join(append(errs, fmt.Errorf("%d mount(s) still there after %s: %v", len(remaining), time.Since(started).Round(time.Millisecond), remaining))...)
	}
}
