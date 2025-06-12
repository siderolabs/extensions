// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/goaux/decowriter"
)

const (
	// FabricManager
	fmCmdFile     = "/usr/bin/nv-fabricmanager"
	fmConfigFile  = "/usr/share/nvidia/nvswitch/fabricmanager.cfg"
	fmStopTimeout = 5 * time.Second

	// NVLSM
	smCmdFile     = "/opt/nvidia/nvlsm/sbin/nvlsm"
	smConfigFile  = "/usr/share/nvidia/nvlsm/nvlsm.conf"
	smPidFile     = "/var/run/nvlsm.pid"
	smSocket      = "/var/run/nvidia-fabricmanager/fm_sm_ipc.socket"
	smStopTimeout = 5 * time.Second
	smSocketWait  = 15 * time.Second
)

func runCommand(ctx context.Context, wg *sync.WaitGroup, doneCb func(), waitDelay time.Duration, path string, arg ...string) {
	wg.Add(1)

	cmd := exec.CommandContext(ctx, path, arg...)
	cmd.WaitDelay = waitDelay
	cmd.Cancel = func() error {
		return cmd.Process.Signal(os.Interrupt)
	}

	// TODO line writer to log module
	name := filepath.Base(path)
	cmd.Stdout = decowriter.New(bufio.NewWriter(os.Stdout), []byte(name+": "), []byte{})
	cmd.Stderr = decowriter.New(bufio.NewWriter(os.Stderr), []byte(name+": "), []byte{})

	go func() {
		log.Printf("nvidia-fabricmanager-wrapper: running command: %s %s\n", path, strings.Join(arg, " "))

		err := cmd.Run()
		if err == nil {
			log.Printf("nvidia-fabricmanager-wrapper: command %s [%d] completed successfully\n", path, cmd.Process.Pid)
		} else if exitErr, ok := err.(*exec.ExitError); ok {
			if exitErr.Exited() {
				log.Printf("nvidia-fabricmanager-wrapper: command %s [%d] exited with code %d\n", path, exitErr.Pid(),
					exitErr.ExitCode())
			} else {
				log.Printf("nvidia-fabricmanager-wrapper: command %s [%d] was terminated\n", path, exitErr.Pid())
			}
		} else {
			log.Printf("nvidia-fabricmanager-wrapper: failed to run command %s: %v\n", path, err)
		}

		wg.Done()
		doneCb()
	}()
}

func waitForFile(ctx context.Context, filepath string, timeout time.Duration) error {
	timer := time.NewTimer(timeout)
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("parent context canceled: %w", ctx.Err())
		case <-timer.C:
			return fmt.Errorf("timeout waiting for file")
		default:
			if _, err := os.Stat(filepath); err == nil {
				return nil
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	var cmdWg sync.WaitGroup

	signal.Ignore(syscall.SIGHUP)

	runCtx, gracefulShutdown := context.WithCancel(context.Background())

	signalsChan := make(chan os.Signal, 1)
	signal.Notify(signalsChan, os.Interrupt)
	signal.Notify(signalsChan, syscall.SIGTERM)

	go func() {
		received := <-signalsChan
		signal.Stop(signalsChan)
		log.Printf("nvidia-fabricmanager-wrapper: received signal '%s', initiating a graceful shutdown\n", received.String())
		gracefulShutdown()
	}()

	nvswitchPorts := findNvswitchMgmtPorts()
	for _, port := range nvswitchPorts {
		log.Printf("nvidia-fabricmanager-wrapper: found NVSwitch LPF: device=%s guid=0x%x\n", port.IBDevice, port.PortGUID)
	}

	fmSmMgmtPortGUID := ""
	if len(nvswitchPorts) > 0 {
		fmSmMgmtPortGUID = fmt.Sprintf("0x%x", nvswitchPorts[0].PortGUID)
		log.Printf("nvidia-fabricmanager-wrapper: using NVSwitch management port GUID: %s\n", fmSmMgmtPortGUID)
	} else {
		log.Println("nvidia-fabricmanager-wrapper: No InfiniBand NVSwitch detected. On Blackwell HGX baseboards and newer",
			"with NVLink 5.0+, please load kernel module 'ib_umad' for NVLSM to run along FabricManager. Otherwise it will",
			"fail to start with error NV_WARN_NOTHING_TO_DO, and GPU workloads will report CUDA_ERROR_SYSTEM_NOT_READY.")
	}

	if fmSmMgmtPortGUID != "" {
		if err := os.Mkdir(filepath.Dir(smSocket), 0755); err != nil {
			log.Printf("nvidia-fabricmanager-wrapper: error creating socket directory: %v\n", err)
		}

		runCommand(runCtx, &cmdWg, gracefulShutdown, smStopTimeout, smCmdFile, "--config", smConfigFile,
			"--guid", fmSmMgmtPortGUID, "--pid_file", smPidFile, "--log_file", "stdout")

		// vendor startup script waits for 5 seconds for NVLSM socket to be available before starting FM
		// let's wait for the actual GRPC socket to be created by the plugin
		log.Println("nvidia-fabricmanager-wrapper: waiting for socket creation at", smSocket)
		err := waitForFile(runCtx, smSocket, smSocketWait)
		if err != nil {
			log.Printf("nvidia-fabricmanager-wrapper: error waiting for socket: %v\n", err)
		} else {
			log.Println("nvidia-fabricmanager-wrapper: socket found at", smSocket)
		}
		// for safety
		time.Sleep(time.Second)
	}

	fmCmdArgs := []string{"--config", fmConfigFile}
	if fmSmMgmtPortGUID != "" {
		fmCmdArgs = append(fmCmdArgs, "--fm-sm-mgmt-port-guid", fmSmMgmtPortGUID)
	}
	runCommand(runCtx, &cmdWg, gracefulShutdown, fmStopTimeout, fmCmdFile, fmCmdArgs...)

	log.Println("nvidia-fabricmanager-wrapper: initialization completed")
	cmdWg.Wait()
}
