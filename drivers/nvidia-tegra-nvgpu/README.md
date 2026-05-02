# nvidia-tegra-nvgpu

Talos system extension providing NVIDIA GA10B GPU kernel modules for **Jetson Orin NX / Orin Nano** (Tegra234 SoC).

## Hardware

| SoM | SoC | GPU |
|-----|-----|-----|
| Jetson Orin NX 8GB / 16GB | Tegra234 | GA10B (Ampere, 1024 CUDA cores) |
| Jetson Orin Nano 4GB / 8GB | Tegra234 | GA10B subset |

## Modules

| Module | Source | Purpose |
|--------|--------|---------|
| `host1x.ko` | OE4T/linux-nv-oot | Syncpoint allocator with GA10B `ERRATA_SYNCPT_INVALID_ID_0` fix |
| `host1x_fence.ko` | OE4T/linux-nv-oot | DMA fence bridge for GPU/CPU synchronization |
| `nvmap.ko` | OE4T/linux-nv-oot | GPU memory allocator (NVIDIA memory management API) |
| `mc_utils.ko` | OE4T/linux-nv-oot | Memory controller EMC frequency helper |
| `governor_pod_scaling.ko` | OE4T/linux-nv-oot | nvhost_podgov devfreq governor for dynamic GPU frequency scaling |
| `nvhost_ctrl_shim.ko` | in-tree | `/dev/nvhost-ctrl` bridge — provides `SYNCPT_WAITMEX` ioctl for JetPack 6 CUDA runtime |
| `nvgpu.ko` | OE4T/linux-nvgpu | Main GA10B GPU driver (Clang-compatible, kernel 6.18) |

## Prerequisites

This extension requires the following to be present on the node:

1. **nvidia-firmware-ext** — GPU firmware blobs (PMU, GSP, etc.) from JetPack r36.5
2. JetPack 6 BSP (L4T r36.x) flashed to the Jetson carrier board for EEPROM / PMIC firmware

## Talos Configuration

Module load order is critical. Add the following to your Talos machine configuration:

```yaml
machine:
  kernel:
    modules:
      - name: host1x
      - name: host1x_fence
      - name: nvhost_ctrl_shim
      - name: nvmap
      - name: mc_utils
      - name: nvgpu
      - name: governor_pod_scaling
```

## CDI

Container Device Interface (CDI) is managed by the `nvidia-cdi-setup` DaemonSet from the
[sbc-jetson](https://github.com/siderolabs/sbc-jetson) overlay. It generates
`/run/cdi/nvidia-jetson.yaml` at runtime, exposing `/dev/nvgpu/igpu0/*`,
`/dev/nvhost-ctrl`, and `/dev/nvmap` to GPU workloads.

## Performance

Tested on Jetson Orin NX 16GB with Talos v1.13.0, kernel 6.18.24, CUDA 12.6 (JetPack 6.2):

| Model | Quantization | tok/s |
|-------|-------------|-------|
| qwen2.5:0.5b | Q4_K_M | ~60 |
| qwen3:4b | Q4_K_M | ~16 |

## References

- [siderolabs/pkgs#1518](https://github.com/siderolabs/pkgs/pull/1518) — kernel package that builds these modules
- [OE4T/linux-nvgpu](https://github.com/OE4T/linux-nvgpu) — GA10B GPU driver (patches-r36 branch)
- [OE4T/linux-nv-oot](https://github.com/OE4T/linux-nv-oot) — NVIDIA OOT modules (patches-r36.5 branch)
