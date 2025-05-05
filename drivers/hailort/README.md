# hailo-driver extension

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

By default, the device will be owned by UID and GID `0` and is only accessible by root.
If you need to change this, you may do this by adding udev rules to your machine configuration like this,
which would change the GID to `44` and give that group read/write permissions.

```yaml
machine:
  udev:
    rules:
      - SUBSYSTEM=="hailo_chardev", MODE="0660", GROUP="44"
```

## Verifiying

You can verify the modules are enabled by reading the `/proc/modules` where it _should_ show the module is live.

For example:

```bash
❯ talosctl -n 192.168.77.73  read /proc/modules
TBD
```

In addition, if you actually have Hailo module installed, you should be able to verify it's presence at `/dev/hailo0`.

For example:

```bash
❯ talosctl -n 192.168.77.73 ls -l /dev/hailo0
TBD
```
