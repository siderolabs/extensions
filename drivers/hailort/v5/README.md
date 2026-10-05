# hailort-v5 extension

This extension provides the Hailo-10 / Hailo-15 family [HailoRT](https://github.com/hailo-ai/hailort-drivers) PCIe driver (`hailo1x_pci`) for Talos Linux. It is built from the upstream `master` branch (hailort-drivers v5.x) and supports Hailo-10H, Hailo-15H, and Hailo-15L accelerators.

For Hailo-8 hardware, use the [`hailort`](../README.md) extension instead.

## Firmware

- Hailo-10H firmware is bundled under `/usr/lib/firmware/hailo/hailo10h/`.
- Hailo-15H / Hailo-15L firmware (`hailo15_fw.bin`, `hailo15l_fw.bin`) is not publicly redistributable; obtain it from the Hailo developer zone and side-load it via a separate firmware extension at `/usr/lib/firmware/hailo/`.

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

```yaml
machine:
  kernel:
    modules:
      - name: hailo1x_pci
```

## Verifying

You can verify the module is loaded by reading `/proc/modules`:

```bash
$ talosctl read /proc/modules
hailo1x_pci 180224 0 - Live 0xffffffffc0674000 (O)
```

If a Hailo-10/15 device is attached, the character device should be present:

```bash
$ talosctl ls -l /dev/hailo0
```
