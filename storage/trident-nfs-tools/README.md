# trident-nfs-tools

This extension provides `mkdir` on the host for **NetApp Trident CSI** NFS volume support on Talos Linux.

## Problem

Trident's `chwrap` mechanism shells out to `mkdir` to create NFS mountpoint directories.
Talos Linux does not ship `mkdir` or any coreutils, causing NFS mounts to fail with:

```
Mkdir failed: exit status 2
error mounting NFS volume ... exit status 32
```

Note: The iSCSI code path in Trident uses Go's native `os.MkdirAll()` and does not
have this problem. Only the NFS code path shells out to the `mkdir` binary.

## What's Included

* **mkdir** — busybox multi-call binary (only depends on musl libc)

## Use Case

To run [trident-operator](https://github.com/NetApp/trident) with NFS, you need this extension
in addition to `nfs-utils` (for NFSv3 locking support):

```yaml
systemExtensions:
  - siderolabs/nfs-utils
  - siderolabs/trident-nfs-tools
```

## Upstream Fix

PR [#1152](https://github.com/NetApp/trident/pull/1152) replaces the `mkdir` shell-out
with Go's native `os.MkdirAll()`. Once merged and released, this extension will no longer
be needed.

## References

- [Trident issue #806](https://github.com/NetApp/trident/issues/806)
- [Trident PR #1152](https://github.com/NetApp/trident/pull/1152)
