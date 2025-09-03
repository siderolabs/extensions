# ZFS

This extension allows easy usage of zfs raid with Talos.

## Prerequisites

To use zfs you have to install an image with this extension and add the zfs kernel module to your machine config:

```yaml
machine:
  kernel:
    modules:
      - name: zfs
```

## Setup zfs raid

To create the zfs raid make sure the node with the zfs extension is running:

```bash
talosctl get extensions --nodes <nodes>
```

You should see the `zfs` extension.

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

# repeat the following command for every disk
wipefs --all /dev/<device>
# example command: wipefs --all /dev/sda
```

Next you have to create a pod which has privileged rights and prepare him for creating the zfs raid

```bash
kubectl -n kube-system debug -it --profile sysadmin --image=alpine node/<node name>
```

You are now in the container and ready to create the zfs raid.

### Create zfs raid

Afterwards create the zfs raid

```bash
# create the debug pod again
kubectl -n kube-system debug -it --profile sysadmin --image=alpine node/<node name>
# you are now in the pod again
# the following command will create the software raid

# This example creates a zfs raidz2
# IMPORTANT: you can only use the /var/<some name> mount [read](https://www.talos.dev/v1.8/learn-more/architecture/#the-file-system)
chroot /host zpool create -f \
-o ashift=12 \
-O mountpoint="/var/<your mount path>" \
-O xattr=sa \
-O compression=zstd \
-O acltype=posixacl \
-O atime=off \
hdd \
raidz2 \
/dev/disk/by-id/<id of device 1> \
/dev/disk/by-id/<id of device 2> \
/dev/disk/by-id/<id of device 3> \
raidz2 \
/dev/disk/by-id/<id of device 4> \
/dev/disk/by-id/<id of device 5> \
/dev/disk/by-id/<id of device 6>
```

Now you can exit the container with the command `exit`

Reboot Talos node

```bash
talosctl reboot --nodes <node>
```

## Check zfs raid

After your node is rebooted check if the mount exsists.

```bash
talosctl mounts | grep <mount of zfs raid you specified (/var/...)>
```

You should see the mount.

## Use existing pools

If you have existing pools, they will be automatically detected and used.

However, you may need to enter a debug shell and modify the pool configuration in order for them to work with Talos.
Turn off Samba and NFS share settings.
In addition, some SELinux pool policies are currently incompatible with Talos. If you see errors like this:
```bash
node_name: kern: warning: [2025-09-03T00:34:28.924235026Z]: SELinux: Unable to set superblock options before the security server is initialized
```
You can resolve this by modifying the ZFS pool settings to disable SELinux on the volume:
```bash
kubectl -n kube-system debug -it --profile sysadmin --image=alpine node/<node name>
chroot /host zfs set context=none fscontext=none defcontext=none rootcontext=none pool_name
chroot /host zfs set context=none fscontext=none defcontext=none rootcontext=none pool_name/child_dataset_name
# Review the settings to make sure they were changed correctly
chroot /host zfs get all
```