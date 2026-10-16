// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package server_test

import (
	"context"
	"errors"
	"io"
	"log"
	"maps"
	"os"
	"slices"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/siderolabs/extensions/storage/mergerfs/mounts/server"
)

type fakeProcess struct {
	path       string
	mountpoint string
	done       chan struct{}
	once       sync.Once

	mu  sync.Mutex
	err error
}

func (process *fakeProcess) Done() <-chan struct{} { return process.done }

func (process *fakeProcess) Err() error {
	process.mu.Lock()
	defer process.mu.Unlock()

	return process.err
}

func (process *fakeProcess) exit(err error) {
	process.mu.Lock()
	process.err = err
	process.mu.Unlock()

	process.once.Do(func() { close(process.done) })
}

type fakeRuntime struct {
	t *testing.T

	mu         sync.Mutex
	configs    map[string][]byte
	running    map[string]*fakeProcess // by config path
	mounted    map[string]int          // stacked mergerfs mounts by mountpoint
	starts     []string
	unmounts   []string
	stopCalls  []string
	draining   bool
	drainErr   error
	drainCalls int
}

func newFakeRuntime(t *testing.T, configs map[string]string) *fakeRuntime {
	t.Helper()

	runtime := &fakeRuntime{
		t:       t,
		configs: map[string][]byte{},
		running: map[string]*fakeProcess{},
		mounted: map[string]int{},
	}

	for name, content := range configs {
		runtime.configs[name] = []byte(content)
	}

	return runtime
}

func (runtime *fakeRuntime) ReadConfigs(string) (map[string][]byte, error) {
	runtime.mu.Lock()
	defer runtime.mu.Unlock()

	return maps.Clone(runtime.configs), nil
}

func (runtime *fakeRuntime) Start(path string) (server.Process, error) {
	runtime.mu.Lock()
	defer runtime.mu.Unlock()

	mountpoint := mountpointFor(runtime.configs, path)

	if runtime.mounted[mountpoint] > 0 {
		runtime.t.Errorf("start %s over an existing mount", path)
	}

	process := &fakeProcess{path: path, mountpoint: mountpoint, done: make(chan struct{})}
	runtime.running[path] = process
	runtime.starts = append(runtime.starts, path)
	runtime.mounted[mountpoint]++

	return &stoppable{fakeProcess: process, runtime: runtime}, nil
}

func (runtime *fakeRuntime) DetachStale(mountpoint string) (int, error) {
	runtime.mu.Lock()
	defer runtime.mu.Unlock()

	detached := runtime.mounted[mountpoint]
	delete(runtime.mounted, mountpoint)

	for range detached {
		runtime.unmounts = append(runtime.unmounts, mountpoint)
	}

	return detached, nil
}

func (runtime *fakeRuntime) IsMounted(mountpoint string) (bool, error) {
	runtime.mu.Lock()
	defer runtime.mu.Unlock()

	return runtime.mounted[mountpoint] > 0, nil
}

func (runtime *fakeRuntime) Draining(time.Time) (bool, error) {
	runtime.mu.Lock()
	defer runtime.mu.Unlock()

	return runtime.draining, runtime.drainErr
}

func (runtime *fakeRuntime) RequestDrain() error {
	runtime.mu.Lock()
	defer runtime.mu.Unlock()

	runtime.drainCalls++

	if runtime.drainErr != nil {
		return runtime.drainErr
	}

	runtime.draining = true

	return nil
}

func (runtime *fakeRuntime) setDraining(draining bool) {
	runtime.mu.Lock()
	defer runtime.mu.Unlock()

	runtime.draining = draining
}

// crash kills the mergerfs process started from path with SIGKILL semantics:
// the process dies, its mount stays.
func (runtime *fakeRuntime) crash(path string) {
	runtime.mu.Lock()
	process := runtime.running[path]
	delete(runtime.running, path)
	runtime.mu.Unlock()

	process.exit(errors.New("signal: killed"))
}

func (runtime *fakeRuntime) setConfig(name, content string) {
	runtime.mu.Lock()
	defer runtime.mu.Unlock()

	if content == "" {
		delete(runtime.configs, name)

		return
	}

	runtime.configs[name] = []byte(content)
}

