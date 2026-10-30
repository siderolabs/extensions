# amd-npu system extension

Kernel module (`amdxdna`) and firmware for the
[AMD XDNA](https://docs.kernel.org/accel/amdxdna/amdnpu.html) NPU found in Ryzen AI
processors (e.g. Strix Point, Strix Halo).

The firmware (`amdnpu/`) comes from `linux-firmware`. This extension is only built for
`amd64`, as the driver depends on `X86_64`.

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

Enable the `amdxdna` module in Talos machine config:

```yaml
machine:
  kernel:
    modules:
      - name: amdxdna
```

The NPU then appears as `/dev/accel/accel0`.

The NPU requires the AMD IOMMU to be enabled: booting with `amd_iommu=off` disables it.
