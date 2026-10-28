// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package server

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/siderolabs/extensions/storage/nfs-server/mounts"
)

const (
	localLibraryPath = "/usr/local/lib"
	cldPipePath      = "/var/lib/nfs/rpc_pipefs/nfsd/cld"
)

var executablePaths = map[string]string{
	"exportfs":   "/usr/local/sbin/exportfs",
	"nfsdcld":    "/usr/local/sbin/nfsdcld",
	"rpc.mountd": "/usr/local/sbin/rpc.mountd",
	"rpc.nfsd":   "/usr/local/sbin/rpc.nfsd",
}

// OSRuntime performs NFS server operations on the Talos host.
type OSRuntime struct{}

// NewOSRuntime returns a Runtime backed by host operating-system primitives.
func NewOSRuntime() OSRuntime {
	return OSRuntime{}
}

// EnsureMounts prepares the NFS kernel filesystems.
func (OSRuntime) EnsureMounts() error {
	return mounts.Ensure(mounts.OSMounter{})
}

// Start launches a supervised NFS userspace daemon.
func (OSRuntime) Start(name string, args ...string) (Process, error) {
	path, err := executablePath(name)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command(path, args...)
	cmd.Env = append(os.Environ(), "LD_LIBRARY_PATH="+localLibraryPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{Pdeathsig: syscall.SIGKILL}

	if err = cmd.Start(); err != nil {
		return nil, err
	}

	process := &osProcess{cmd: cmd, done: make(chan struct{})}
	go process.wait()

	return process, nil
}

// WaitReady waits until a daemon has completed its required initialization.
func (OSRuntime) WaitReady(ctx context.Context, name string, process Process) error {
	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-process.Done():
			return childExitError(name, process.Err())
		default:
		}

		ready, err := processReady(name, process.PID())
		if err == nil && ready {
			select {
			case <-process.Done():
				return childExitError(name, process.Err())
			default:
				return nil
			}
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-process.Done():
			return childExitError(name, process.Err())
		case <-ticker.C:
		}
	}
}

// Run executes a bounded NFS utility command.
func (OSRuntime) Run(ctx context.Context, name string, args ...string) error {
	path, err := executablePath(name)
	if err != nil {
		return err
	}

	cmd := exec.CommandContext(ctx, path, args...)
	cmd.Env = append(os.Environ(), "LD_LIBRARY_PATH="+localLibraryPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{Pdeathsig: syscall.SIGKILL}

	return cmd.Run()
}

// ReadExports reads the same export files consumed by exportfs -r.
func (OSRuntime) ReadExports() ([]byte, error) {
	return readExports(exportsPath)
}

func readExports(path string) ([]byte, error) {
	contents, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	entries, err := os.ReadDir(path + ".d")
	if errors.Is(err, os.ErrNotExist) {
		return contents, nil
	}
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), ".") || !strings.HasSuffix(entry.Name(), ".exports") {
			continue
		}
		filePath := filepath.Join(path+".d", entry.Name())
		dropIn, readErr := os.ReadFile(filePath)
		if readErr != nil {
			return nil, fmt.Errorf("read %s: %w", filePath, readErr)
		}
		// Keep logical lines separate even if the previous file has no final
		// newline or ends with a line continuation.
		contents = append(contents, '\n', '\n')
		contents = append(contents, dropIn...)
	}

	return contents, nil
}

// IsMountPoint checks the host mount table, not path existence (which can be true
// before a Talos user volume is mounted).
func (OSRuntime) IsMountPoint(path string) (bool, error) {
	contents, err := os.ReadFile("/proc/self/mountinfo")
	if err != nil {
		return false, err
	}

	return mountInfoContains(contents, path), nil
}

func mountInfoContains(contents []byte, path string) bool {
	for line := range strings.SplitSeq(string(contents), "\n") {
		fields := strings.Fields(line)
		if len(fields) < 5 {
			continue
		}
		mountpoint := strings.NewReplacer(`\040`, " ", `\011`, "\t", `\012`, "\n", `\134`, `\`).Replace(fields[4])
		if mountpoint == path {
			return true
		}
	}

	return false
}

func executablePath(name string) (string, error) {
	path, ok := executablePaths[name]
	if !ok {
		return "", fmt.Errorf("unknown NFS executable %q", name)
	}

	return path, nil
}

func processOwnsTCPPort(pid, port int) (bool, error) {
	entries, err := os.ReadDir(fmt.Sprintf("/proc/%d/fd", pid))
	if err != nil {
		return false, err
	}

	socketInodes := map[string]struct{}{}
	for _, entry := range entries {
		target, readErr := os.Readlink(filepath.Join("/proc", strconv.Itoa(pid), "fd", entry.Name()))
		if readErr != nil || !strings.HasPrefix(target, "socket:[") || !strings.HasSuffix(target, "]") {
			continue
		}

		socketInodes[strings.TrimSuffix(strings.TrimPrefix(target, "socket:["), "]")] = struct{}{}
	}

	for _, table := range []string{"/proc/net/tcp", "/proc/net/tcp6"} {
		listening, tableErr := ownsListeningSocket(table, port, socketInodes)
		if tableErr != nil {
			return false, tableErr
		}
		if listening {
			return true, nil
		}
	}

	return false, nil
}

func ownsListeningSocket(path string, port int, socketInodes map[string]struct{}) (bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer func() {
		_ = file.Close()
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 10 || fields[3] != "0A" {
			continue
		}

		address := strings.Split(fields[1], ":")
		if len(address) != 2 {
			continue
		}

		parsedPort, parseErr := strconv.ParseUint(address[1], 16, 16)
		if parseErr != nil || int(parsedPort) != port {
			continue
		}

		if _, ok := socketInodes[fields[9]]; ok {
			return true, nil
		}
	}

	return false, scanner.Err()
}

func processReady(name string, pid int) (bool, error) {
	switch name {
	case "rpc.mountd":
		return processOwnsTCPPort(pid, 20048)
	case "nfsdcld", "nfsdcld-attached":
		entries, err := os.ReadDir(fmt.Sprintf("/proc/%d/fd", pid))
		if err != nil {
			return false, err
		}

		var hasDatabase, hasInotify, hasCLDPipe bool

		for _, entry := range entries {
			target, readErr := os.Readlink(filepath.Join("/proc", fmt.Sprintf("%d", pid), "fd", entry.Name()))
			if readErr != nil {
				continue
			}

			switch {
			case target == cldPipePath:
				hasCLDPipe = true
			case target == "anon_inode:inotify":
				hasInotify = true
			case filepath.Base(target) == "main.sqlite":
				hasDatabase = true
			}
		}

		if name == "nfsdcld-attached" {
			return hasCLDPipe, nil
		}

		return hasDatabase && hasInotify, nil
	default:
		return false, fmt.Errorf("no readiness check for %q", name)
	}
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

func (process *osProcess) PID() int {
	return process.cmd.Process.Pid
}

func (process *osProcess) Done() <-chan struct{} {
	return process.done
}

func (process *osProcess) Err() error {
	process.mu.Lock()
	defer process.mu.Unlock()

	return process.err
}

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
