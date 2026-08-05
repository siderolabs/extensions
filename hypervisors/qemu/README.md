# QEMU extension

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Included functionality

The extension provides hardware-accelerated QEMU system emulation and disk
image tooling for the host architecture:

- `qemu-system-x86_64` on amd64
- `qemu-system-aarch64` on arm64
- `qemu-img`, `qemu-io`, `qemu-nbd`, and `qemu-storage-daemon`

KVM, NUMA, Linux AIO, and Ceph RBD support are enabled. User-mode emulation and
graphical frontends are disabled.
