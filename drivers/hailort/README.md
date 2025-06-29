# hailort extension

This extension provides [HailoRT](https://github.com/hailo-ai/hailort-drivers) drivers for Talos Linux.

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

```yaml
machine:
  kernel:
    modules:
      - name: hailort
```

## Verifiying

You can verify the modules are enabled by reading the `/proc/modules` where it _should_ show the module is live.

For example:

```bash
$ talosctl read /proc/modules
hailo_pci 135168 0 - Live 0xffffffffc0674000 (O)
```

In addition, if you actually have Hailo module installed, you should be able to verify it's presence at `/dev/hailo0`.

For example:
```bash
$ talosctl ls -l /dev/hailo0
```
