// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package server

import (
	"context"
	"errors"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/siderolabs/extensions/storage/mergerfs/mounts"
)

// MergerfsPath is where the extension installs the mergerfs binary.
const MergerfsPath = "/usr/local/bin/mergerfs"

// OSRuntime runs mergerfs on the Talos host.
type OSRuntime struct {
	drain Drain
}

// NewOSRuntime returns a Runtime backed by host operating-system primitives.
func NewOSRuntime() OSRuntime {
	return OSRuntime{
		drain: Drain{Dir: DefaultStateDir, MaxAge: DefaultDrainMaxAge},
	}
}

// ReadConfigs implements Runtime.
func (OSRuntime) ReadConfigs(dir string) (map[string][]byte, error) {
	entries, err := os.ReadDir(dir)
	if errors.Is(err, fs.ErrNotExist) {
		return map[string][]byte{}, nil
	}

	if err != nil {
		return nil, err
	}

	configs := map[string][]byte{}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".ini") {
			continue
		}

		content, readErr := os.ReadFile(filepath.Join(dir, entry.Name()))
		if readErr != nil {
			return nil, readErr
		}

		configs[entry.Name()] = content
	}

	return configs, nil
}

// Start implements Runtime.
func (OSRuntime) Start(path string) (Process, error) {
	cmd := exec.Command(MergerfsPath, "-f", "-o", "config="+path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// If the supervisor dies, mergerfs gets SIGTERM and unmounts on its own;
	// the next supervisor clears whatever is left before remounting.
	cmd.SysProcAttr = &syscall.SysProcAttr{Pdeathsig: syscall.SIGTERM}

	if err := cmd.Start(); err != nil {
		return nil, err
	}

	process := &osProcess{cmd: cmd, done: make(chan struct{})}
	go process.wait()

	return process, nil
}

// DetachStale implements Runtime.
func (OSRuntime) DetachStale(mountpoint string) (int, error) {
	return mounts.DetachStale(mounts.OSMounter{}, mountpoint)
}

// IsMounted implements Runtime.
func (OSRuntime) IsMounted(mountpoint string) (bool, error) {
	return mounts.IsMounted(mounts.OSMounter{}, mountpoint)
}

// Draining implements Runtime.
func (r OSRuntime) Draining(now time.Time) (bool, error) {
	return r.drain.Active(now)
}

// RequestDrain implements Runtime.
func (r OSRuntime) RequestDrain() error {
	return r.drain.Request()
}

type osProcess struct {
	cmd  *exec.Cmd
	done chan struct{}

	mu  sync.Mutex
	err error
}

func (process *osProcess) wait() {
	err := process.cmd.Wait()

	process.mu.Lock()
	process.err = err
	process.mu.Unlock()

	close(process.done)
}

func (process *osProcess) Done() <-chan struct{} {
	return process.done
}

func (process *osProcess) Err() error {
	process.mu.Lock()
	defer process.mu.Unlock()

	return process.err
}

// Stop sends SIGTERM, on which mergerfs detaches its mount and exits, and
// falls back to SIGKILL when ctx expires. It does not wait for the kill to
// take effect: a process stuck in the kernel would hold up the supervisor
// (the same choice as the nfs-server extension makes).
func (process *osProcess) Stop(ctx context.Context) error {
	select {
	case <-process.done:
		return nil
	default:
	}

	if err := process.cmd.Process.Signal(syscall.SIGTERM); err != nil && !errors.Is(err, os.ErrProcessDone) {
		return err
	}

	select {
	case <-process.done:
		return nil
	case <-ctx.Done():
		if err := process.cmd.Process.Kill(); err != nil && !errors.Is(err, os.ErrProcessDone) {
			return errors.Join(ctx.Err(), err)
		}

		return ctx.Err()
	}
}
