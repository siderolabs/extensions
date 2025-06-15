# DVB m88ds3103 Firmware Extension

This system extension provides the `dvb-demod-m88ds3103.fw` firmware file.

## Prerequisites

This extension only provides the firmware. It is intended to be used alongside an existing DVB driver extension (e.g., `dvb-cx23885` or a similar extension) that supplies the necessary kernel modules and drivers for your DVB hardware.

Ensure that a compatible DVB driver extension is installed and configured on your Talos system for this firmware to be utilized.

## Usage

Once installed, this extension makes the `dvb-demod-m88ds3103.fw` firmware available in the standard firmware path (`/usr/lib/firmware/`) for DVB drivers to load.
