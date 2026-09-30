# NVIDIA Dynamic Boost

This optional extension runs the unmodified `nvidia-powerd` binary supplied by
the installed NVIDIA toolkit extension. A private D-Bus daemon and a service
wrapper provide its runtime on Talos without systemd or a host-wide D-Bus service.

## Requirements

- An amd64 system with a GPU and firmware supporting NVIDIA Dynamic Boost.
- Matching NVIDIA kernel module and `nvidia-container-toolkit` extensions from
  the same driver branch and Talos release. The toolkit supplies `nvidia-powerd`,
  NVIDIA userspace libraries and glibc; this extension does not duplicate them.
- Native CPU MSR access for the Intel path: install the `msr` extension and
  explicitly load the `msr` module.
- The kernel CPUFreq interface and any firmware power-management interfaces
  required by the hardware.

The arm64 image contains extension metadata only. AMD CPU behavior has not been
validated; the initial service definition requires the MSR device on amd64.

The proposed `msr` extension requires a Talos kernel package that supplies the
signed module. Stock Talos currently disables `CONFIG_X86_MSR`, and its kernel
hardening check rejects enabling it without an approved exception. This
extension does not work around that prerequisite.

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).
Select `nvidia-powerd`, `msr`, and the matching NVIDIA kernel and toolkit
extensions together when building the Talos boot image. These are runtime
requirements; Talos does not automatically install other extensions.

For example, with the open LTS driver:

```yaml
customization:
  systemExtensions:
    officialExtensions:
      - siderolabs/nvidia-open-gpu-kernel-modules-lts
      - siderolabs/nvidia-container-toolkit-lts
      - siderolabs/msr
      - siderolabs/nvidia-powerd
```

`msr` and `nvidia-powerd` in this example are proposed extensions, not currently
available selections in the official Image Factory catalog. Use custom
extension images while developing them.

Load MSR in the machine configuration:

```yaml
machine:
  kernel:
    modules:
      - name: msr
```

Use the kernel's write policy required by the stock daemon. A read-only MSR
configuration may prevent its CPU power controller from working; native
application behavior must be validated on the target CPU.

## Behavior

The service waits for `udevd`, NVIDIA control and CPU MSR devices, and the
toolkit's binary and glibc loader. The wrapper exposes NVIDIA and CPU devices
inside the container, starts a private D-Bus instance, then starts `nvidia-powerd`.
The D-Bus socket and syslog relay are private to the container. Daemon syslog
messages appear in Talos service logs.

If either daemon exits, the wrapper stops the other and exits with an error so
Talos can restart the service. Shutdown stops `nvidia-powerd` before D-Bus so
the daemon can restore its power state.

This extension does not change EC settings, force CPU governors, disable runtime
power management, or impose a GPU power target. NVIDIA and the firmware manage
the supported CPU/GPU power budget. Remove any other running `nvidia-powerd`
instance before enabling this service; only one instance should control a node.

## Verification

```console
talosctl -n <IP> service ext-nvidia-powerd
talosctl -n <IP> logs ext-nvidia-powerd
talosctl -n <IP> read /proc/modules | grep '^msr '
talosctl -n <IP> ls /dev/cpu/0
```

Verify Dynamic Boost support in `/proc/driver/nvidia/gpus/<PCI-ID>/power`.
From a GPU workload, measure the enforced GPU limit, GPU power, CPU package
power and temperatures under both GPU-only and combined CPU/GPU load.
A running service alone does not prove Dynamic Boost is working.

## Development checks

The wrapper build runs process-supervision tests. Run the private D-Bus
integration test separately with the repository's usual build arguments:

```console
make docker-nvidia-powerd-runtime-test PLATFORM=linux/amd64
```

This test uses the matching glibc runtime in an isolated filesystem and checks
that the private bus permits the NVIDIA service name while rejecting other names.
It does not start `nvidia-powerd` or claim to validate GPU power management.
