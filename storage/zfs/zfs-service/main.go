// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"context"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"time"

	"golang.org/x/sys/unix"
)

// defaultExportTimeout bounds the shutdown `zpool export -a` so this PID 1 exits
// before Talos's fixed 10s SIGTERM->SIGKILL grace for extension services,
// leaving headroom for `zfs unmount -au`, the 1s WaitDelay, and a clean process
// exit so the containerd shim can unmount the /dev rbind in order. Override with
// the ZFS_EXPORT_TIMEOUT env var (a Go duration); 0 skips the export entirely
// (pre-v1.13.0 behavior, before #1028 added the export).
const defaultExportTimeout = 8 * time.Second

func main() {
	cmd := exec.Command("/usr/local/sbin/zpool", "import", "-fal")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("zfs-service: zpool import error: %v\n", err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, unix.SIGINT, unix.SIGTERM)
	<-ch

	cmd = exec.Command("/usr/local/sbin/zfs", "unmount", "-au")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Fatalf("zfs-service: zfs unmount error: %v\n", err)
	}

	// ZFS_EXPORT_TIMEOUT overrides the bound; 0 skips the export entirely.
	timeout := defaultExportTimeout
	if v := os.Getenv("ZFS_EXPORT_TIMEOUT"); v != "" {
		d, err := time.ParseDuration(v)
		if err != nil {
			log.Printf("zfs-service: invalid ZFS_EXPORT_TIMEOUT %q: %v; using %s\n", v, err, timeout)
		} else {
			timeout = d
		}
	}

	if timeout == 0 {
		log.Printf("zfs-service: ZFS_EXPORT_TIMEOUT=0, skipping zpool export; pools will be re-imported on next boot\n")

		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	cmd = exec.CommandContext(ctx, "/usr/local/sbin/zpool", "export", "-a")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.WaitDelay = time.Second

	if err := cmd.Run(); err != nil {
		// Do NOT log.Fatalf: a non-zero exit blocks the containerd shim from
		// unmounting the /dev rbind, which wedges shutdown for ~20 minutes.
		// Un-exported pools are recovered next boot by `zpool import -fal`.
		log.Printf("zfs-service: zpool export did not finish within %s: %v; leaving pools for boot-time import\n", timeout, err)
	}
}
