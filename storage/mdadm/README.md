# mdadm

This tool activates the healthy software RAID (mdraid).
It's important to note that Talos doesn't handle software RAID maintenance.
You'll need to create and manage it yourself.

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

### Preparation

Delete every disk which should be part of the raid.

```bash
# Check that you have selected the correct kubeconfig
# Create a privileged container
kubectl -n kube-system run root-alpine --restart=Never --rm -it  --image alpine --privileged --overrides '{"spec":{"hostPID": true, "hostNetwork": true}}' /bin/sh
# You are now in the container
# Install the needed linux tools
apk add util-linux
# delete every disk which should be part of the raid (repeat the following command for every disk)
sfdisk --delete <path to device file> e.g. /dev/sda
# delete superblock for every disk which should be part of the raid (repeat the following command for every disk)
wipefs --all <path of device file> e.g. /dev/sda
# exit the container
exit
# reboot talos node
talosctl reboot --nodes <ip of the node>
# check after reboot if the disks have nothing in the DISCOVERED section (*NOTE:* if there is something their like gtp wipe the superblocks again for the disk)
# if DISCOVERED is empty you are ready to go
talosctl get discoveredvolumes
> NODE            NAMESPACE   TYPE               ID          VERSION   TYPE        SIZE     DISCOVERED   LABEL       PARTITIONLABEL
> storage         runtime     DiscoveredVolume   sda         1         disk        18 TB
```

### Create soft raid

Now you are ready to create a software raid.

```bash
# create the privileged pod again
kubectl -n kube-system run root-alpine --restart=Never --rm -it  --image alpine --privileged --overrides '{"spec":{"hostPID": true, "hostNetwork": true}}' /bin/sh
# you are now in the pod again
# the following command will create the software raid
nsenter --mount=/proc/1/ns/mnt -- mdadm --create /dev/md/<some name> --raid-devices=<amount of devices> --metadata=1.2 --level=<raid level>  <all device paths>
# this is an example usage of the command
# nsenter --mount=/proc/1/ns/mnt -- mdadm --create /dev/md/testmd --raid-devices=4 --metadata=1.2 --level=10  /dev/sda /dev/sdb /dev/sdc /dev/sdd
# after you can exit the container and restart the talos node again
exit
talosctl reboot --nodes <ip of the node>
```

### Format the raid

Now you you are ready to format the software raid.

* Get the uuid of the created software raid.

```bash
# copy the output of the following command
talosctl ls /dev/disk/by-id | grep md-uuid | awk '{print $2}'
```

#### Let Talos format software raid

* Mount raid1 to the folder `/var/data`. Talos will create the partition and format it.

```yaml
machine:
  ...
  disks:
    - device: /dev/disk/by-id/<output of the previous command>
      partitions:
        - mountpoint: /var/data
```

* Reboot the talos node:

```bash
talosctl reboot --nodes <ip of the node>
```

* Result will be like:

```bash
# talosctl --nodes $NODE mounts | grep '/dev/md'
192.168.1.1   /dev/md0p1                 10.67      0.11       10.56           1.01%          /var/data
```

#### Manually format the software raid

As an alternative you can format it on your own.

* Start the privileged pod again and install the package `e2fsprogs`

```bash
kubectl -n kube-system run r00t-ubuntu --restart=Never --rm -it  --image alpine --privileged --overrides '{"spec":{"hostPID": true, "hostNetwork": true}}' /bin/sh
# Now you are in the container
apk add e2fsprogs
# check if md127 is there
ls /dev/md127
# format the software raid (example here is ext4)
mkfs.ext4 /dev/md127
```
