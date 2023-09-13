# binfmt_misc

Miscellaneous Binary Format is a capability of the Linux kernel that allows arbitrary executable file formats to be recognized and passed to certain user space applications, such as emulators and virtual machines.

## Installation

Add the extension to your machine config and enable the modules.

```yaml
machine:
  install:
    extensions:
      - image: ghcr.io/siderolabs/binfmt-misc:<VERSION>
  kernel:
    modules:
      - name: binfmt_misc
```
