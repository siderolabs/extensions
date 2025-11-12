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

and **additionally** the **kernel** modules:
```yaml
kernel:
    modules:
      - name: scsi-transport-iscsi
      - name: libiscsi_tcp
      - name: iscsi_tcp
      - name: scsi_transport_fc
      - name: dm_multipath
      - name: dm_round-robin
```
To configure multipath accordingly you need to apply:
```yaml
apiVersion: v1alpha1
kind: ExtensionServiceConfig
name: multipathd
configFiles:
    - content: |
        defaults {
            user_friendly_names yes
            find_multipaths no
            path_selector "round-robin 0"
        }
      mountPath: /etc/multipath.conf
```

## References

- [lsscsi man page](https://linux.die.net/man/8/lsscsi)
- [Related Trident issue](https://github.com/NetApp/trident/issues/806#issuecomment-2399332314)