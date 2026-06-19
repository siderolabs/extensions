# hyperv-guest-agent extension

This system extension provides the Hyper-V Linux guest integration daemons from the
kernel's `tools/hv` tree, run as Talos [extension services](https://docs.siderolabs.com/talos/v1.12/build-and-extend-talos/custom-images-and-development/extension-services):

| Service | Binary | Purpose |
| --- | --- | --- |
| `ext-hyperv-kvp` | `hv_kvp_daemon` | Key-Value Pair exchange — reports guest IPv4/IPv6, hostname, and OS details to the host so `Get-VMNetworkAdapter` is populated |
| `ext-hyperv-vss` | `hv_vss_daemon` | VSS — freezes/thaws the data partition so Hyper-V can take application-consistent checkpoints |

Talos already ships the kernel side (`hv_vmbus`, `hv_utils`); this extension adds the
missing user-space daemons. The daemons are statically linked against the base toolchain
and require no shared libraries at runtime.

## Prerequisites

The daemons attach to `/dev/vmbus/hv_kvp` and `/dev/vmbus/hv_vss`, which are created by the
`hv_utils` kernel module. Make sure the corresponding integration services are enabled on
the VM (on the Hyper-V host):

```powershell
Enable-VMIntegrationService -VMName '<vm-name>' -Name 'Key-Value Pair Exchange'
Enable-VMIntegrationService -VMName '<vm-name>' -Name 'VSS'
```

If `hv_utils` is not auto-loaded on your Talos version, add it via machine config:

```yaml
machine:
  kernel:
    modules:
      - name: hv_utils
```

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Testing

Confirm the extension services are running:

```bash
$ talosctl -n <node> service ext-hyperv-kvp
NODE     <node>
ID       ext-hyperv-kvp
STATE    Running
HEALTH   ?
EVENTS   [Running]: Started task ext-hyperv-kvp ...
```

```bash
$ talosctl -n <node> logs ext-hyperv-kvp
```

On the Hyper-V host, the guest IP should now be reported:

```powershell
Get-VMNetworkAdapter -VMName '<vm-name>' | Select-Object -ExpandProperty IPAddresses
```

## Notes & limitations

- **DNS/DHCP reporting is best-effort.** `hv_kvp_daemon` shells out to the distro helper
  scripts `hv_get_dns_info` / `hv_get_dhcp_info` to report DNS servers and DHCP status to
  the host. Those scripts are not shipped, so those two fields stay empty; IP address,
  hostname, and OS reporting are unaffected (they come from the kernel/sysfs, not the
  helpers).
- **Host → guest IP injection is not supported.** The KVP feature that lets the host push a
  static IP into the guest relies on the same (unshipped) helper scripts and a writeable
  network-config path; Talos networking is managed by machine config instead.
- **VSS freezes the `/var` (ephemeral) partition.** The read-only squashfs root needs no
  freeze. Validate checkpoint consistency against your workload before relying on it.
