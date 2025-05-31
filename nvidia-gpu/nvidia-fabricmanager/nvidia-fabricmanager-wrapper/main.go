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
	fmCmdFile      = "/usr/local/bin/nv-fabricmanager"
	fmConfigFile   = "/usr/local/share/nvidia/nvswitch/fabricmanager.cfg"
	fmStopDeadline = 5 * time.Second

	// NVLSM
	smCmdFile      = "/usr/bin/nvlsm"
	smConfigFile   = "/usr/share/nvidia/nvlsm/nvlsm.cfg"
	smStateFolder  = "/var/run/nvidia-nvlsm"
	smPidFile      = smStateFolder + "/" + "nvlsm.pid"
	smStopDeadline = 5 * time.Second
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
		runCommand(runCtx, &cmdWg, gracefulShutdown, smStopDeadline, smCmdFile, "--config", smConfigFile,
			"--guid", fmSmMgmtPortGUID, "--pid_file", smPidFile, "--log_file", "stdout")
		// runCommand(runCtx, &cmdWg, gracefulShutdown, smStopDeadline, "/usr/bin/sleep", "2")
	}

	fmCmdArgs := []string{"--config", fmConfigFile}
	if fmSmMgmtPortGUID != "" {
		fmCmdArgs = append(fmCmdArgs, "--fm-sm-mgmt-port-guid", fmSmMgmtPortGUID)
	}
	runCommand(runCtx, &cmdWg, gracefulShutdown, fmStopDeadline, fmCmdFile, fmCmdArgs...)
	// runCommand(runCtx, &cmdWg, gracefulShutdown, fmStopDeadline, "/usr/bin/sleep", "2")

	log.Println("nvidia-fabricmanager-wrapper: initialization completed")
	cmdWg.Wait()
}
