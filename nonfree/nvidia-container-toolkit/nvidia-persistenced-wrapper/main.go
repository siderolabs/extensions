package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strconv"

	"golang.org/x/sys/unix"
)

const (
	stateFolder = "/var/run/nvidia-persistenced"
	pidFile     = stateFolder + "/" + "nvidia-persistenced.pid"
)

func main() {
	// first check if the pid file exists,
	// then check if the process is running,
	// if running try to kill it, then start the new process
	if _, err := os.Stat(pidFile); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			log.Fatalf("nvidia-persistenced-wrapper: failed to stat pid file: %s%v\n", pidFile, err)
		}
	} else {
		pid, err := getProcessId()

		if err != nil {
			log.Fatalf("nvidia-persistenced-wrapper: error reading pid file: %s%v\n", pidFile, err)
		}
		if err := killProcess(pid); err != nil {
			log.Fatalf("nvidia-persistenced-wrapper: error killing process: %d%v\n", pid, err)
		}
		// now we can remove the state directory
		if err := os.RemoveAll(stateFolder); err != nil {
			log.Fatalf("nvidia-persistenced-wrapper: error removing state directory: %s%v\n", stateFolder, err)
		}
	}

	cmd := exec.Command("/usr/local/bin/nvidia-persistenced",
		[]string{
			"--no-persistence-mode",
			"--verbose",
		}...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		log.Fatalf("nvidia-persistenced-wrapper: error starting nvidia-persistenced: %v\n", err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, unix.SIGINT, unix.SIGTERM)

	if err := cmd.Process.Signal(<-ch); err != nil {
		log.Fatalf("nvidia-persistenced-wrapper: error sending signal to nvidia-persistenced: %v\n", err)
	}

	if _, err := cmd.Process.Wait(); err != nil {
		log.Fatalf("nvidia-persistenced-wrapper: error waiting for nvidia-persistenced to exit: %v\n", err)
	}
}

func getProcessId() (int, error) {
	pidData, err := ioutil.ReadFile(pidFile)
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
