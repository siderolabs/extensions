// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"bytes"
	"errors"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strconv"
	"syscall"

	"golang.org/x/sys/unix"
)

const (
	stateFolder = "/var/run/nvidia-persistenced"
	pidFile     = stateFolder + "/" + "nvidia-persistenced.pid"

	driverSource = "/usr/local/glibc"
	driverTarget = "/run/nvidia/driver"
)

func main() {
	bindMounDriver()
	execPersistenced()
}

func bindMounDriver() {
	err := os.MkdirAll(driverTarget, 0755)
	if err != nil {
		log.Fatalf("nvidia-driver: failed mkdir: %v", err)
	}
	err = syscall.Mount(driverSource, driverTarget, "", syscall.MS_BIND|syscall.MS_REC, "")
	if err != nil {
		log.Fatalf("nvidia-driver: failed mount BIND: %v", err)
	}
}

func execPersistenced() {
	// ref: https://docs.nvidia.com/deploy/driver-persistence/index.html
	// first check if the pid file exists,
	// then check if the process is running,
	// if running try to kill it, then start the new process
	if _, err := os.Stat(pidFile); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			log.Fatalf("nvidia-driver: failed to stat pid file: %s%v\n", pidFile, err)
		}
	} else {
		pid, err := getProcessId()
		if err != nil {
			log.Fatalf("nvidia-driver: error reading pid file: %s%v\n", pidFile, err)
		}
		if err := killProcess(pid); err != nil {
			log.Fatalf("nvidia-driver: error killing process: %d%v\n", pid, err)
		}
		// now we can remove the state directory
		if err := os.RemoveAll(stateFolder); err != nil {
			log.Fatalf("nvidia-driver: error removing state directory: %s%v\n", stateFolder, err)
		}
	}

	cmd := exec.Command(filepath.Join(driverTarget, "/usr/bin/nvidia-persistenced"))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		log.Fatalf("nvidia-driver: error starting nvidia-persistenced: %v\n", err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, unix.SIGINT, unix.SIGTERM)

	if err := cmd.Process.Signal(<-ch); err != nil {
		log.Fatalf("nvidia-driver: error sending signal to nvidia-persistenced: %v\n", err)
	}

	if _, err := cmd.Process.Wait(); err != nil {
		log.Fatalf("nvidia-driver: error waiting for nvidia-persistenced to exit: %v\n", err)
	}
}

func getProcessId() (int, error) {
	pidData, err := os.ReadFile(pidFile)
	if err != nil {
		return 0, err
	}
	// remove any newlines
	pidData = bytes.TrimRight(pidData, "\n")
	pid, err := strconv.Atoi(string(pidData))
	if err != nil {
		return 0, err
	}
	return int(pid), nil
}

func killProcess(pid int) error {
	p, err := os.FindProcess(pid)
	if err != nil {
		return err
	}
	// ignore if process is already dead
	if err := p.Kill(); !errors.Is(err, os.ErrProcessDone) {
		return err
	}
	return nil
}
