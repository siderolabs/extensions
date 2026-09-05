// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package server_test

import (
	"context"
	"errors"
	"net"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/siderolabs/extensions/storage/nfs-server/mounts/server"
)

type fakeProcess struct {
	pid    int
	done   chan struct{}
	err    error
	onStop func()
	once   sync.Once
}

func newFakeProcess(pid int, onStop func()) *fakeProcess {
	return &fakeProcess{pid: pid, done: make(chan struct{}), onStop: onStop}
}

func (process *fakeProcess) PID() int              { return process.pid }
func (process *fakeProcess) Done() <-chan struct{} { return process.done }
func (process *fakeProcess) Err() error            { return process.err }
func (process *fakeProcess) Stop(context.Context) error {
	process.onStop()
	process.once.Do(func() { close(process.done) })

	return nil
}
func (process *fakeProcess) fail(err error) {
	process.err = err
	process.once.Do(func() { close(process.done) })
}

type fakeRuntime struct {
	mu        sync.Mutex
	calls     []string
	processes map[string]*fakeProcess
	exports   []byte
	reloaded  chan struct{}
	blocked   map[string]bool
}

func newFakeRuntime() *fakeRuntime {
	return &fakeRuntime{
		processes: map[string]*fakeProcess{},
		exports:   []byte("/srv *(rw)\n"),
		reloaded:  make(chan struct{}, 1),
		blocked:   map[string]bool{},
	}
}

func (runtime *fakeRuntime) record(call string) {
	runtime.mu.Lock()
	defer runtime.mu.Unlock()

	runtime.calls = append(runtime.calls, call)

	if call == "run exportfs -r -v" {
		select {
		case runtime.reloaded <- struct{}{}:
		default:
		}
	}
}

func (runtime *fakeRuntime) EnsureMounts() error {
	runtime.record("mount filesystems")

	return nil
}

func (runtime *fakeRuntime) Start(name string, args ...string) (server.Process, error) {
	runtime.record("start " + name)

	process := newFakeProcess(len(runtime.processes)+100, func() { runtime.record("stop " + name) })
	runtime.processes[name] = process

	return process, nil
}

func (runtime *fakeRuntime) WaitReady(_ context.Context, name string, _ server.Process) error {
	runtime.record("ready " + name)

	return nil
}

func (runtime *fakeRuntime) Run(ctx context.Context, name string, args ...string) error {
	call := "run " + name
	for _, arg := range args {
		call += " " + arg
	}

	runtime.record(call)

	runtime.mu.Lock()
	blocked := runtime.blocked[call]
	runtime.mu.Unlock()

	if blocked {
		<-ctx.Done()

		return ctx.Err()
	}

	return nil
}

func (runtime *fakeRuntime) ReadFile(string) ([]byte, error) {
	runtime.mu.Lock()
	defer runtime.mu.Unlock()

	return slices.Clone(runtime.exports), nil
}

func (runtime *fakeRuntime) setExports(contents string) {
	runtime.mu.Lock()
	defer runtime.mu.Unlock()

	runtime.exports = []byte(contents)
}

func (runtime *fakeRuntime) block(call string) {
	runtime.mu.Lock()
	defer runtime.mu.Unlock()

	runtime.blocked[call] = true
}

func (runtime *fakeRuntime) snapshot() []string {
	runtime.mu.Lock()
	defer runtime.mu.Unlock()

	return slices.Clone(runtime.calls)
}

func TestRunStartsComponentsOnlyAfterTheirPrerequisitesAreReady(t *testing.T) {
	t.Parallel()

	runtime := newFakeRuntime()
	ctx, cancel := context.WithCancel(t.Context())

	errCh := make(chan error, 1)
	go func() {
		errCh <- server.Run(ctx, runtime, server.Config{ReloadInterval: time.Hour})
	}()

	select {
	case <-runtime.reloaded:
	case <-time.After(time.Second):
		t.Fatal("server startup timed out")
	}

	cancel()

	if err := <-errCh; err != nil {
		t.Fatalf("Run() error = %v", err)
	}

	wantPrefix := []string{
		"mount filesystems",
		"start rpc.mountd",
		"ready rpc.mountd",
		"start nfsdcld",
		"ready nfsdcld",
		"run exportfs -r -v",
		"run rpc.nfsd --port 2049 8",
		"ready nfsdcld-attached",
		"run rpc.nfsd 0",
		"run exportfs -a -u",
	}

	got := runtime.snapshot()
	if len(got) != len(wantPrefix)+2 || !slices.Equal(got[:len(wantPrefix)], wantPrefix) {
		t.Fatalf("Run() calls = %q, want prefix %q followed by both child stops", got, wantPrefix)
	}
	for _, call := range []string{"stop nfsdcld", "stop rpc.mountd"} {
		if !slices.Contains(got[len(wantPrefix):], call) {
			t.Fatalf("Run() calls = %q, want %q after cleanup commands", got, call)
		}
	}
}

func TestRunReloadsExportsWhenConfigurationChanges(t *testing.T) {
	t.Parallel()

	runtime := newFakeRuntime()
	ctx, cancel := context.WithCancel(t.Context())
	defer cancel()

	errCh := make(chan error, 1)
	go func() {
		errCh <- server.Run(ctx, runtime, server.Config{ReloadInterval: time.Millisecond})
	}()

	select {
	case <-runtime.reloaded: // initial load
	case <-time.After(time.Second):
		t.Fatal("initial export load timed out")
	}

	runtime.setExports("/srv *(ro)\n")

	select {
	case <-runtime.reloaded:
	case <-time.After(time.Second):
		t.Fatal("export reload timed out")
	}

	cancel()

	if err := <-errCh; err != nil {
		t.Fatalf("Run() error = %v", err)
	}
}

