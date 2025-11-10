# Realtek RTL8127 10GbE Driver

## Description

This extension provides the Realtek RTL8127 10 Gigabit Ethernet driver for Talos Linux.

## Hardware Support

- **Device ID**: `10ec:8127` (Realtek RTL8127)
- **Alternate ID**: `10ec:0e10`
- **Speed**: 10 Gigabit Ethernet
- **Interface**: PCIe

## Tested On

- NVIDIA DGX Spark (GB10 Grace Blackwell)

## Features

- Full PTP (Precision Time Protocol) support
- RSS (Receive Side Scaling) support  
- VLAN support
- ASPM (Active State Power Management)
- EEE (Energy Efficient Ethernet)

## Installation

Include this extension when generating Talos boot assets:

```bash
docker run --rm -t \
  -v $PWD/_out:/out \
  ghcr.io/siderolabs/imager:v1.11.0 iso \
  --arch arm64 \
  --system-extension-image ghcr.io/siderolabs/realtek-r8127:v1.11.0
```

## Driver Version

- **r8127**: 11.015.00
- **Source**: OpenWRT rtl8127 (https://github.com/openwrt/rtl8127)

