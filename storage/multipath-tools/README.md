# multipath-tools

This extension provides the `multipathd` daemon on the host for handling device-mapper multipathing.
It enables consistent, fault-tolerant access to storage devices that expose multiple I/O paths.

## What's Included

* **multipathd**: Multipath daemon
* **multipath**: Device mapper target autoconfig
* **multipathc**: Interactive client for multipathd
* **mpathpersist**: Manages SCSI persistent reservations on dm multipath devices
* **kpartx**: Create device maps from partition tables.

## Use Case

To run this daemon you need the following **kernel** modules:
```yaml
kernel:
    modules:
      - name: dm_multipath
      - name: dm_round-robin # or dm-queue-length
```
To configure multipath accordingly you need to apply something like:
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
You probably need to use `round-robin` or `queue-length` (if that is supported by your use case) as the default `service-time` is not supported, since `dm-service-time` kernel module is not jet in talos build at the moment.


## References

- [multipath man page](https://linux.die.net/man/8/multipath)
- [multipathd man page](https://linux.die.net/man/8/multipathd)
- [Related Kernel module issue](https://github.com/siderolabs/talos/issues/9515)