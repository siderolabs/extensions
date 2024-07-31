# mei extension

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

Provides:

* `mei_wdt`
* `mei_txe`
* `mei_gsc`
* `mei_pxp`
* `mei_hdcp`
* `mei_me`
* `mei`

Modules are automatically loaded.

## Verifiying

You can verify the modules are enabled by reading the `/proc/modules` where it _should_ show the modules are live.

For example:

```
❯ talosctl -n 192.168.227.5 read /proc/modules
mei_wdt 12288 - - Live 0xffffffffc030b000
mei_txe 28672 - - Live 0xffffffffc02d0000
mei_gsc 12288 - - Live 0xffffffffc0247000
mei_pxp 12288 - - Live 0xffffffffc02d4000
mei_hdcp 16384 - - Live 0xffffffffc02bd000
mei_me 45056 - - Live 0xffffffffc0267000
mei 131072 - - Live 0xffffffffc0286000
```
