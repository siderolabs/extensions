# MediaTek Bluetooth Firmware

This system extension provides the single Linux firmware binary required by
the validated MediaTek MT7922 USB Bluetooth function (`0e8d:7922`). It is kept
separate from `bluetooth-usb-drivers` so the driver extension remains generic
and firmware licensing and provenance remain independently auditable.

The binary and `LICENCE.mediatek` are copied from the pinned Sidero Labs
`linux-firmware` package image used by the matching Talos release. No wireless
LAN firmware or other MediaTek device firmware is included.

Select this extension together with `bluetooth-usb-drivers` when assembling an
installer for MT7922 hardware. Realtek hardware should continue to use the
official `siderolabs/realtek-firmware` extension instead.
