# VDUSE Extension for Talos Linux

This extension provides VDUSE (vDPA Device in Userspace) kernel modules for Talos Linux.

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

Enable the VDUSE modules in your Talos machine configuration:

```yaml
machine:
  kernel:
    modules:
      - name: vdpa
      - name: vduse
      - name: virtio_vdpa
```

## Verification

Verify the modules are loaded:

```bash
# Check if modules are loaded
talosctl -n <node-ip> read /proc/modules | grep -E "(vdpa|vduse|virtio_vdpa)"

# Check if VDUSE device is available
talosctl -n <node-ip> ls /dev/ | grep vduse
```

## Requirements

- Talos Linux >= v1.8.0 (Linux kernel 6.18+)
- Compatible with both AMD64 and ARM64 architectures
