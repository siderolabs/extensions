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
The extension ships no default configuration. The service waits until you provide `/etc/multipath.conf` via [`EtcFileConfig`](https://www.talos.dev/v1.14/reference/configuration/runtime/etcfileconfig/), which is bind-mounted read-only into the service container:

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

You probably need to use `round-robin` or `queue-length` (if that is supported by your use case) as the default `service-time` is not supported, since the `dm-service-time` kernel module is not yet in the Talos build.


## References

- [multipath man page](https://linux.die.net/man/8/multipath)
- [multipathd man page](https://linux.die.net/man/8/multipathd)
- [Related Kernel module issue](https://github.com/siderolabs/talos/issues/9515)