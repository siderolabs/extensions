# Bluetooth USB Drivers

This system extension provides the Bluetooth core and USB HCI kernel modules
required to use supported USB and integrated USB Bluetooth controllers on
Talos Linux.

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

Install the firmware extension required by the controller alongside this
extension. Realtek controllers use `realtek-firmware`; MediaTek controllers use
`mediatek-bluetooth-firmware`.

## Usage

Enable `btusb` in the Talos machine configuration. Kernel dependencies and the
matching chipset helper module load automatically.

```yaml
machine:
  kernel:
    modules:
      - name: btusb
```

## Verifying

Check that the Bluetooth modules are loaded:

```console
talosctl -n <IP> read /proc/modules | grep -E "bluetooth|btusb|btintel|btrtl|btmtk"
```

Check that the kernel created an HCI device:

```console
talosctl -n <IP> ls /sys/class/bluetooth
```

Kernel initialization and firmware errors are available in the node log:

```console
talosctl -n <IP> dmesg | grep -i bluetooth
```

## Notes

- The extension contains kernel modules only. BlueZ, D-Bus policy and
  application configuration belong in the workload.
- Firmware is provided by separate vendor extensions and is not duplicated in
  this package.
- The workload running BlueZ must use the host network namespace to access the
  host Bluetooth management interface.
