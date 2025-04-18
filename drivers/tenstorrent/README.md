# tenstorrent extension

This extension provides [Tennstorrent](https://github.com/tenstorrent/tt-kmd) drivers for Talos Linux.

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

```yaml
machine:
  kernel:
    modules:
      - name: tenstorrent
```

## Verifying

You can verify the modules are enabled by reading the `/proc/modules` where it _should_ show the module is live.

For example:

```bash
❯ talosctl read /proc/modules
```

In addition, if you actually have Tenstorrent hardware installed, you should be able to verify the devices that begin with `/dev/tenstorrent/*`.

For example:

```bash
❯ talosctl ls -l /dev/tenstorrent/*
NODE           MODE          UID   GID   SIZE(B)   LASTMOD           LABEL                           NAME
192.168.4.20   drwxr-xr-x    0     0     60        May 20 15:26:39   system_u:object_r:device_t:s0   .
192.168.4.20   Dcrw-rw-rw-   0     0     0         May 20 15:26:40   system_u:object_r:device_t:s0   0
```
