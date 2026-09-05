// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package mounts_test

import (
	"errors"
	"io/fs"
	"syscall"
	"testing"

	"github.com/siderolabs/extensions/storage/nfs-server/mounts"
)

type call struct {
	source string
	target string
	fstype string
}

type fakeMounter struct {
	mkdirs    []string
	mounts    []call
	mountErrs []error
}

func (m *fakeMounter) MkdirAll(path string, _ fs.FileMode) error {
	m.mkdirs = append(m.mkdirs, path)

	return nil
}

func (m *fakeMounter) Mount(source, target, fstype string, _ uintptr, _ string) error {
	m.mounts = append(m.mounts, call{source: source, target: target, fstype: fstype})

	if len(m.mountErrs) == 0 {
		return nil
	}

	err := m.mountErrs[0]
	m.mountErrs = m.mountErrs[1:]

	return err
}

func TestEnsureMountsKernelFilesystemsOnHost(t *testing.T) {
	t.Parallel()

	mounter := &fakeMounter{}

	if err := mounts.Ensure(mounter); err != nil {
		t.Fatalf("Ensure() error = %v", err)
	}

	wantMounts := []call{
		{source: "nfsd", target: "/proc/fs/nfsd", fstype: "nfsd"},
		{source: "rpc_pipefs", target: "/var/lib/nfs/rpc_pipefs", fstype: "rpc_pipefs"},
	}

	if len(mounter.mounts) != len(wantMounts) {
		t.Fatalf("Ensure() mounted %d filesystems, want %d", len(mounter.mounts), len(wantMounts))
	}

	for i := range wantMounts {
		if mounter.mounts[i] != wantMounts[i] {
			t.Errorf("mount %d = %#v, want %#v", i, mounter.mounts[i], wantMounts[i])
		}
	}
}

func TestEnsureAcceptsAlreadyMountedFilesystems(t *testing.T) {
	t.Parallel()

	mounter := &fakeMounter{mountErrs: []error{syscall.EBUSY, syscall.EBUSY}}

	if err := mounts.Ensure(mounter); err != nil {
		t.Fatalf("Ensure() error = %v, want nil for already-mounted filesystems", err)
	}
}

func TestEnsureReturnsMountFailure(t *testing.T) {
	t.Parallel()

	mounter := &fakeMounter{mountErrs: []error{errors.New("mount failed")}}

	if err := mounts.Ensure(mounter); err == nil {
		t.Fatal("Ensure() error = nil, want mount failure")
	}
}
