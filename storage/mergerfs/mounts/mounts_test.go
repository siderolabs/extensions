// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package mounts_test

import (
	"errors"
	"slices"
	"strings"
	"syscall"
	"testing"

	"github.com/siderolabs/extensions/storage/mergerfs/mounts"
)

const mountinfo = `22 1 0:21 / / ro,relatime - squashfs /dev/loop0 ro
98 30 0:52 / /var/mnt/disk1 rw,noatime shared:40 - zfs disk1 rw,xattr,posixacl
120 30 0:61 / /var/mnt/pool rw,relatime shared:55 - fuse.mergerfs pool rw,user_id=0,group_id=0,allow_other
121 30 0:62 / /var/mnt/with\040space rw,relatime shared:56 - fuse.mergerfs spaced rw
122 30 0:63 / /var/mnt/other-fuse rw,relatime shared:57 - fuse.sshfs remote rw
123 30 0:64 / /var/mnt/covered rw,relatime shared:58 - fuse.mergerfs stale rw
124 30 0:65 / /var/mnt/covered rw,relatime shared:59 - tmpfs tmpfs rw
125 30 0:66 / /var/mnt/stacked rw,relatime shared:60 - fuse.mergerfs stale rw
126 30 0:67 / /var/mnt/stacked rw,relatime shared:61 - fuse.mergerfs stale rw
`

type call struct {
	target string
	flags  int
}

type fakeMounter struct {
	mountinfo  string
	readErr    error
	unmounts   []call
	unmountErr error
	// sticky keeps the mount table unchanged after an unmount.
	sticky bool
}

func (m *fakeMounter) ReadMountinfo() ([]byte, error) {
	return []byte(m.mountinfo), m.readErr
}

// Unmount drops the last mountinfo entry for target, like the kernel does.
func (m *fakeMounter) Unmount(target string, flags int) error {
	m.unmounts = append(m.unmounts, call{target: target, flags: flags})

	if m.unmountErr != nil {
		return m.unmountErr
	}

	if m.sticky {
		return nil
	}

	lines := strings.Split(strings.TrimSuffix(m.mountinfo, "\n"), "\n")
	for i := len(lines) - 1; i >= 0; i-- {
		if strings.Fields(lines[i])[4] == target {
			m.mountinfo = strings.Join(slices.Delete(lines, i, i+1), "\n") + "\n"

			break
		}
	}

	return nil
}

func TestIsMergerfsMount(t *testing.T) {
	t.Parallel()

	for path, want := range map[string]bool{
		"/var/mnt/pool":       true,
		"/var/mnt/with space": true,
		"/var/mnt/stacked":    true,
		"/var/mnt/disk1":      false, // not mergerfs: never ours to unmount
		"/var/mnt/other-fuse": false,
		"/var/mnt/covered":    false, // a stale mergerfs mount under something else is left alone
		"/var/mnt/missing":    false,
	} {
		if got := mounts.IsMergerfsMount([]byte(mountinfo), path); got != want {
			t.Errorf("IsMergerfsMount(%q) = %v, want %v", path, got, want)
		}
	}
}

func TestDetachStaleDetachesOnlyMergerfsMounts(t *testing.T) {
	t.Parallel()

	for _, tc := range []struct {
		mountpoint string
		detached   int
	}{
		{mountpoint: "/var/mnt/pool", detached: 1},
		{mountpoint: "/var/mnt/stacked", detached: 2},
		{mountpoint: "/var/mnt/covered", detached: 0},
		{mountpoint: "/var/mnt/disk1", detached: 0},
		{mountpoint: "/var/mnt/missing", detached: 0},
	} {
		mounter := &fakeMounter{mountinfo: mountinfo}

		detached, err := mounts.DetachStale(mounter, tc.mountpoint)
		if err != nil {
			t.Fatalf("DetachStale(%q) error = %v", tc.mountpoint, err)
		}

		if detached != tc.detached {
			t.Fatalf("DetachStale(%q) = %v, want %v", tc.mountpoint, detached, tc.detached)
		}

		var want []call
		for range tc.detached {
			want = append(want, call{target: tc.mountpoint, flags: syscall.MNT_DETACH})
		}

		if !slices.Equal(mounter.unmounts, want) {
			t.Fatalf("DetachStale(%q) unmounts = %v, want %v", tc.mountpoint, mounter.unmounts, want)
		}

		if mounts.IsMergerfsMount([]byte(mounter.mountinfo), tc.mountpoint) {
			t.Fatalf("DetachStale(%q) left a mergerfs mount on top", tc.mountpoint)
		}
	}
}

func TestDetachStaleReturnsErrors(t *testing.T) {
	t.Parallel()

	readErr := errors.New("read failed")
	if _, err := mounts.DetachStale(&fakeMounter{readErr: readErr}, "/var/mnt/pool"); !errors.Is(err, readErr) {
		t.Fatalf("DetachStale() error = %v, want %v", err, readErr)
	}

	if _, err := mounts.DetachStale(&fakeMounter{mountinfo: mountinfo, unmountErr: syscall.EINVAL}, "/var/mnt/pool"); !errors.Is(err, syscall.EINVAL) {
		t.Fatalf("DetachStale() error = %v, want EINVAL", err)
	}

	// An unmount that never takes effect must not loop forever.
	sticky := &fakeMounter{mountinfo: mountinfo, sticky: true}
	if _, err := mounts.DetachStale(sticky, "/var/mnt/pool"); err == nil {
		t.Fatal("DetachStale() with a sticky mount returned no error")
	}
}

func TestIsMounted(t *testing.T) {
	t.Parallel()

	mounted, err := mounts.IsMounted(&fakeMounter{mountinfo: mountinfo}, "/var/mnt/pool")
	if err != nil || !mounted {
		t.Fatalf("IsMounted(pool) = %v, %v; want true", mounted, err)
	}

	mounted, err = mounts.IsMounted(&fakeMounter{mountinfo: mountinfo}, "/var/mnt/covered")
	if err != nil || mounted {
		t.Fatalf("IsMounted(covered) = %v, %v; want false", mounted, err)
	}
}