func TestRunStopsKernelServerWhenTrackedDaemonExits(t *testing.T) {
	t.Parallel()

	runtime := newFakeRuntime()
	errCh := make(chan error, 1)
	go func() {
		errCh <- server.Run(t.Context(), runtime, server.Config{ReloadInterval: time.Hour})
	}()

	select {
	case <-runtime.reloaded:
	case <-time.After(time.Second):
		t.Fatal("server startup timed out")
	}

	runtime.processes["nfsdcld"].fail(errors.New("nfsdcld crashed"))

	if err := <-errCh; !errors.Is(err, server.ErrChildExited) {
		t.Fatalf("Run() error = %v, want ErrChildExited", err)
	}

	calls := runtime.snapshot()
	stopIndex := slices.Index(calls, "run rpc.nfsd 0")
	unexportIndex := slices.Index(calls, "run exportfs -a -u")
	if stopIndex < 0 || unexportIndex < 0 || stopIndex > unexportIndex {
		t.Fatalf("shutdown calls = %q, want rpc.nfsd 0 before exportfs -a -u", calls)
	}
}

func TestRunTimesOutBlockedStartupCommand(t *testing.T) {
	t.Parallel()

	runtime := newFakeRuntime()
	runtime.block("run exportfs -r -v")

	err := server.Run(t.Context(), runtime, server.Config{
		ReloadInterval: time.Hour,
		CommandTimeout: 20 * time.Millisecond,
		CleanupTimeout: 100 * time.Millisecond,
	})
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("Run() error = %v, want context deadline exceeded", err)
	}

	calls := runtime.snapshot()
	for _, call := range []string{"stop nfsdcld", "stop rpc.mountd"} {
		if !slices.Contains(calls, call) {
			t.Fatalf("Run() calls = %q, want %q", calls, call)
		}
	}
}

func TestRunAttemptsAllCleanupAfterBlockedKernelStop(t *testing.T) {
	t.Parallel()

	runtime := newFakeRuntime()
	runtime.block("run rpc.nfsd 0")

	ctx, cancel := context.WithCancel(t.Context())
	errCh := make(chan error, 1)
	go func() {
		errCh <- server.Run(ctx, runtime, server.Config{
			ReloadInterval: time.Hour,
			CommandTimeout: 20 * time.Millisecond,
			CleanupTimeout: 100 * time.Millisecond,
		})
	}()

	select {
	case <-runtime.reloaded:
	case <-time.After(time.Second):
		t.Fatal("server startup timed out")
	}

	cancel()

	if err := <-errCh; !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("Run() error = %v, want blocked cleanup deadline", err)
	}

	calls := runtime.snapshot()
	for _, call := range []string{
		"run rpc.nfsd 0",
		"run exportfs -a -u",
		"stop nfsdcld",
		"stop rpc.mountd",
	} {
		if !slices.Contains(calls, call) {
			t.Fatalf("Run() calls = %q, want cleanup attempt %q", calls, call)
		}
	}
}

func TestOSRuntimeRejectsMountdPortOwnedByAnotherProcess(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:20048")
	if err != nil {
		t.Skipf("port 20048 is unavailable: %v", err)
	}
	defer func() {
		if closeErr := listener.Close(); closeErr != nil {
			t.Errorf("close listener: %v", closeErr)
		}
	}()

	ctx, cancel := context.WithTimeout(t.Context(), 50*time.Millisecond)
	defer cancel()

	process := newFakeProcess(1, func() {})
	err = server.NewOSRuntime().WaitReady(ctx, "rpc.mountd", process)
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("WaitReady() error = %v, want context deadline exceeded", err)
	}
}

func TestDefaultCleanupDeadlinePrecedesTalosKillGrace(t *testing.T) {
	t.Parallel()

	const talosKillGrace = 10 * time.Second
	if cleanupTimeout := server.DefaultConfig().CleanupTimeout; cleanupTimeout >= talosKillGrace {
		t.Fatalf("cleanup timeout = %s, must be less than Talos kill grace %s", cleanupTimeout, talosKillGrace)
	}
}

func TestManifestRequiresHostRunnerCapableTalos(t *testing.T) {
	t.Parallel()

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller() failed")
	}

	contents, err := os.ReadFile(filepath.Join(filepath.Dir(filename), "..", "..", "manifest.yaml.tmpl"))
	if err != nil {
		t.Fatalf("read manifest: %v", err)
	}

	if !slices.Contains([]string{">= v1.14.0", ">= v1.15.0"}, manifestTalosVersion(string(contents))) {
		t.Fatalf("manifest Talos compatibility does not require host-runner support: %s", contents)
	}
}

func manifestTalosVersion(contents string) string {
	for _, line := range strings.Split(contents, "\n") {
		const prefix = "      version: \""
		if strings.HasPrefix(line, prefix) && strings.HasSuffix(line, "\"") {
			return strings.TrimSuffix(strings.TrimPrefix(line, prefix), "\"")
		}
	}

	return ""
}
