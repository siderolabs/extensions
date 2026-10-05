# hailort extension

This extension provides the Hailo-8 family [HailoRT](https://github.com/hailo-ai/hailort-drivers) PCIe driver (`hailo_pci`) for Talos Linux. It is built from the upstream `hailo8` branch (hailort-drivers v4.x) and supports Hailo-8, Hailo-8L, and Hailo-8R accelerators.

For Hailo-10 / Hailo-15 hardware, use the [`hailort-v5`](v5/README.md) extension instead.

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

```yaml
machine:
  kernel:
    modules:
      - name: hailo_pci
```

## Verifying

You can verify the module is loaded by reading `/proc/modules`:

```bash
$ talosctl read /proc/modules
hailo_pci 135168 0 - Live 0xffffffffc0674000 (O)
```

If a Hailo-8 device is attached, the character device should be present:

```bash
$ talosctl ls -l /dev/hailo0
```
