// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"context"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
	"time"
)

func testDir(t *testing.T) string {
	t.Helper()
	dir, err := os.MkdirTemp("/tmp", "powerd-test-")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(dir) })
	return dir
}

func helper(role, socket, record string) *exec.Cmd {
	cmd := exec.Command(os.Args[0], "-test.run=^TestHelperProcess$")
	cmd.Env = append(os.Environ(), "POWERD_TEST_ROLE="+role, "POWERD_TEST_SOCKET="+socket, "POWERD_TEST_RECORD="+record)
	return cmd
}

func TestHelperProcess(t *testing.T) {
	role := os.Getenv("POWERD_TEST_ROLE")
	if role == "" {
		return
	}
	if role == "fail" {
		os.Exit(3)
	}
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer cancel()
	if role == "bus" {
		conn, err := net.ListenUnix("unix", &net.UnixAddr{Name: os.Getenv("POWERD_TEST_SOCKET"), Net: "unix"})
		if err != nil {
			os.Exit(2)
		}
		defer conn.Close()
	}
	record := os.Getenv("POWERD_TEST_RECORD")
	f, err := os.OpenFile(record, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		os.Exit(2)
	}
	_, _ = f.WriteString("start-" + role + "\n")
	<-ctx.Done()
	_, _ = f.WriteString("stop-" + role + "\n")
	_ = f.Close()
	os.Exit(0)
}

func TestPowerFailureStopsBus(t *testing.T) {
	dir := testDir(t)
	socket, record := filepath.Join(dir, "bus.sock"), filepath.Join(dir, "record")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := supervise(ctx, helper("bus", socket, record), helper("fail", socket, record), socket)
	if err == nil || !strings.Contains(err.Error(), "nvidia-powerd exited unexpectedly") {
		t.Fatalf("unexpected result: %v", err)
	}
	content, err := os.ReadFile(record)
	if err != nil || !strings.Contains(string(content), "stop-bus") {
		t.Fatalf("D-Bus was not stopped: %s, %v", content, err)
	}
}

func TestCancellationStopsPowerBeforeBus(t *testing.T) {
	dir := testDir(t)
	socket, record := filepath.Join(dir, "bus.sock"), filepath.Join(dir, "record")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	done := make(chan error, 1)
	go func() { done <- supervise(ctx, helper("bus", socket, record), helper("power", socket, record), socket) }()
	deadline := time.Now().Add(3 * time.Second)
	for {
		content, _ := os.ReadFile(record)
		if strings.Contains(string(content), "start-power") {
			break
		}
		if time.Now().After(deadline) {
			t.Fatal("power process did not start")
		}
		time.Sleep(10 * time.Millisecond)
	}
	cancel()
	select {
	case err := <-done:
		if err != nil {
			t.Fatal(err)
		}
	case <-time.After(3 * time.Second):
		t.Fatal("shutdown did not complete")
	}
	content, _ := os.ReadFile(record)
	power, bus := strings.Index(string(content), "stop-power"), strings.Index(string(content), "stop-bus")
	if power < 0 || bus < 0 || power >= bus {
		t.Fatalf("incorrect shutdown order: %s", content)
	}
}

func TestBusFailureDoesNotStartPower(t *testing.T) {
	dir := testDir(t)
	socket, record := filepath.Join(dir, "bus.sock"), filepath.Join(dir, "record")
	err := supervise(context.Background(), helper("fail", socket, record), helper("power", socket, record), socket)
	if err == nil || !strings.Contains(err.Error(), "D-Bus exited before becoming ready") {
		t.Fatalf("unexpected result: %v", err)
	}
	if _, err := os.Stat(record); !os.IsNotExist(err) {
		t.Fatal("power process started without D-Bus")
	}
}

func TestDeviceLinksKeepSyslogPrivate(t *testing.T) {
	host, dev := t.TempDir(), t.TempDir()
	for _, name := range []string{"nvidia0", "nvidiactl", "cpu", "log", "null"} {
		if err := os.WriteFile(filepath.Join(host, name), nil, 0600); err != nil {
			t.Fatal(err)
		}
	}
	if err := linkDevices(host, dev); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"nvidia0", "nvidiactl", "cpu"} {
		target, err := os.Readlink(filepath.Join(dev, name))
		if err != nil || target != filepath.Join(host, name) {
			t.Fatalf("incorrect device link %s: %s %v", name, target, err)
		}
	}
	for _, name := range []string{"log", "null"} {
		if _, err := os.Lstat(filepath.Join(dev, name)); !os.IsNotExist(err) {
			t.Fatalf("host %s should not be linked", name)
		}
	}
}
