# bcache

This extension provides kernel module needed to use bcache. 

Bcache is a Linux kernel block layer cache. It allows one or more fast disk drives such as flash-based solid state drives (SSDs) to act as a cache for one or more slower hard disk drives.

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

Enable the module in Talos machine config:

```yaml
machine:
  kernel:
    modules:
      - name: bcache
```
