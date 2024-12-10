# dvb-cx23885 system extension

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

Enable the `cx23885` module in Talos machine config to enable the tuner.

```yaml
machine:
  kernel:
    modules:
      - name: cx23885
```