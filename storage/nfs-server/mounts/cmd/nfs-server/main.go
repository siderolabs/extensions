// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/siderolabs/extensions/storage/nfs-server/mounts/server"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	if err := server.Run(ctx, server.NewOSRuntime(), server.DefaultConfig()); err != nil {
		fmt.Fprintf(os.Stderr, "run NFS server: %v\n", err)

		os.Exit(1)
	}
}
