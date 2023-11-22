# mdadm

This tool activates the healthy software RAID (mdraid).
It's important to note that Talos doesn't handle software RAID maintenance.
You'll need to create and manage it yourself.

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

Create soft raid1

```shell
# mdadm --create --verbose /dev/md0 --metadata=0.90 --level=1 --raid-devices=2 /dev/sdb /dev/sdc
# mdadm --detail --scan
ARRAY /dev/md/0 metadata=0.90 UUID=86345734:5d4de8db:5a848257:e2917ea5
```

Mount raid1 to the folder `/var/data`.
Talos will create the partition and format it.

```yaml
machine:
    disks:
        - device: /dev/disk/by-id/md-uuid-86345734:5d4de8db:5a848257:e2917ea5
          partitions:
            - mountpoint: /var/data
```

Result will be like:

```shell
# talosctl --nodes $NODE mounts | grep '/dev/md'
192.168.1.1   /dev/md0p1                 10.67      0.11       10.56           1.01%          /var/data
```
