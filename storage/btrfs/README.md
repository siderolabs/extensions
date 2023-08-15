# btrfs

This extension provides kernel modules needed to mount btrfs filesystems.

## Installation

Add the extension to your machine config and enable the modules.

```yaml
machine:
  install:
    extensions:
      - image: ghcr.io/siderolabs/btrfs:<VERSION>
  kernel:
    modules:
      - name: btrfs
```
