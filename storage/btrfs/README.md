# btrfs

This extension provides kernel modules and userspace tools needed to support btrfs filesystems.

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

Enable the module in Talos machine config:

```yaml
machine:
  kernel:
    modules:
      - name: btrfs
```
