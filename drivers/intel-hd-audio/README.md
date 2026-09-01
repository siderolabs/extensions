# Intel HD Audio Drivers

This system extension bundles the necessary audio kernel modules for Talos Linux so that onboard Intel HD Audio interfaces can be used. Only x86_64/amd64 is supported, because this chipset is only available on intel boards.

## Installation

See Installing Extensions: https://github.com/siderolabs/extensions#installing-extensions

## Usage

Enable at least the snd-hda-intel kernel module in the Talos Linux machine configuration. Likely you will also need to list the necessary codecs (see below):

```yaml
machine:
  kernel:
    modules:
      - name: snd-hda-codec-alc882
      - name: snd-hda-codec-intelhdmi
      - name: snd-hda-intel
```

You may encounter the following error in dmesg:
```
snd_hda_intel 0000:00:1f.3: Cannot probe codecs, giving up
```

This happens because the codecs must be resident before the controller is probed. To control this, snd_hda_intel is blacklisted from autoloading by udev at PCI detection time, but loaded later by specifying which modules to load in `machine.kernel.modules`.

Swap the codec modules to match your specific HDA chipset / GPU (see the available codec modules in `files/modules-x86_64.txt`).

Mount `/dev/snd` devices into your privileged pod as needed.

## Finding which codecs to load

If you don't know which codec modules your hardware needs, you can discover them at runtime.

Make sure to load `snd_hda_intel`, then in a privileged container rebind the interface:
```
echo 0000:00:1f.3 > /sys/bus/pci/drivers/snd_hda_intel/unbind
echo 0000:00:1f.3 > /sys/bus/pci/drivers/snd_hda_intel/bind
```
(Replace the PCI address with that of your controller.)

Once the codecs come up, dmesg shows which ones were loaded, with the exception of the HDMI ones:

```
talosctl -n <IP> dmesg | grep snd_hda_codec

snd_hda_codec_alc882 hdaudioC0D0: ALC1150: picked fixup  for PCI SSID 1462:0000
```

List those modules in `machine.kernel.modules` together with `snd-hda-intel`.

For the HDMI codecs that don't show up there, pick the right one for your gpu. (See hdmi section in `files/modules-x86_64.txt`)

## Verifying

Check that modules are loaded:

```
talosctl -n <IP> get modules | grep -E 'snd'
talosctl -n <IP> read /proc/modules | grep -E "snd"
```

Check for sound devices:

```
talosctl -n <IP> ls -l /dev/snd
talosctl read -n <IP> /proc/asound/cards
talosctl read -n <IP> /proc/asound/devices
```

And to check for actual sound, start a privileged container mounting `/dev/snd` and install alsa-utils
```
aplay -l
speaker-test -D hw:0,3 -c 2 -t sine -l 1
```
(Replace hardware address with the right one from the aplay -l list)
