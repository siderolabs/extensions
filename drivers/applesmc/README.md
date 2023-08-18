# applesmc-drivers extension

## Installation

Add the extension to your machine config and enable the `applesmc` module.

```yaml
machine:
  install:
    extensions:
      - image: ghcr.io/siderolabs/applesmc:<VERSION>
  kernel:
    modules:
      - name: applesmc
```

## Verifiying

You can verify the modules are enabled by reading the `/proc/modules` where it _should_ show the module is live.

For example:

```
❯ talosctl -n 192.168.42.15  read /proc/modules
applesmc 299008 - - Live 0xffffffffc03ca00
```
