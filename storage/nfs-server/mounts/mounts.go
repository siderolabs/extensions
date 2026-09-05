// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package mounts prepares the kernel filesystems required by the NFS server.
package mounts

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"syscall"
)

// Mounter provides the filesystem operations required by Ensure.
type Mounter interface {
	MkdirAll(path string, perm fs.FileMode) error
	Mount(source, target, fstype string, flags uintptr, data string) error
}

// OSMounter performs mount operations using the host operating system.
type OSMounter struct{}

// MkdirAll creates a directory and any missing parents.
func (OSMounter) MkdirAll(path string, perm fs.FileMode) error {
	return os.MkdirAll(path, perm)
}

// Mount mounts a filesystem.
func (OSMounter) Mount(source, target, fstype string, flags uintptr, data string) error {
	return syscall.Mount(source, target, fstype, flags, data)
}

// Ensure mounts the kernel nfsd and rpc_pipefs filesystems idempotently.
func Ensure(mounter Mounter) error {
	mounts := []struct {
		source string
		target string
		fstype string
	}{
		{source: "nfsd", target: "/proc/fs/nfsd", fstype: "nfsd"},
		{source: "rpc_pipefs", target: "/var/lib/nfs/rpc_pipefs", fstype: "rpc_pipefs"},
	}

	for _, mount := range mounts {
		if err := mounter.MkdirAll(mount.target, 0o755); err != nil {
			return fmt.Errorf("create %s mountpoint: %w", mount.fstype, err)
		}

		if err := mounter.Mount(mount.source, mount.target, mount.fstype, 0, ""); err != nil && !errors.Is(err, syscall.EBUSY) {
			return fmt.Errorf("mount %s at %s: %w", mount.fstype, mount.target, err)
		}
	}

	return nil
}
