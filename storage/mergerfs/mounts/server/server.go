// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package server runs one mergerfs process per config file and stops them all
// when one of them exits or the configuration changes, so Talos restarts the
// service.
package server

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"log"
	"maps"
	"path/filepath"
	"slices"
	"strings"
	"sync"
	"time"
)

const (
	// DefaultConfigDir holds one `<name>.ini` per mergerfs mount, supplied with EtcFileConfig.
	DefaultConfigDir = "/etc/mergerfs"

	defaultPollInterval = time.Second
	defaultStopTimeout  = 8 * time.Second
	defaultHookWait     = 20 * time.Second
	defaultHookPoll     = 100 * time.Millisecond
)

// Process is a running mergerfs instance.
type Process interface {
	Done() <-chan struct{}
	Err() error
	// Stop asks the process to exit and returns once it has, or once ctx
	// expires, in which case it has been killed.
	Stop(context.Context) error
}

// Runtime provides the host operations the supervisor and the pre-shutdown hook need.
type Runtime interface {
	// ReadConfigs returns the contents of every `*.ini` file in dir, keyed by file name.
	// A missing directory is not an error: it means no mounts are configured.
	ReadConfigs(dir string) (map[string][]byte, error)
	// Start launches `mergerfs -f -o config=<path>`.
	Start(path string) (Process, error)
	// DetachStale lazily detaches the mergerfs mounts left at mountpoint and
	// returns how many there were.
	DetachStale(mountpoint string) (int, error)
	// IsMounted reports whether a mergerfs mount is visible at mountpoint.
	IsMounted(mountpoint string) (bool, error)
	// Draining reports whether a drain request is in force.
	Draining(now time.Time) (bool, error)
	// RequestDrain asks the supervisor to unmount everything and keep it unmounted.
	RequestDrain() error
}

// Config controls supervision.
type Config struct {
	Dir          string
	PollInterval time.Duration
	// StopTimeout bounds stopping the mounts. It stays below Talos's
	// ten-second host-process termination grace period.
	StopTimeout time.Duration
	// HookWait is how long the pre-shutdown hook waits for the supervisor to
	// unmount everything before detaching what is left itself.
	HookWait time.Duration
	HookPoll time.Duration
	Logger   *log.Logger
}

// DefaultConfig returns production settings.
func DefaultConfig() Config {
	return Config{
		Dir:          DefaultConfigDir,
		PollInterval: defaultPollInterval,
		StopTimeout:  defaultStopTimeout,
		HookWait:     defaultHookWait,
		HookPoll:     defaultHookPoll,
		Logger:       log.Default(),
	}
}

func normalizeConfig(config Config) Config {
	defaults := DefaultConfig()

	if config.Dir == "" {
		config.Dir = defaults.Dir
	}

	if config.PollInterval <= 0 {
		config.PollInterval = defaults.PollInterval
	}

	if config.StopTimeout <= 0 {
		config.StopTimeout = defaults.StopTimeout
	}

	if config.HookWait <= 0 {
		config.HookWait = defaults.HookWait
	}

	if config.HookPoll <= 0 {
		config.HookPoll = defaults.HookPoll
	}

	if config.Logger == nil {
		config.Logger = defaults.Logger
	}

	return config
}

type mount struct {
	name       string
	path       string
	mountpoint string
	process    Process
}

type supervisor struct {
	runtime Runtime
	config  Config
	mounts  []*mount
}

// Run mounts every configured pool and supervises the mergerfs processes until
// one of them exits, the configuration changes, a drain is requested (see
// PreShutdown) or ctx is canceled. It stops every mount before returning, and
// Talos restarts the service. The returned error says why it stopped.
func Run(ctx context.Context, runtime Runtime, config Config) error {
	config = normalizeConfig(config)

	configs, err := runtime.ReadConfigs(config.Dir)
	if err != nil {
		return fmt.Errorf("read %s: %w", config.Dir, err)
	}

	mounts, err := plan(config.Dir, configs)
	if err != nil {
		return err
	}

	if len(mounts) == 0 {
		config.Logger.Printf("no config files in %s", config.Dir)
	}

	s := &supervisor{runtime: runtime, config: config, mounts: mounts}

	ticker := time.NewTicker(config.PollInterval)
	defer ticker.Stop()

	// A drain request keeps the pools unmounted until it expires.
	for waited := false; ; waited = true {
		draining, err := runtime.Draining(time.Now())
		if err != nil {
			return fmt.Errorf("check drain request: %w", err)
		}

		if !draining {
			break
		}

		if !waited {
			config.Logger.Printf("drain request in force, waiting")
		}

		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
		}
	}

	if err = s.startAll(); err != nil {
		return errors.Join(err, s.stopAll())
	}

	for {
		select {
		case <-ctx.Done():
			return s.stopAll()
		case <-ticker.C:
		}

		if name, err := s.exited(); name != "" {
			return errors.Join(fmt.Errorf("%s: mergerfs exited: %w", name, err), s.stopAll())
		}

		draining, err := runtime.Draining(time.Now())
		if err != nil {
			return errors.Join(fmt.Errorf("check drain request: %w", err), s.stopAll())
		}

		if draining {
			config.Logger.Printf("draining: stopping all mounts")

			return s.stopAll()
		}

		// A transient read error must not unmount anything.
		current, err := runtime.ReadConfigs(config.Dir)
		if err != nil || maps.EqualFunc(current, configs, bytes.Equal) {
			continue
		}

		config.Logger.Printf("config changed, restarting")

		return s.stopAll()
	}
}

