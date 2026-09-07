# trident-iscsi-tools

This extension provides basic **Linux** tools like 'cat', 'ls', and other binaries on the host used by the **Trident CSI**, which are apparently not jet bundled in to Trident itself.


## What's Included

* **lsscsi**
* **coreutils**: ls, cat, dd
* **procps**: free, pgrep
* **blockdev**


## Use Case

To run [trident-operator](https://github.com/NetApp/trident) with iSCSI, you need to enable the **extensions**:
* `iscsi-tools`
* `multipath-tools`
* `util-linux` (for blkid)
* `trident-iscsi-tools`

and **additionally** the following **kernel** module configuration documents:

```yaml
apiVersion: v1alpha1
kind: KernelModuleConfig
name: scsi_transport_iscsi
---
apiVersion: v1alpha1
kind: KernelModuleConfig
name: libiscsi_tcp
---
apiVersion: v1alpha1
kind: KernelModuleConfig
name: iscsi_tcp
---
apiVersion: v1alpha1
kind: KernelModuleConfig
name: scsi_transport_fc
---
apiVersion: v1alpha1
kind: KernelModuleConfig
name: dm_multipath
---
apiVersion: v1alpha1
kind: KernelModuleConfig
name: dm_round_robin
```

To configure multipath accordingly you need to apply:

```yaml
apiVersion: v1alpha1
kind: EtcFileConfig
name: multipath.conf
mode: 0o644
contents: |
  defaults {
      user_friendly_names yes
      find_multipaths no
      path_selector "round-robin 0"
  }
```

## References

- [lsscsi man page](https://linux.die.net/man/8/lsscsi)
- [Related Trident issue](https://github.com/NetApp/trident/issues/806#issuecomment-2399332314)
