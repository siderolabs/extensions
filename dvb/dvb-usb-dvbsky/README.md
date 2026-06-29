# dvb-usb-dvbsky system extension

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

Enable the `dvb_usb_dvbsky` module in Talos machine config to enable the tuner.

```yaml
machine:
  kernel:
    modules:
      - name: dvb_usb_dvbsky
```
