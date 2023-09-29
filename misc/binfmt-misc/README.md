# binfmt_misc

Miscellaneous Binary Format is a capability of the Linux kernel that allows arbitrary executable file formats to be recognized and passed to certain user space applications, such as emulators and virtual machines.

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

Enable the kernel module in Talos machine config:

```yaml
machine:
  kernel:
    modules:
      - name: binfmt_misc
```
