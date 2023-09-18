# chelsio-drivers extension

## Installation

Add the extension to your machine config.

Provides:

* `cxgb`
* `cxgb3`
* `cxgb4`
* `cxgb4vf`

```yaml
machine:
  install:
    extensions:
      - image: ghcr.io/siderolabs/chelsio-drivers:<VERSION>
```

These modules should be auto-loaded by udev based on the NIC present.
