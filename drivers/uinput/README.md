# uinput extension

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

Enable the `uinput` module in Talos machine config.

```yaml
machine:
  kernel:
    modules:
      - name: uinput
```

## Verifiying

You can verify the modules are enabled by  the `/proc/modules` where it _should_ show the module is live.

For example:

```
❯ talosctl -n 192.168.42.15  read /proc/modules
uinput 24576 - - Live 0xffffffffc0414000
```

In addition, the `/dev/uinput` device should be present.

For example:

```
❯ talosctl -n 192.168.42.15 ls /dev/uinput'
NODE            NAME
192.168.42.15   uinput
```