// plan turns the config files into mounts, sorted by name.
func plan(dir string, configs map[string][]byte) ([]*mount, error) {
	var mounts []*mount

	owners := map[string]string{}

	for _, name := range slices.Sorted(maps.Keys(configs)) {
		mountpoint := Mountpoint(configs[name])
		if mountpoint == "" {
			return nil, fmt.Errorf("%s: no mountpoint= in config", name)
		}

		if owner, taken := owners[mountpoint]; taken {
			return nil, fmt.Errorf("%s: mountpoint %s is used by %s", name, mountpoint, owner)
		}

		owners[mountpoint] = name

		mounts = append(mounts, &mount{
			name:       name,
			path:       filepath.Join(dir, name),
			mountpoint: mountpoint,
		})
	}

	return mounts, nil
}

func (s *supervisor) startAll() error {
	for _, m := range s.mounts {
		// Mounting over a leftover mergerfs mount would stack a second one on top.
		if err := s.detachStale(m); err != nil {
			return err
		}

		process, err := s.runtime.Start(m.path)
		if err != nil {
			return fmt.Errorf("%s: start mergerfs: %w", m.name, err)
		}

		m.process = process

		s.config.Logger.Printf("%s: started mergerfs on %s", m.name, m.mountpoint)
	}

	return nil
}

// exited returns the first mount whose mergerfs process has exited.
func (s *supervisor) exited() (string, error) {
	for _, m := range s.mounts {
		if m.process == nil {
			continue
		}

		select {
		case <-m.process.Done():
		default:
			continue
		}

		err := m.process.Err()
		if err == nil {
			err = errors.New("exited without an error")
		}

		return m.name, err
	}

	return "", nil
}

// stopAll stops every mount concurrently, bounded by StopTimeout. The service
// is stopping either way, so the stop gets its own budget.
func (s *supervisor) stopAll() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.config.StopTimeout)
	defer cancel()

	var (
		wg   sync.WaitGroup
		mu   sync.Mutex
		errs []error
	)

	for _, m := range s.mounts {
		wg.Go(func() {
			if err := s.stop(ctx, m); err != nil {
				mu.Lock()
				errs = append(errs, fmt.Errorf("%s: %w", m.name, err))
				mu.Unlock()
			}
		})
	}

	wg.Wait()

	return errors.Join(errs...)
}

func (s *supervisor) stop(ctx context.Context, m *mount) error {
	var errs []error

	if m.process != nil {
		if err := m.process.Stop(ctx); err != nil {
			errs = append(errs, fmt.Errorf("stop mergerfs: %w", err))
		}

		m.process = nil
	}

	// mergerfs detaches its own mount on SIGTERM; this covers a SIGKILL.
	if err := s.detachStale(m); err != nil {
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}

// detachStale removes the mergerfs mounts a dead instance left at the
// mountpoint, which answer ENOTCONN.
func (s *supervisor) detachStale(m *mount) error {
	detached, err := s.runtime.DetachStale(m.mountpoint)
	if detached > 0 {
		s.config.Logger.Printf("%s: unmounted %d leftover mount(s) at %s", m.name, detached, m.mountpoint)
	}

	return err
}

// Mountpoint returns the mountpoint declared by a mergerfs config file, cleaned.
//
// mergerfs reads `key=value` lines, ignoring blank lines and `#` comments, and
// accepts both `mountpoint` and its alias `mount`. Options are applied in
// order, so the last declaration wins.
func Mountpoint(config []byte) string {
	var mountpoint string

	for line := range bytes.Lines(config) {
		text := strings.TrimSpace(string(line))
		if text == "" || strings.HasPrefix(text, "#") {
			continue
		}

		key, value, ok := strings.Cut(text, "=")
		if !ok {
			continue
		}

		switch strings.TrimSpace(key) {
		case "mountpoint", "mount":
			mountpoint = strings.TrimSpace(value)
		}
	}

	if mountpoint == "" {
		return ""
	}

	return filepath.Clean(mountpoint)
}
