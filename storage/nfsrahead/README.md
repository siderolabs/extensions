# nfsrahead

This extension provides [nfsrahead](https://man7.org/linux/man-pages/man5/nfsrahead.5.html) from `nfs-utils`, a udev helper that sets `read_ahead_kb` on NFS mounts. A larger readahead window can significantly improve sequential read throughput over NFS.

## How It Works

The extension installs a udev rule that runs `nfsrahead` whenever the kernel registers a backing device for an NFS mount. The tool looks up the readahead for the mount's NFS version in `/etc/nfs.conf` and applies it to the device.

## Configuration

The extension ships no default configuration. Without one, every NFS mount gets a readahead of 128 KiB. Provide `/etc/nfs.conf` via [`EtcFileConfig`](https://docs.siderolabs.com/talos/v1.14/reference/configuration/runtime/etcfileconfig):

```yaml
apiVersion: v1alpha1
kind: EtcFileConfig
name: nfs.conf
mode: 0o644
contents: |
  [nfsrahead]
  nfs=15000
  nfs4=16000
  default=128
```

| Key | Applies to |
|-----|------------|
| `nfs` | NFSv3 mounts |
| `nfs4` | NFSv4 mounts |
| `default` | Mounts with no matching key above |

Values are in KiB. Drop-in files under `/etc/nfs.conf.d/*.conf` are also read, so the section can be split across multiple `EtcFileConfig` documents.

The rule only fires when a mount is created, so configuration changes apply to new mounts; existing mounts keep their readahead until remounted.

## References

- [nfsrahead man page](https://man7.org/linux/man-pages/man5/nfsrahead.5.html)
- [nfs.conf man page](https://man7.org/linux/man-pages/man5/nfs.conf.5.html)
