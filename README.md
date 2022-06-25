# Talos Linux System Extensions

This repo serves as a central place for publishing supported extensions to Talos Linux.
Extensions allow for additional functionality on top of the default Talos Linux capabilities.
Things like gVisor, GPU support, etc. are good candidates for extensions.

## Using Extensions

Extensions in this repo are published as container images.
These images can be specified in the Talos Linux machine configuration and, when present, will get extracted and laid down as part of the installation process.
The image is composed of a `manifest.yaml` file that provides information and compatibility information, as well as a `rootfs` that contains things like compiled binaries that are bind mounted into the system.

## Extension Catalog

All system extensions provided by Sidero Labs can be found in the [ghcr.io registry](https://github.com/orgs/siderolabs/packages?tab=packages&q=repo%3Asiderolabs%2Fextensions).

### Container Runtimes

| Name                                | Image                                                                                       | Description                                     | Version Format                     |
| ----------------------------------- | ------------------------------------------------------------------------------------------- | ----------------------------------------------- | ---------------------------------- |
| [gvisor](container-runtime/gvisor/) | [ghcr.io/siderolabs/gvisor](https://github.com/siderolabs/extensions/pkgs/container/gvisor) | [gVisor](https://gvisor.dev/) container runtime | `upstream version`-`talos version` |

### Firmware

| Name                                 | Image                                                                                                 | Description                 | Version Format           |
| ------------------------------------ | ----------------------------------------------------------------------------------------------------- | --------------------------- | ------------------------ |
| [amd-ucode](firmware/amd-ucode/)     | [ghcr.io/siderolabs/amd-ucode](https://github.com/siderolabs/extensions/pkgs/container/amd-ucode)     | AMD CPU microcode updates   | `linux firmware version` |
| [bnx2-bnx2x](firmware/bnx2-bnx2x/)   | [ghcr.io/siderolabs/bnx2-bnx2x](https://github.com/siderolabs/extensions/pkgs/container/bnx2-bnx2x)   | Broadcom NetXtreme firmware | `linux firmware version` |
| [intel-ucode](firmware/intel-ucode/) | [ghcr.io/siderolabs/intel-ucode](https://github.com/siderolabs/extensions/pkgs/container/intel-ucode) | Intel CPU microcode updates | `upstream version`       |

### Storage

| Name                                | Image                                                                                                 | Description      | Version Format |
| ----------------------------------- | ----------------------------------------------------------------------------------------------------- | ---------------- | -------------- |
| [iscsi-tools](storage/iscsi-tools/) | [ghcr.io/siderolabs/iscsi-tools](https://github.com/siderolabs/extensions/pkgs/container/iscsi-tools) | Open iSCSI tools | `v0.1.0`       |

### Non-free

System extensions in this section are not distributed by Sidero Labs, they should be built manually following the build instructions.

| Name                                                          | Description                                                                                                                        | Version Format                     |
| ------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------- |
| [nvidia-container-toolkit](nonfree/nvidia-container-toolkit/) | Tools to run [NVIDIA GPU workloads](https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/overview.html) in containers | `driver version`-`toolkit version` |
| [nvidia-fabricmanager](nonfree/nvidia-fabricmanager/)         | [NVIDIA fabric manager](https://docs.nvidia.com/datacenter/tesla/pdf/fabric-manager-user-guide.pdf) support for GPU workloads      | `driver version`                   |

## Building Extensions

In the current form, building extensions requires the use of our [bldr](https://github.com/siderolabs/bldr) tool.
It is highly recommended to take a look at an existing extensions as a template for building your own.
The rough flow should look like the following:

- Create a `manifest.yaml` file that contains information about your system extension. See instructions below for this file.
- Create a `pkg.yaml` file that details the full flow of downloading, building, installing your application.
- Once you have these, add your extension to the `TARGETS` list in the `Makefile`.
- You can now build your extension using make like `make <extension-name> PLATFORM=linux/amd64`
- If you wish to output the contents of the image and validate your install, you can issue `make local-<extension-name> PLATFORM=linux/amd64 DEST=_out`. The contents will then be present in the `_out` directory.

### Creating `manifest.yaml`

The `manifest.yaml` file should match the following format:

```yaml
version: v1alpha1
metadata:
  name: <extension name>
  version: <version of the package the extension installs>-<version of the extensions repo (tracks with talos version)>
  author: Andrew Rynhard
  description: |
    <detailed description of the extension/package>
  ## The compatibility section is "optional" but highly recommended to specify a Talos version that
  ## has been tested and known working for this extension.
  compatibility:
    talos:
      version: ">= v1.0.0"
```

### Creating `pkg.yaml`

Creating a `pkg.yaml` file is the normal process from bldr.
See instructions [here](https://github.com/siderolabs/bldr#pkgyaml) for details and examples on this format.
Using other existing extensions in this repo for tips is also highly recommended.
One important note is that the final directory tree of the generated package should look like this example from the `gvisor` package:

```bash
├── manifest.yaml
└── rootfs
    ├── etc
    │   └── cri
    │       └── conf.d
    │           └── gvisor.part
    └── usr
        └── local
            └── bin
                ├── containerd-shim-runsc-v1
                └── runsc

```

Note that the `manifest.yaml` file lives at the root, while all installed files live under `/rootfs` with the full tree of where they should live on the eventual Talos Linux install.

### `rootfs` Restrictions

The following restrictions are applied to the contents of the `rootfs` of the system extension:

- no special files (FIFOs, devices, etc.)
- no world-writeable files or directories

Any paths in the `rootfs` should be contained within the following hierarchies:

- `/etc/cri/conf.d/`
- `/lib/firmware/`
- `/usr/etc/udev/rules.d`
- `/usr/local/`

## Dependency Diagram

![Dependency Diagram](/deps.png)
