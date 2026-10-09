# mergerfs

This extension provides [mergerfs](https://github.com/trapexit/mergerfs), a FUSE union
filesystem that pools several filesystems ("branches") into one mount. The `ext-mergerfs`
host service runs one mergerfs process per config file in `/etc/mergerfs`.

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

Requires Talos v1.14.0 or newer for host-mode extension services. FUSE is built into the
Talos kernel, so no kernel module or other extension is needed.

## Usage

Describe each pool in an `EtcFileConfig` document. mergerfs reads all of its options from
the file as `key=value` lines:

```yaml
apiVersion: v1alpha1
kind: EtcFileConfig
name: mergerfs/media.ini
mode: 0o644
contents: |
  branches=/var/mnt/disk1:/var/mnt/disk2
  mountpoint=/var/mnt/media
  fsname=media
  category.create=pfrd
  cache.files=partial
  # wait for the branches to be mounted before pooling them
  branches-mount-timeout=300
  branches-mount-timeout-fail=true
```

Branches and mountpoints have to live under `/var`, the writable part of the Talos
filesystem: user volumes are mounted at `/var/mnt/<name>`, and ZFS datasets need a
`mountpoint` under `/var`. The mountpoint directory must exist (create it once from a
privileged debug pod) and must not be a symlink. `allow_other` is set automatically since
mergerfs runs as root. Everything else is covered by the
[mergerfs options](https://trapexit.github.io/mergerfs/latest/config/options/).

### Exporting over NFS

A pool can be exported with the `nfs-server` extension like any other directory. mergerfs
recommends `inodecalc=path-hash` and `never-forget-nodes=true` in the pool config, leaving
`lazy-umount-mountpoint` at its default, and an explicit `fsid=` on the export because
FUSE filesystems share a device number. The `mp` export option keeps the export inactive
until the pool is mounted:

```text
/var/mnt/media 10.0.0.0/24(rw,sync,no_subtree_check,mp,fsid=0b6a5c1e-3c5b-4bd8-9e0d-2e8b6f1c7a42)
```

## How It Works

The service reads `/etc/mergerfs` when it starts. When a file is added, changed or removed,
or a mergerfs process exits, it stops every pool and exits, and Talos restarts it five
seconds later with the current configuration; the mount a crashed mergerfs leaves behind is
detached first. A file without `mountpoint=`, or two files with the same mountpoint, keep
the service from starting until fixed, and the log says which.

### Waiting for branches

`ext-mergerfs` starts as soon as `/var` is mounted, which can be before the branches are,
for example before the `zfs` extension has imported its pools. `branches-mount-timeout`
makes mergerfs wait until every branch is a mount of its own, and
`branches-mount-timeout-fail=true` makes it exit rather than mount a partial pool. Branches
that are plain directories on the same filesystem as the mountpoint need a
`.mergerfs.branch` marker file instead.

While waiting, mergerfs tries to run `mount` for each missing branch and aborts when there
is no such command, as on Talos. The service then restarts every five seconds, so the pool
still comes up once the branches are mounted; until then the log shows one
`mergerfs exited: signal: aborted` line per attempt.

### Shutdown

On reboot, shutdown, reset and upgrade, the extension's pre-shutdown hook unmounts every
pool while the services that own the branches (such as `ext-zfs-service`) and the ones
serving the pools (such as `ext-nfs-server`) are still running, so nothing races mergerfs
for the branches. A pool that has not unmounted after 20 seconds is lazily detached. The
hook never fails, since a failed hook aborts the whole sequence.

## Testing

Once the node is up, `talosctl service ext-mergerfs` is `Running` and the pools show up in
the service log and in the mount table:

```bash
$ talosctl logs ext-mergerfs
10.5.0.3: mergerfs-supervisor: media.ini: started mergerfs on /var/mnt/media
$ talosctl mounts | grep /var/mnt/media
10.5.0.3   media   42.95   0.00   42.95   0.00%   /var/mnt/media
```

A pod with a `hostPath` volume on the mountpoint and `mountPropagation: HostToContainer`
sees the files of every branch. After `talosctl reboot` the log shows
`draining: stopping all mounts` before the other services stop, and the pools are mounted
again after boot.

## References

- [mergerfs documentation](https://trapexit.github.io/mergerfs/)
- [mergerfs and NFS](https://trapexit.github.io/mergerfs/latest/remote_filesystems/)
