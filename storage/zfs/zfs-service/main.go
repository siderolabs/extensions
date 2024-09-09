// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"log"
	"os"
	"os/exec"
	"os/signal"

	"golang.org/x/sys/unix"
)

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
}
