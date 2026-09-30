# MSR extension

Provides the Linux `msr` kernel module on amd64. The arm64 image contains
extension metadata only, following the convention used by other x86-only
module extensions.

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).
Use an extension built from the kernel package matching the target Talos release.

The kernel package must be built with `CONFIG_X86_MSR=m`. The extension copies
the signed module from that package without rebuilding, stripping or signing it
again. It fails to build on amd64 if the module is unavailable.

## Usage

Load the module explicitly in the Talos machine configuration:

```yaml
machine:
  kernel:
    modules:
      - name: msr
```

The kernel creates `/dev/cpu/<CPU>/msr` for each online logical CPU. A workload
using the interface needs access to the corresponding host devices and
`CAP_SYS_RAWIO`. Kernel lockdown and device access policies still apply.

For workloads that only need reads, disable writes with the module parameter:

```yaml
machine:
  kernel:
    modules:
      - name: msr
        parameters:
          - allow_writes=off
```

Raw MSR writes can change CPU power limits and other hardware state. Grant
access only to trusted workloads. This extension does not change the module's
default write policy or relax module signature verification.

The extension does not include `nvidia-powerd`, change CPU or GPU power limits,
or configure a workload. It supplies the native device interface used by
applications that require CPU MSR access.

## Verifying

```console
talosctl -n <IP> read /proc/modules | grep '^msr '
talosctl -n <IP> ls /dev/cpu/0
talosctl -n <IP> dmesg | grep -i msr
```

Confirm that `msr` is loaded and the `msr` character device exists. Application
validation must use registers supported by the specific CPU model.
