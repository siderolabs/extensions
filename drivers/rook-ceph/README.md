# rook-ceph-drivers extension

## Installation

Add the extension to your machine config and enable the modules you need.


```yaml
machine:
  install:
    extensions:
      - image: ghcr.io/siderolabs/rook-ceph-drivers:<VERSION>
  kernel:
    modules:
      # rook / ceph specific modules
      - name: ceph
      - name: nbd
      - name: rbd
      # LVM required modules for ceph https://rook.io/docs/rook/v1.10/Getting-Started/Prerequisites/prerequisites/#lvm-package
      - name: dm_raid
      - name: dm_integrity
      - name: raid0
      - name: raid1
      - name: raid10
      - name: raid456
```

## Verifiying

You can verify the modules are enabled by reading the `/proc/modules` where it _should_ show the module is live.

For example:

```
❯ talosctl -n 192.168.32.5  read /proc/modules
dm_raid 45056 - - Live 0xffffffffc0399000
raid456 155648 - - Live 0xffffffffc0369000
async_raid6_recov 16384 - - Live 0xffffffffc0362000
async_memcpy 16384 - - Live 0xffffffffc035b000
async_pq 16384 - - Live 0xffffffffc0352000
async_xor 16384 - - Live 0xffffffffc034b000
xor 24576 - - Live 0xffffffffc0340000
async_tx 16384 - - Live 0xffffffffc0337000
raid6_pq 118784 - - Live 0xffffffffc0314000
r8169 94208 - - Live 0xffffffffc02fc000
...
...
```

In addition, you should have all the drivers necessary for LVM raid + RAID since they are a subset of ceph drivers, and you can verify your volumes presence at `/dev/mapper/*`, if you have any present on attatched storage.

For example:

```
❯ talosctl -n 192.168.32.5  ls -l /dev/mapper
NODE             MODE          UID   GID   SIZE(B)   LASTMOD           NAME
192.168.99.136   drwxr-xr-x    0     0     160       Aug 31 10:19:35   .
192.168.99.136   Dcrw-------   0     0     0         Aug 31 10:19:33   control
192.168.99.136   Lrwxrwxrwx    167   167   7         Aug 31 10:20:35   data-data -> ../dm-4
192.168.99.136   Lrwxrwxrwx    0     0     7         Aug 31 10:19:34   data-data_rimage_0 -> ../dm-1
192.168.99.136   Lrwxrwxrwx    0     0     7         Aug 31 10:19:34   data-data_rimage_1 -> ../dm-3
192.168.99.136   Lrwxrwxrwx    0     0     7         Aug 31 10:19:34   data-data_rmeta_0 -> ../dm-0
192.168.99.136   Lrwxrwxrwx    0     0     7         Aug 31 10:19:34   data-data_rmeta_1 -> ../dm-2
```

## Configurtion

Talos currently does not have an OS level method of creation, management, or removal of ceph, raid, nor LVM volumes, so you will
probably want to ignore it on the system level, pass the modems `/dev/` devices to a pod, and manage them from the pod.

However the presence of the modules make it possible to use rook-ceph and configure the rook-ceph cluster as normal, along with
maintain the volumes on the kernel level like synchronising lvm-raid volumes in the background.
