# dvb-pt3 system extension

Kernel modules for the [Earthsoft PT3](https://earthsoft.jp/PT3/) ISDB-T/S PCIe tuner card,
common for terrestrial and satellite TV recording in Japan. The card needs no firmware blob.

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

Enable the `earth-pt3` module in Talos machine config to enable the tuner.

```yaml
machine:
  kernel:
    modules:
      - name: earth-pt3
```

The card then appears as `/dev/dvb/adapter*`. A container that tunes needs the device
exposed to it, for example through a `hostPath` volume on `/dev/dvb`.

Decoding ISDB broadcasts also needs a B-CAS card reader, which is a separate USB device
served by `pcscd` in userspace and is not part of this extension.
