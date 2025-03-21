# xdma-driver extension

This extension provides [Xilinx PCIe DMA](https://github.com/Xilinx/dma_ip_drivers/tree/master/XDMA/linux-kernel) drivers for Talos Linux.

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

```yaml
machine:
  kernel:
    modules:
      - name: xdma
```

## Verifying

You can verify the modules are enabled by reading the `/proc/modules` where it _should_ show the module is live.

For example:

```bash
❯ talosctl read /proc/modules
```

In addition, if you actually have XDMA module installed, you should be able to verify with the presence of character devices that begin with `/dev/xdma*`.

For example:

```bash
❯ talosctl ls -l /dev/xdma*
NODE           MODE          UID   GID   SIZE(B)   LASTMOD           NAME
```
