# mdadm

This tool activates the healthy software RAID (mdraid).
It's important to note that Talos doesn't handle software RAID maintenance.
You'll need to create and manage it yourself.

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

### Preparation

First check if the disks have any partitions or superblocks on them.
This can be checked by running the following command:

```bash
talosctl get discoveredvolumes
```

If the output for the `DISCOVERED` field for the disks you want to use is not empty, you need to delete the partitions and superblocks.

```text
❯ talosctl -n 10.5.0.3 get discoveredvolumes
NODE       NAMESPACE   TYPE               ID      VERSION   TYPE        SIZE     DISCOVERED   LABEL       PARTITIONLABEL
10.5.0.3   runtime     DiscoveredVolume   loop0   1         disk        463 kB   squashfs
10.5.0.3   runtime     DiscoveredVolume   loop1   1         disk        82 MB    squashfs
10.5.0.3   runtime     DiscoveredVolume   sda     1         disk        11 GB
10.5.0.3   runtime     DiscoveredVolume   sdb     1         disk        11 GB
10.5.0.3   runtime     DiscoveredVolume   sdc     1         disk        11 GB
10.5.0.3   runtime     DiscoveredVolume   vda     1         disk        6.4 GB   gpt
10.5.0.3   runtime     DiscoveredVolume   vda1    1         partition   105 MB   vfat                     EFI
10.5.0.3   runtime     DiscoveredVolume   vda2    1         partition   1.0 MB                            BIOS
10.5.0.3   runtime     DiscoveredVolume   vda3    1         partition   982 MB   xfs          BOOT        BOOT
10.5.0.3   runtime     DiscoveredVolume   vda4    1         partition   524 kB   talosmeta                META
10.5.0.3   runtime     DiscoveredVolume   vda5    1         partition   92 MB    xfs          STATE       STATE
10.5.0.3   runtime     DiscoveredVolume   vda6    1         partition   5.1 GB   xfs          EPHEMERAL   EPHEMERAL
```

This can be done by creating a debug container and running the following commands:

```bash
kubectl -n kube-system debug -it --profile sysadmin --image=alpine node/<node name>
```

Once in the debug container, you can run the following commands:

```bash
apk add wipefs

# Delete the partitions
wipefs --all /dev/<device>
```

### Create soft raid

Now you are ready to create a software raid.

```bash
# create the debug pod again
kubectl -n kube-system debug -it --profile sysadmin --image=alpine node/<node name>
# you are now in the pod again
# the following command will create the software raid
chroot /host -- mdadm --create /dev/md/<some name> --raid-devices=<amount of devices> --metadata=1.2 --level=<raid level>  <all device paths>
# this is an example usage of the command
# chroot /host -- mdadm --create /dev/md/testmd --raid-devices=4 --metadata=1.2 --level=10  /dev/sda /dev/sdb /dev/sdc /dev/sdd
# after you can exit the container and restart the talos node again
exit
```

### Format the raid

Now you you are ready to format the software raid.

- Get the uuid of the created software raid.

```bash
# copy the output of the following command
talosctl ls /dev/disk/by-id | grep md-uuid | awk '{print $2}'
```

#### Let Talos format software raid

- Mount raid1 to the folder `/var/data`. Talos will create the partition and format it.

```yaml
machine:
  ...
  disks:
    - device: /dev/disk/by-id/<output of the previous command>
      partitions:
        - mountpoint: /var/data
```

- Reboot the talos node:

```bash
talosctl reboot --nodes <ip of the node>
```

- Result will be like:

```bash
# talosctl --nodes $NODE mounts | grep '/dev/md'
192.168.1.1   /dev/md0p1                 10.67      0.11       10.56           1.01%          /var/data
```

#### Manually format the software raid

As an alternative you can format it on your own.

- Start the debug pod again and install the package `e2fsprogs`

```bash
kubectl -n kube-system debug -it --profile sysadmin --image=alpine node/<node name>
# Now you are in the container
apk add e2fsprogs
# check if md127 is there
ls /dev/md127
# format the software raid (example here is ext4)
mkfs.ext4 /dev/md127
# exit the container and restart talosctl
exit
talosctl reboot --nodes <ip of the nodes>
# after the reboot check if the formatted device shows extfs value in DISCOVERED coloumn
talosctl get discoveredvolumes | grep md127
```
