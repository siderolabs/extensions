# NVIDIA GPUDirect RDMA (gdrcopy) Driver
This extension provides the NVIDIA GPUDirect RDMA (gdrcopy) device for Talos Linux.
## Description
gdrcopy provide a low-latency GPU memory kernel-mode driver for NVIDIA GPUs. This enables:
- Low-latency access to GPU memory from CPU
- RDMA transfers to/from GPU memory
- Support for GPUDirect RDMA technology
The extension creates the `/dev/gdrdrv` device required to use gdrcopy. The kernel modules required are already bundled in the nvidia-open-gpu-kernel-modules.
## Installation
See [Installing Extensions](https://www.talos.dev/latest/talos-guides/configuration/system-extensions/) for general information on installing system extensions.
## Usage
After installing this extension, the `gdrdrv` kernel module will can be loaded by adding it to your machine config:
```yaml
machine:
  kernel:
    modules:
      - name: gdrdrv
        parameters:
          - dbg_enabled=0
          - info_enabled=0
          - use_persistent_mapping=1
```
The extension will then create the `/dev/gdrdrv` character device that applications can use to interact with GPU memory.
## Requirements
- NVIDIA GPU with GPUDirect RDMA support
- nvidia-open-gpu-kernel-modules-production Talos extension
