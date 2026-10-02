# qaic system extension

Kernel modules (`qaic` and its `mhi` bus dependency) and firmware for
[Qualcomm Cloud AI](https://docs.kernel.org/accel/qaic/index.html) inference accelerators
(e.g. Cloud AI 100).

The firmware (`qcom/aic*/`) comes from `linux-firmware`. All files are required:
the driver loads `sbl.bin`, and the card then requests the runtime images.

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

Enable the `qaic` module in Talos machine config:

```yaml
machine:
  kernel:
    modules:
      - name: qaic
```

Each card function then appears as `/dev/accel/accel*`, and should report `Ready`
in `qaic-util -q` from the Qualcomm Cloud AI SDK.
