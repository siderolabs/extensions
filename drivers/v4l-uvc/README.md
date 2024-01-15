# Video4Linux USB Video Class extension

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

Enable the `uvcvideo` module in Talos machine config to enable `/dev/video` devices from USB devices supporting the USB video device class.

```yaml
machine:
  kernel:
    modules:
      - name: uvcvideo
```

## Verifiying

You can verify the modules are enabled by reading the `/proc/modules` where it _should_ show the module is live.

For example:

```
❯ talosctl -n 192.168.42.15  read /proc/modules
uvcvideo 122880 - - Live 0xffffffffc065f000
videobuf2_vmalloc 16384 - - Live 0xffffffffc063b000
videobuf2_memops 16384 - - Live 0xffffffffc0588000
videobuf2_v4l2 28672 - - Live 0xffffffffc05b3000
videobuf2_common 61440 - - Live 0xffffffffc064f000
videodev 237568 - - Live 0xffffffffc0600000
mc 49152 - - Live 0xffffffffc05f3000
```

In addition, you should be able to verify presence of the video device if the USB device is plugged in, checking `/dev` directory.

For example:

```
❯ talosctl -n 192.168.42.15 ls /dev | grep video
192.168.42.15   video0
192.168.42.15   video1
```

You can also verify everything in dmesg:

```
❯ talosctl -n 192.168.42.15 dmesg
# look for lines like these:
kern:    info: [2024-01-15T19:37:30.689914441Z]: videodev: Linux video capture interface: v2.00
kern:    info: [2024-01-15T19:37:34.222751441Z]: usbcore: registered new interface driver uvcvideo
```
