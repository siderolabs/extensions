# DVB m88ds3103 Firmware Extension

This system extension provides the `dvb-demod-m88ds3103.fw` firmware file, primarily for DVB-S/S2 PCIe cards like the DVBSky S952.

## Prerequisites

*   This extension only provides the firmware. It is intended to be used alongside an existing DVB driver extension (e.g., `dvb-cx23885` or a similar extension) that supplies the necessary kernel modules and drivers for your DVB hardware.
Ensure that a compatible DVB driver extension is installed and configured on your Talos system for this firmware to be utilized.

## Usage

Once installed, this extension performs the following:

1.  Makes the `dvb-demod-m88ds3103.fw` firmware available in the standard firmware path (`/usr/lib/firmware/`) for DVB drivers to load.
2.  Configures the `cx23885` driver for IR support by creating `/etc/modprobe.d/cx23885.conf` with the option `enable_885_ir=1`. This is beneficial if you are using the `dvb-cx23885` extension and your hardware supports IR.

## Firmware Source

The `dvb-demod-m88ds3103.fw` firmware is sourced from the [OpenELEC dvb-firmware project](https://github.com/OpenELEC/dvb-firmware/blob/master/firmware/dvb-demod-m88ds3103.fw).

For more general information on the DVBSky S952 card, you can check this guide: https://www.linuxtv.org/wiki/index.php/DVBSky_S952
