// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

const (
	loader      = "/usr/local/lib/ld-linux-x86-64.so.2"
	libraryPath = "/usr/lib:/usr/local/lib"
	busSocket   = "/run/nvidia-powerd/system_bus_socket"
)

type child struct {
	cmd  *exec.Cmd
	done chan struct{}
	err  error
}

func start(cmd *exec.Cmd) (*child, error) {
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	if err := cmd.Start(); err != nil {
		return nil, err
	}
	p := &child{cmd: cmd, done: make(chan struct{})}
	go func() {
		p.err = cmd.Wait()
		close(p.done)
	}()
	return p, nil
}

func (p *child) stop() {
	select {
	case <-p.done:
		return
	default:
	}
	_ = p.cmd.Process.Signal(syscall.SIGTERM)
	timer := time.NewTimer(5 * time.Second)
	defer timer.Stop()
	select {
	case <-p.done:
	case <-timer.C:
		_ = p.cmd.Process.Kill()
		<-p.done
	}
}

func waitForBus(ctx context.Context, bus *child, socket string) error {
	timer := time.NewTimer(10 * time.Second)
	defer timer.Stop()
	tick := time.NewTicker(20 * time.Millisecond)
	defer tick.Stop()
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-bus.done:
			return fmt.Errorf("D-Bus exited before becoming ready: %v", bus.err)
		case <-timer.C:
			return errors.New("timed out waiting for private D-Bus socket")
		case <-tick.C:
			info, err := os.Stat(socket)
			if err == nil && info.Mode()&os.ModeSocket != 0 {
				return nil
			}
		}
	}
}

func supervise(ctx context.Context, busCmd, powerCmd *exec.Cmd, socket string) error {
	bus, err := start(busCmd)
	if err != nil {
		return fmt.Errorf("start D-Bus: %w", err)
	}
	defer bus.stop()
	if err := waitForBus(ctx, bus, socket); err != nil {
		return err
	}
	power, err := start(powerCmd)
	if err != nil {
		return fmt.Errorf("start nvidia-powerd: %w", err)
	}
	// Stop powerd first, allowing it to restore power state while D-Bus is alive.
	defer power.stop()
	log.Println("nvidia-powerd and private D-Bus started")
	select {
	case <-ctx.Done():
		return nil
	case <-bus.done:
		return fmt.Errorf("D-Bus exited unexpectedly: %v", bus.err)
	case <-power.done:
		return fmt.Errorf("nvidia-powerd exited unexpectedly: %v", power.err)
	}
}

func linkDevices(hostDev, dev string) error {
	entries, err := os.ReadDir(hostDev)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.Name() != "cpu" && !strings.HasPrefix(entry.Name(), "nvidia") {
			continue
		}
		if err := os.Symlink(filepath.Join(hostDev, entry.Name()), filepath.Join(dev, entry.Name())); err != nil && !errors.Is(err, os.ErrExist) {
			return err
		}
	}
	return nil
}

func relaySyslog(socket string) (func(), error) {
	conn, err := net.ListenUnixgram("unixgram", &net.UnixAddr{Name: socket, Net: "unixgram"})
	if err != nil {
		return nil, err
	}
	go func() {
		buf := make([]byte, 8192)
		for {
			n, _, err := conn.ReadFromUnix(buf)
			if err != nil {
				return
			}
			message := strings.TrimSpace(strings.TrimRight(string(buf[:n]), "\x00"))
			if strings.HasPrefix(message, "<") {
				if end := strings.IndexByte(message, '>'); end >= 0 {
					message = message[end+1:]
				}
			}
			log.Print(message)
		}
	}()
	return func() { _ = conn.Close(); _ = os.Remove(socket) }, nil
}

func run(ctx context.Context) error {
	if err := linkDevices("/host/dev", "/dev"); err != nil {
		return fmt.Errorf("expose CPU and NVIDIA devices: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(busSocket), 0700); err != nil {
		return err
	}
	// D-Bus clients need a machine ID; this bus is private to the service.
	if _, err := os.Stat("/etc/machine-id"); errors.Is(err, os.ErrNotExist) {
		var id [16]byte
		if _, err := rand.Read(id[:]); err != nil {
			return err
		}
		if err := os.WriteFile("/etc/machine-id", []byte(hex.EncodeToString(id[:])+"\n"), 0644); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	// The runtime directory and /dev/log are private to this service container.
	for _, socket := range []string{busSocket, "/dev/log"} {
		if err := os.Remove(socket); err != nil && !errors.Is(err, os.ErrNotExist) {
			return err
		}
	}
	stopLog, err := relaySyslog("/dev/log")
	if err != nil {
		return fmt.Errorf("start private syslog relay: %w", err)
	}
	defer stopLog()
	bus := exec.Command(loader, "--library-path", libraryPath, "/usr/bin/dbus-daemon", "--nofork", "--nopidfile", "--config-file=/etc/dbus-1/system.conf")
	power := exec.Command(loader, "--library-path", libraryPath, "/usr/bin/nvidia-powerd")
	bus.Env = append(os.Environ(), "LD_LIBRARY_PATH="+libraryPath)
	power.Env = append(os.Environ(), "LD_LIBRARY_PATH="+libraryPath, "DBUS_SYSTEM_BUS_ADDRESS=unix:path="+busSocket)
	return supervise(ctx, bus, power, busSocket)
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()
	if err := run(ctx); err != nil && !errors.Is(err, context.Canceled) {
		log.Print(err)
		os.Exit(1)
	}
}
