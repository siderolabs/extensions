# gasket-driver extension

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
      - SUBSYSTEM=="apex", MODE="0660", GROUP="44"
```

## Verifiying

You can verify the modules are enabled by reading the `/proc/modules` where it _should_ show the module is live.

For example:

```bash
❯ talosctl -n 192.168.32.5  read /proc/modules
apex 20480 - - Live 0xffffffffc01c9000 (O)
gasket 94208 - - Live 0xffffffffc01aa000 (O)
```

In addition, if you actually have Coral module installed, you should be able to verify it's presence at `/dev/apex_0`.

For example:

```bash
❯ talosctl -n 192.168.32.5  ls -l /dev/apex_0
NODE           MODE          UID   GID   SIZE(B)   LASTMOD           NAME
192.168.32.5   Dcrw-rw----   0     44    0         Sep 10 18:15:52   apex_0
```
