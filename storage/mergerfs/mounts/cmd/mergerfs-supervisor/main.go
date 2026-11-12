// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Command mergerfs-supervisor runs one mergerfs per /etc/mergerfs/*.ini.
//
// Without arguments it supervises the mounts. `mergerfs-supervisor pre-shutdown`
// is the Talos pre-shutdown hook: it makes the running supervisor unmount
// everything while the services owning the branches (ZFS, NFS) still run.
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/siderolabs/extensions/storage/mergerfs/mounts/server"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("mergerfs-supervisor: ")

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	runtime := server.NewOSRuntime()
	config := server.DefaultConfig()

	switch {
	case len(os.Args) == 1:
		if err := server.Run(ctx, runtime, config); err != nil {
			log.Print(err)

			os.Exit(1)
		}
	case len(os.Args) == 2 && os.Args[1] == "pre-shutdown":
		// A failed hook makes Talos abort the reboot, shutdown or upgrade;
		// a mount that is still there is the lesser problem, so only log.
		if err := server.PreShutdown(ctx, runtime, config); err != nil {
			log.Printf("pre-shutdown: %v", err)
		}
	default:
		log.Printf("usage: %s [pre-shutdown]", filepath.Base(os.Args[0]))

		os.Exit(1)
	}
}