func (runtime *fakeRuntime) snapshot() (starts, unmounts, stops []string, mounted []string) {
	runtime.mu.Lock()
	defer runtime.mu.Unlock()

	for mountpoint, layers := range runtime.mounted {
		if layers > 0 {
			mounted = append(mounted, mountpoint)
		}
	}

	slices.Sort(mounted)

	return slices.Clone(runtime.starts), slices.Clone(runtime.unmounts), slices.Clone(runtime.stopCalls), mounted
}

// stoppable models mergerfs's SIGTERM behavior: it detaches its own mount.
type stoppable struct {
	*fakeProcess

	runtime *fakeRuntime
}

func (process *stoppable) Stop(context.Context) error {
	// Like osProcess.Stop: nothing to do for a process that has exited; its
	// leftover mount is the supervisor's problem.
	select {
	case <-process.done:
		return nil
	default:
	}

	process.runtime.mu.Lock()
	process.runtime.stopCalls = append(process.runtime.stopCalls, process.path)
	process.runtime.mounted[process.mountpoint]--

	if process.runtime.running[process.path] == process.fakeProcess {
		delete(process.runtime.running, process.path)
	}
	process.runtime.mu.Unlock()

	process.exit(nil)

	return nil
}

// mountpointFor resolves the mountpoint of the config at path; tests key
// configs by file name under /etc/mergerfs.
func mountpointFor(configs map[string][]byte, path string) string {
	for name, content := range configs {
		if "/etc/mergerfs/"+name == path {
			return server.Mountpoint(content)
		}
	}

	return ""
}

func testConfig() server.Config {
	return server.Config{
		Dir:          "/etc/mergerfs",
		PollInterval: 5 * time.Millisecond,
		StopTimeout:  50 * time.Millisecond,
		HookWait:     200 * time.Millisecond,
		HookPoll:     5 * time.Millisecond,
		Logger:       log.New(io.Discard, "", 0),
	}
}

type harness struct {
	cancel context.CancelFunc
	result chan error
}

func start(t *testing.T, runtime *fakeRuntime) *harness {
	t.Helper()

	ctx, cancel := context.WithCancel(t.Context())

	h := &harness{
		cancel: cancel,
		result: make(chan error, 1),
	}

	go func() {
		h.result <- server.Run(ctx, runtime, testConfig())
	}()

	t.Cleanup(func() {
		cancel()
		<-h.result
	})

	return h
}

// wait returns Run's result, leaving it for the cleanup.
func (h *harness) wait(t *testing.T) error {
	t.Helper()

	select {
	case err := <-h.result:
		h.result <- err

		return err
	case <-time.After(5 * time.Second):
		t.Fatal("Run did not return")

		return nil
	}
}

func eventually(t *testing.T, what string, cond func() bool) {
	t.Helper()

	deadline := time.Now().Add(5 * time.Second)
	for !cond() {
		if time.Now().After(deadline) {
			t.Fatalf("timed out waiting for %s", what)
		}

		time.Sleep(time.Millisecond)
	}
}

const (
	poolConfig     = "branches=/var/mnt/ssd:/var/mnt/disk1\nmountpoint=/var/mnt/pool\n"
	poolSlowConfig = "# slow tier\nbranches=/var/mnt/disk1\nmountpoint = /var/mnt/pool-slow\n"
)

func TestStartsOneMergerfsPerConfig(t *testing.T) {
	t.Parallel()

	runtime := newFakeRuntime(t, map[string]string{"pool.ini": poolConfig, "pool-slow.ini": poolSlowConfig})
	start(t, runtime)

	eventually(t, "both mounts", func() bool {
		_, _, _, mounted := runtime.snapshot()

		return slices.Equal(mounted, []string{"/var/mnt/pool", "/var/mnt/pool-slow"})
	})

	starts, _, _, _ := runtime.snapshot()
	if !slices.Equal(starts, []string{"/etc/mergerfs/pool-slow.ini", "/etc/mergerfs/pool.ini"}) {
		t.Fatalf("unexpected starts %v", starts)
	}
}

