# cachefilesd

This extension provides [cachefilesd](https://www.kernel.org/doc/html/latest/filesystems/caching/cachefiles.html) - the userspace daemon for FS-Cache, the Linux kernel filesystem caching facility. It ships both the `cachefiles` kernel module and the `cachefilesd` daemon.

FS-Cache allows network filesystems (NFS, Ceph, etc.) to cache data on local disk, dramatically reducing cold-start latency for read-heavy workloads. The cache is transparent: applications see no change in behavior, but repeated reads from the network filesystem are served from local disk instead.

## Usage

The kernel module is not loaded automatically. Add `cachefiles` with `KernelModuleConfig` in your machine config:

```yaml
apiVersion: v1alpha1
kind: KernelModuleConfig
name: cachefiles
```

The `cachefilesd` extension service starts automatically once `/dev/cachefiles` appears (i.e., the module is loaded) and the user has provided a configuration file (see below).

### Configuration

Create a user volume which will serve as the cache location.
Any user volume should work, e.g. using a partition or whole disk. If you want to share cache with `EPHEMERAL`, use directory type volume.

```yaml
apiVersion: v1alpha1
kind: UserVolumeConfig
name: fscache
volumeType: partition
provisioning:
    diskSelector:
        match: disk.transport == "nvme" 
    maxSize: 50GiB
```

The extension ships no default config. The service waits until you provide `/etc/cachefilesd.conf` via [`EtcFileConfig`](https://www.talos.dev/v1.14/reference/configuration/runtime/etcfileconfig/), which is bind-mounted into the service container:

```yaml
apiVersion: v1alpha1
kind: EtcFileConfig
name: cachefilesd.conf
mode: 0o644
contents: |
  dir /var/mnt/fscache
  tag mycache
  brun 10%
  bcull 7%
  bstop 3%
  frun 10%
  fcull 7%
  fstop 3%
```

> Note: The `dir` line must point to a directory on a user volume `fscache`, so the value should be `/var/mnt/fscache` always, but the user volume backing it can be any user volume.

### Enabling fscache on NFS mounts

To cache an NFS export, add `fsc` to the mount options. For a Kubernetes PVC backed by NFS, this is typically done in the CSI driver or storage class configuration. Adjust the `dir` line in the config above if a different cache location is needed.

## Compatibility

Requires a Talos kernel built with `CONFIG_FSCACHE=y` and `CONFIG_NFS_FSCACHE=y` (if using NFS fscache). The `cachefiles` module (`CONFIG_CACHEFILES=m`) is provided by this extension.
