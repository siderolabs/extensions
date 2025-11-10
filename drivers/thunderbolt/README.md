# thunderbolt-drivers extension

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

Enable the `thunderbolt` module in Talos machine config.
If you need Thunderbolt/USB4 networking,
enable `thunderbolt_net` module as well.

```yaml
machine:
  kernel:
    modules:
      - name: thunderbolt
      - name: thunderbolt_net
```

## Verifiying

You can verify the modules are enabled by reading the `/proc/modules` where it _should_ show the module is live.

For example:

```
❯ talosctl -n 192.168.42.15  read /proc/modules
thunderbolt_net 24576 - - Live 0xffffffffc0414000
thunderbolt 299008 - - Live 0xffffffffc03ca00
```

In addition, if you're using networking, you should be able to verify presence of the network interfaces, checking `/sys/class/net` directory.

For example:

```
❯ talosctl -n 192.168.42.15 ls /sys/class/net/ | grep -E 'NODE|thunderbolt'
NODE            NAME
192.168.42.15   thunderbolt0
192.168.42.15   thunderbolt1
```

You can also verify everything in dmesg:

```
❯ talosctl -n 192.168.42.15 dmesg
# look for lines like these:
10.100.52.1: kern:    info: [2023-07-23T16:47:28.22083266Z]: ACPI: bus type thunderbolt registered
10.100.52.1: kern:    info: [2023-07-23T16:47:30.48512066Z]: thunderbolt 0-0:1.1: new retimer found, vendor=0x8087 device=0x15ee
 SUBSYSTEM=thunderbolt
 DEVICE=+thunderbolt:0-0:1.1
10.100.52.1: kern:    info: [2023-07-23T16:47:32.76328066Z]: thunderbolt 1-0:1.1: new retimer found, vendor=0x8087 device=0x15ee
 SUBSYSTEM=thunderbolt
 DEVICE=+thunderbolt:1-0:1.1
10.100.52.1: kern:    info: [2023-07-23T16:47:37.34770966Z]: thunderbolt 0-1: new host found, vendor=0x8086 device=0x1
 SUBSYSTEM=thunderbolt
 DEVICE=+thunderbolt:0-1
10.100.52.1: kern:    info: [2023-07-23T16:47:37.34917566Z]: thunderbolt 0-1: Intel Corp. talos-node-2
 SUBSYSTEM=thunderbolt
 DEVICE=+thunderbolt:0-1
10.100.52.1: kern:    info: [2023-07-23T16:47:39.74636466Z]: thunderbolt 1-1: new host found, vendor=0x8086 device=0x1
 SUBSYSTEM=thunderbolt
 DEVICE=+thunderbolt:1-1
10.100.52.1: kern:    info: [2023-07-23T16:47:39.74767966Z]: thunderbolt 1-1: Intel Corp. talos-node-3
 SUBSYSTEM=thunderbolt
 DEVICE=+thunderbolt:1-1
```

## Security Warning

This extension automatically authorizes all Thunderbolt devices during system boot, which poses potential security risks. Use at your own discretion.

## Power Management

This extension automatically disabled power savings on all Thunderbolt ports, if you would like to re-enable that use the following udev rule in your machineconfig.

```
ACTION=="add", SUBSYSTEM=="thunderbolt", ATTR{power/control}="auto"
```
