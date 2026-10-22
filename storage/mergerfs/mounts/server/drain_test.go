// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package server_test

import (
	"errors"
	"os"
	"path/filepath"
	"slices"
	"testing"
	"time"

	"github.com/siderolabs/extensions/storage/mergerfs/mounts/server"
)

func TestDrainFile(t *testing.T) {
	t.Parallel()

	drain := server.Drain{Dir: filepath.Join(t.TempDir(), "state"), MaxAge: time.Minute}

	now := time.Now()

	active, err := drain.Active(now)
	if err != nil || active {
		t.Fatalf("Active() before any request = %v, %v; want false", active, err)
	}

	if err = drain.Request(); err != nil {
		t.Fatalf("Request() error = %v", err)
	}

	active, err = drain.Active(now)
	if err != nil || !active {
		t.Fatalf("Active() after a request = %v, %v; want true", active, err)
	}

	// An aborted shutdown: the request expires and is removed.
	active, err = drain.Active(now.Add(2 * time.Minute))
	if err != nil || active {
		t.Fatalf("Active() after MaxAge = %v, %v; want false", active, err)
	}

	if _, err = os.Stat(filepath.Join(drain.Dir, "drain")); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("expired request was not removed: %v", err)
	}

	// A repeated request is a fresh one.
	if err = drain.Request(); err != nil {
		t.Fatalf("Request() error = %v", err)
	}

	active, err = drain.Active(now)
	if err != nil || !active {
		t.Fatalf("Active() after a repeated request = %v, %v; want true", active, err)
	}
}

func TestPreShutdownWaitsForTheSupervisor(t *testing.T) {
	t.Parallel()

	runtime := newFakeRuntime(t, map[string]string{"pool.ini": poolConfig, "pool-slow.ini": poolSlowConfig})
	start(t, runtime)

	eventually(t, "both mounts", func() bool {
		_, _, _, mounted := runtime.snapshot()

		return len(mounted) == 2
	})

	if err := server.PreShutdown(t.Context(), runtime, testConfig()); err != nil {
		t.Fatalf("PreShutdown() error = %v", err)
	}

	_, unmounts, stops, mounted := runtime.snapshot()
	if len(stops) != 2 || len(mounted) != 0 || len(unmounts) != 0 {
		t.Fatalf("expected the supervisor to stop both mounts: stops %v, unmounts %v, mounted %v", stops, unmounts, mounted)
	}

	if runtime.drainCalls != 1 {
		t.Fatalf("drain requested %d times, want 1", runtime.drainCalls)
	}
}

func TestPreShutdownDetachesLeftoversWithoutASupervisor(t *testing.T) {
	t.Parallel()

	runtime := newFakeRuntime(t, map[string]string{"pool.ini": poolConfig, "pool-slow.ini": poolSlowConfig})
	runtime.mounted["/var/mnt/pool"] = 1

	err := server.PreShutdown(t.Context(), runtime, testConfig())
	if err == nil {
		t.Fatal("PreShutdown() with nobody to drain returned no error")
	}

	_, unmounts, _, mounted := runtime.snapshot()
	if !slices.Equal(unmounts, []string{"/var/mnt/pool"}) || len(mounted) != 0 {
		t.Fatalf("expected the leftover to be detached: unmounts %v, mounted %v", unmounts, mounted)
	}
}

func TestPreShutdownWithoutMountsReturnsAtOnce(t *testing.T) {
	t.Parallel()

	runtime := newFakeRuntime(t, map[string]string{"pool.ini": poolConfig})

	started := time.Now()

	if err := server.PreShutdown(t.Context(), runtime, testConfig()); err != nil {
		t.Fatalf("PreShutdown() error = %v", err)
	}

	if elapsed := time.Since(started); elapsed > testConfig().HookWait {
		t.Fatalf("PreShutdown() waited %s with nothing mounted", elapsed)
	}
}

func TestPreShutdownReportsAFailedRequest(t *testing.T) {
	t.Parallel()

	runtime := newFakeRuntime(t, map[string]string{"pool.ini": poolConfig})
	runtime.drainErr = errors.New("read-only file system")

	if err := server.PreShutdown(t.Context(), runtime, testConfig()); !errors.Is(err, runtime.drainErr) {
		t.Fatalf("PreShutdown() error = %v, want %v", err, runtime.drainErr)
	}
}