func TestDetachesLeftoverMountsBeforeStarting(t *testing.T) {
	t.Parallel()

	runtime := newFakeRuntime(t, map[string]string{"pool.ini": poolConfig})
	runtime.mounted["/var/mnt/pool"] = 2 // two crashed instances stacked
	start(t, runtime)

	// fakeRuntime.Start fails the test if it is asked to mount over a leftover.
	eventually(t, "start", func() bool {
		starts, _, _, _ := runtime.snapshot()

		return len(starts) == 1
	})

	_, unmounts, _, mounted := runtime.snapshot()
	if !slices.Equal(unmounts, []string{"/var/mnt/pool", "/var/mnt/pool"}) || !slices.Equal(mounted, []string{"/var/mnt/pool"}) {
		t.Fatalf("expected the leftovers to be detached first: unmounts %v, mounted %v", unmounts, mounted)
	}
}

func TestExitedMergerfsStopsEverything(t *testing.T) {
	t.Parallel()

	runtime := newFakeRuntime(t, map[string]string{"pool.ini": poolConfig, "pool-slow.ini": poolSlowConfig})
	h := start(t, runtime)

	eventually(t, "both mounts", func() bool {
		_, _, _, mounted := runtime.snapshot()

		return len(mounted) == 2
	})

	runtime.crash("/etc/mergerfs/pool.ini")

	err := h.wait(t)
	if err == nil || !strings.Contains(err.Error(), "pool.ini: mergerfs exited: signal: killed") {
		t.Fatalf("Run() error = %v, want the exited mount named", err)
	}

	// The other pool is stopped, the crashed one's leftover detached, and Talos restarts the service.
	_, unmounts, stops, mounted := runtime.snapshot()
	if !slices.Equal(stops, []string{"/etc/mergerfs/pool-slow.ini"}) || !slices.Equal(unmounts, []string{"/var/mnt/pool"}) || len(mounted) != 0 {
		t.Fatalf("stops %v, unmounts %v, mounted %v", stops, unmounts, mounted)
	}
}

func TestConfigChangeStopsEverything(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		name   string
		change func(*fakeRuntime)
	}{
		{"added", func(r *fakeRuntime) { r.setConfig("pool-slow.ini", poolSlowConfig) }},
		{"changed", func(r *fakeRuntime) { r.setConfig("pool.ini", poolConfig+"minfreespace=4G\n") }},
		{"removed", func(r *fakeRuntime) { r.setConfig("pool.ini", "") }},
	} {
		runtime := newFakeRuntime(t, map[string]string{"pool.ini": poolConfig})
		h := start(t, runtime)

		eventually(t, "first start", func() bool {
			starts, _, _, _ := runtime.snapshot()

			return len(starts) == 1
		})

		tc.change(runtime)

		if err := h.wait(t); err != nil {
			t.Fatalf("%s: Run() error = %v", tc.name, err)
		}

		// Nothing is started by this instance: the restarted service reads the new configuration.
		starts, _, stops, mounted := runtime.snapshot()
		if len(starts) != 1 || !slices.Equal(stops, []string{"/etc/mergerfs/pool.ini"}) || len(mounted) != 0 {
			t.Fatalf("%s: starts %v, stops %v, mounted %v", tc.name, starts, stops, mounted)
		}
	}
}

func TestRejectsBadConfigs(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		configs map[string]string
		want    string
	}{
		{map[string]string{"broken.ini": "branches=/var/mnt/disk1\n", "pool.ini": poolConfig}, "broken.ini: no mountpoint= in config"},
		{map[string]string{"a.ini": poolConfig, "b.ini": "branches=/var/mnt/other\nmountpoint=/var/mnt/pool\n"}, "b.ini: mountpoint /var/mnt/pool is used by a.ini"},
	} {
		runtime := newFakeRuntime(t, tc.configs)
		runtime.mounted["/var/mnt/pool"] = 1 // whatever is there is not ours to touch

		err := server.Run(t.Context(), runtime, testConfig())
		if err == nil || !strings.Contains(err.Error(), tc.want) {
			t.Fatalf("Run() error = %v, want %q", err, tc.want)
		}

		starts, unmounts, _, _ := runtime.snapshot()
		if len(starts) != 0 || len(unmounts) != 0 {
			t.Fatalf("a rejected configuration touched the host: starts %v, unmounts %v", starts, unmounts)
		}
	}
}

