# coretemp-drivers extension

## Installation

Add the extension to your machine config and enable the `coretemp` module.

```yaml
machine:
  install:
    extensions:
      - image: ghcr.io/siderolabs/coretemp:<VERSION>
  kernel:
    modules:
      - name: coretemp
```

## Verifiying

You can verify the modules are enabled by reading the `/proc/modules` where it _should_ show the module is live.

For example:

```
❯ talosctl -n 192.168.42.15  read /proc/modules
coretemp 299008 - - Live 0xffffffffc03ca00
```
