# nvidia-tegra-nvgpu extension

This extension provides the GPU kernel modules and the GA10B firmware for NVIDIA Jetson Orin modules (AGX Orin, Orin NX, Orin Nano; Tegra234 with the GA10B GPU).
The modules are built from NVIDIA's Jetson Linux r39.2.1 (JetPack 7.2) sources.

The CUDA userspace libraries are not part of this extension.

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

Load the modules in this order:

```yaml
machine:
  kernel:
    modules:
      - name: ivc_ext
      - name: tegra_hv
      - name: host1x
      - name: host1x_nvhost
      - name: host1x_fence
      - name: nvhwpm
      - name: tegra_drm
      - name: nvmap
      - name: mc_utils
      - name: nvgpu
      - name: governor_pod_scaling
```

`host1x` and `tegra_drm` replace the in-tree modules of the same name.

## Verifying

The modules are loaded:

```bash
$ talosctl read /proc/modules | grep -E 'nvgpu|tegra_drm|host1x'
```

The DRM render node exists:

```bash
$ talosctl ls /dev/dri
```

The kernel reports silicon (not a pre-silicon platform):

```bash
$ talosctl read /sys/devices/soc0/revision
Silicon A01
```