func TestDrainStopsEverythingUntilItEnds(t *testing.T) {
	t.Parallel()

	runtime := newFakeRuntime(t, map[string]string{"pool.ini": poolConfig, "pool-slow.ini": poolSlowConfig})
	h := start(t, runtime)

	eventually(t, "both mounts", func() bool {
		_, _, _, mounted := runtime.snapshot()

		return len(mounted) == 2
	})

	runtime.setDraining(true)

	if err := h.wait(t); err != nil {
		t.Fatalf("Run() error = %v", err)
	}

	if _, _, stops, mounted := runtime.snapshot(); len(stops) != 2 || len(mounted) != 0 {
		t.Fatalf("expected both mounts stopped: stops %v, mounted %v", stops, mounted)
	}

	// The restarted service waits while the request is in force.
	start(t, runtime)
	time.Sleep(50 * time.Millisecond)

	if starts, _, _, _ := runtime.snapshot(); len(starts) != 2 {
		t.Fatalf("mounts came back while draining: starts %v", starts)
	}

	// An aborted shutdown: the request expires and everything comes back.
	runtime.setDraining(false)

	eventually(t, "both mounts again", func() bool {
		_, _, _, mounted := runtime.snapshot()

		return slices.Equal(mounted, []string{"/var/mnt/pool", "/var/mnt/pool-slow"})
	})
}

func TestCancelStopsEverything(t *testing.T) {
	t.Parallel()

	runtime := newFakeRuntime(t, map[string]string{"pool.ini": poolConfig, "pool-slow.ini": poolSlowConfig})
	h := start(t, runtime)

	eventually(t, "both mounts", func() bool {
		_, _, _, mounted := runtime.snapshot()

		return len(mounted) == 2
	})

	h.cancel()

	if err := h.wait(t); err != nil {
		t.Fatalf("Run() error = %v", err)
	}

	_, _, stops, mounted := runtime.snapshot()
	if len(stops) != 2 || len(mounted) != 0 {
		t.Fatalf("expected both mounts stopped, got stops %v, mounted %v", stops, mounted)
	}
}

func TestMountpoint(t *testing.T) {
	t.Parallel()

	for name, tc := range map[string]struct {
		config string
		want   string
	}{
		"plain":           {"branches=/a:/b\nmountpoint=/var/mnt/pool\n", "/var/mnt/pool"},
		"spaces":          {"  mountpoint = /var/mnt/pool  \n", "/var/mnt/pool"},
		"trailing slash":  {"mountpoint=/var/mnt/pool/\n", "/var/mnt/pool"},
		"alias":           {"mount=/var/mnt/pool\n", "/var/mnt/pool"},
		"last wins":       {"mountpoint=/var/mnt/a\nmountpoint=/var/mnt/b\n", "/var/mnt/b"},
		"comment ignored": {"# mountpoint=/var/mnt/a\nmountpoint=/var/mnt/b\n", "/var/mnt/b"},
		"no newline":      {"mountpoint=/var/mnt/pool", "/var/mnt/pool"},
		"missing":         {"branches=/a\n", ""},
		"flag without =":  {"mountpoint\n", ""},
	} {
		if got := server.Mountpoint([]byte(tc.config)); got != tc.want {
			t.Errorf("%s: Mountpoint() = %q, want %q", name, got, tc.want)
		}
	}
}

func TestManifestRequiresHostRunnerCapableTalos(t *testing.T) {
	t.Parallel()

	contents, err := os.ReadFile("../../manifest.yaml.tmpl")
	if err != nil {
		t.Fatalf("read manifest: %v", err)
	}

	if !slices.Contains([]string{">= v1.14.0", ">= v1.15.0"}, manifestTalosVersion(string(contents))) {
		t.Fatalf("manifest Talos compatibility does not require host-runner support: %s", contents)
	}
}

func manifestTalosVersion(contents string) string {
	const prefix = "      version: \""

	for _, line := range strings.Split(contents, "\n") {
		if strings.HasPrefix(line, prefix) && strings.HasSuffix(line, "\"") {
			return strings.TrimSuffix(strings.TrimPrefix(line, prefix), "\"")
		}
	}

	return ""
}
