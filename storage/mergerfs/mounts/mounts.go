// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Package mounts inspects the mount table for mergerfs mounts and detaches the
// ones left behind by a mergerfs process that died.
package mounts

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"syscall"
)

const (
	mountinfoPath = "/proc/self/mountinfo"

	// Filesystem type mergerfs mounts are recorded as.
	fsType = "fuse.mergerfs"

	// maxStaleLayers bounds DetachStale's loop; each restart of a crashed
	// mergerfs adds at most one layer.
	maxStaleLayers = 64
)

// Mounter provides the filesystem operations required by DetachStale.
type Mounter interface {
	ReadMountinfo() ([]byte, error)
	Unmount(target string, flags int) error
}

// OSMounter performs mount operations using the host operating system.
type OSMounter struct{}

// ReadMountinfo returns the mount table of the calling process.
func (OSMounter) ReadMountinfo() ([]byte, error) {
	return os.ReadFile(mountinfoPath)
}

// Unmount unmounts a filesystem.
func (OSMounter) Unmount(target string, flags int) error {
	return syscall.Unmount(target, flags)
}

// IsMounted reports whether a mergerfs mount is what is visible at mountpoint.
//
// The mount table records the path a mount was made at, so mountpoint has to
// be that path: cleaned and without symlinks. It is never looked up on disk,
// which would hang on a wedged FUSE mount.
func IsMounted(mounter Mounter, mountpoint string) (bool, error) {
	mountinfo, err := mounter.ReadMountinfo()
	if err != nil {
		return false, fmt.Errorf("read mount table: %w", err)
	}

	return IsMergerfsMount(mountinfo, mountpoint), nil
}

// DetachStale lazily detaches the mergerfs mounts stacked at mountpoint (see
// IsMounted for the form it takes) and returns how many there were.
//
// A mergerfs process that dies without unmounting leaves its mount behind,
// answering ENOTCONN; mounting a new instance on top of it would stack a second
// mount. Only the mount visible at mountpoint is considered, and only while it
// is a `fuse.mergerfs` mount: a different filesystem mounted over a stale
// mergerfs mount is neither detached nor looked under.
func DetachStale(mounter Mounter, mountpoint string) (int, error) {
	detached := 0

	for range maxStaleLayers {
		mountinfo, err := mounter.ReadMountinfo()
		if err != nil {
			return detached, fmt.Errorf("read mount table: %w", err)
		}

		if !IsMergerfsMount(mountinfo, mountpoint) {
			return detached, nil
		}

		if err = mounter.Unmount(mountpoint, syscall.MNT_DETACH); err != nil {
			return detached, fmt.Errorf("detach %s: %w", mountpoint, err)
		}

		detached++
	}

	return detached, fmt.Errorf("detach %s: still mounted after %d detaches", mountpoint, detached)
}

// IsMergerfsMount reports whether the mount visible at path is a fuse.mergerfs mount.
//
// mountinfo lists mounts in the order they were made, so of several mounts at
// the same path the last entry is the one on top.
func IsMergerfsMount(mountinfo []byte, path string) bool {
	fstype, ok := TopMount(mountinfo, path)

	return ok && fstype == fsType
}

// TopMount returns the filesystem type of the mount visible at path.
func TopMount(mountinfo []byte, path string) (fstype string, ok bool) {
	for line := range bytes.Lines(mountinfo) {
		// 36 35 98:0 /mnt1 /mnt2 rw,noatime master:1 - fuse.mergerfs pool rw,...
		pre, post, found := strings.Cut(strings.TrimRight(string(line), "\n"), " - ")
		if !found {
			continue
		}

		fields := strings.Fields(pre)
		postFields := strings.Fields(post)

		if len(fields) < 5 || len(postFields) < 1 {
			continue
		}

		if unescape(fields[4]) == path {
			fstype, ok = postFields[0], true
		}
	}

	return fstype, ok
}

// unescape decodes the octal escapes (\040 for space, \011, \012, \134) the kernel
// uses for whitespace and backslashes in mountinfo paths.
func unescape(field string) string {
	if !strings.Contains(field, `\`) {
		return field
	}

	var builder strings.Builder

	for i := 0; i < len(field); i++ {
		if field[i] == '\\' && i+3 < len(field) && isOctal(field[i+1]) && isOctal(field[i+2]) && isOctal(field[i+3]) {
			builder.WriteByte((field[i+1]-'0')<<6 | (field[i+2]-'0')<<3 | (field[i+3] - '0'))
			i += 3

			continue
		}

		builder.WriteByte(field[i])
	}

	return builder.String()
}

func isOctal(c byte) bool {
	return c >= '0' && c <= '7'
}
