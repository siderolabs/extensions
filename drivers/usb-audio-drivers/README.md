# USB Audio Drivers

This system extension bundles ALSA USB Audio kernel modules for Talos Linux so that class-compliant USB audio devices (DACs, headsets, sound cards, MIDI-over-USB) work.

## Installation

See Installing Extensions: https://github.com/siderolabs/extensions#installing-extensions

## Usage

Enable the `snd-usb-audio` module in the Talos Linux machine configuration. Kernel dependencies will auto-load.

```yaml
machine:
  kernel:
    modules:
      - name: snd-usb-audio
```

If you also need raw MIDI over USB, the above will pull in `snd-usbmidi-lib` and `snd-rawmidi` automatically. No extra config is required in most cases.

User space ALSA libraries and tools belong in your workload container, not in the Talos host. Mount `/dev/snd` devices into your pod as needed.

## Verifying

Check that modules are loaded:

```
❯ talosctl -n <IP> read /proc/modules | grep -E "snd|usb"
snd_usb_audio ... Live
snd_usbmidi_lib ... Live
snd_pcm ... Live
snd ... Live
```

Check for sound devices:

```
❯ talosctl -n <IP> ls -l /dev/snd
```

You should also see dmesg lines similar to:

```
❯ talosctl -n <IP> dmesg
kern: info: usbcore: registered new interface driver snd-usb-audio
```

## Notes

- This extension copies a minimal set of kernel modules from the Talos kernel image and runs `depmod` so dependency resolution works at boot.
- For HDMI/DP audio on GPUs, use the relevant GPU driver extension; this package only targets USB audio devices.
