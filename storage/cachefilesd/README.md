# cachefilesd

This extension provides [cachefilesd](https://www.kernel.org/doc/html/latest/filesystems/caching/cachefiles.html) - the userspace daemon for FS-Cache, the Linux kernel filesystem caching facility. It ships both the `cachefiles` kernel module and the `cachefilesd` daemon.

FS-Cache allows network filesystems (NFS, Ceph, etc.) to cache data on local disk, dramatically reducing cold-start latency for read-heavy workloads. The cache is transparent: applications see no change in behavior, but repeated reads from the network filesystem are served from local disk instead.

## Usage

The kernel module is not loaded automatically. Add `cachefiles` to `.machine.kernel.modules` in your machine config:

```yaml
machine:
  kernel:
    modules:
      - name: cachefiles
```

The `cachefilesd` extension service starts automatically once `/dev/cachefiles` appears (i.e., the module is loaded). It reads its configuration from `/usr/local/etc/cachefilesd.conf`.

### Enabling fscache on NFS mounts

To cache an NFS export, add `fsc` to the mount options. For a Kubernetes PVC backed by NFS, this is typically done in the CSI driver or storage class configuration.

The default cache directory is `/var/cache/fscache` on the node's ephemeral partition. Adjust the `dir` line in the config file (via Talos machine config extension service overrides) if a different location is needed.

## Compatibility

Requires a Talos kernel built with `CONFIG_FSCACHE=y` and `CONFIG_NFS_FSCACHE=y` (if using NFS fscache). The `cachefiles` module (`CONFIG_CACHEFILES=m`) is provided by this extension.
