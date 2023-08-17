# device-mapper

This extension provides kernel modules needed for non-stock device-mapper functionality.

## Installation

Add the extension to your machine config and enable the modules.

```yaml
machine:
  install:
    extensions:
      - image: ghcr.io/siderolabs/device-mapper:<VERSION>
  kernel:
    modules:
      - name: dm-thin-pool
```