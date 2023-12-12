# Talos Linux System Extensions

This repo serves as a central place for publishing supported extensions to Talos Linux.
Extensions allow for additional functionality on top of the default Talos Linux capabilities.
Things like gVisor, GPU support, etc. are good candidates for extensions.

## Using Extensions

Extensions in this repo are published as container images.
These images can be added to the the Talos Linux [boot asset](https://www.talos.dev/latest/talos-guides/install/boot-assets/) to produce a final boot asset containing a base Talos `initramfs` and
a set of [system extensions](https://www.talos.dev/latest/talos-guides/configuration/system-extensions/) appended to it.

The extension image is composed of a `manifest.yaml` file that provides information and compatibility information, as well as a `rootfs` that contains things like compiled binaries that are bind mounted into the system.

## Installing Extensions

In order to find a container reference for a system extension compatible with your Talos Linux version, you can use the following command:

```bash
crane export ghcr.io/siderolabs/extensions:v<talos-version> | tar x -O image-digests | grep <extension-name>
```

For example, to find a compatible version of the `gasket-driver` extension for Talos v1.5.3, you can run:

```bash
$ crane export ghcr.io/siderolabs/extensions:v1.5.3 | tar x -O image-digests | grep gasket-driver
ghcr.io/siderolabs/gasket-driver:97aeba58-v1.5.3@sha256:c786edb356edae3b451cb82d5322f94e54ea0710195181b93ae37ccc8e7ba908
```

Please always use the pinned digest when referencing an extension image.

All extensions are signed with Google Accounts OIDC issuer matching `@siderolabs.com` domain, so the image signatures can be verified, for example:

```bash
cosign verify --certificate-identity-regexp '@siderolabs\.com$' --certificate-oidc-issuer https://accounts.google.com ghcr.io/siderolabs/extensions:v1.5.3
cosign verify --certificate-identity-regexp '@siderolabs\.com$' --certificate-oidc-issuer https://accounts.google.com ghcr.io/siderolabs/gasket-driver:97aeba58-v1.5.3@sha256:c786edb356edae3b451cb82d5322f94e54ea0710195181b93ae37ccc8e7ba908
```

## Extension Catalog

### Container Runtimes

| Name                                | Image                                                                                       | Description                                     | Version Format                     |
| ----------------------------------- | ------------------------------------------------------------------------------------------- | ----------------------------------------------- | ---------------------------------- |
| [gvisor](container-runtime/gvisor/) | [ghcr.io/siderolabs/gvisor](https://github.com/siderolabs/extensions/pkgs/container/gvisor) | [gVisor](https://gvisor.dev/) container runtime | `upstream version`-`talos version` |
| [stargz-snapshotter](container-runtime/stargz-snapshotter/) | [ghcr.io/siderolabs/stargz-snapshotter](https://github.com/siderolabs/extensions/pkgs/container/stargz-snapshotter) | [Stargz Snapshotter](https://github.com/containerd/stargz-snapshotter) container runtime | `upstream version`-`talos version` |
| [ecr-credential-provider](container-runtime/ecr-credential-provider) | [ghcr.io/siderolabs/ecr-credential-provider](https://github.com/siderolabs/extensions/pkgs/container/ecr-credential-provider) | [ECR Credential Provider](https://github.com/kubernetes/cloud-provider-aws/tree/master/cmd/ecr-credential-provider) kubelet plugin | `upstream version` |

### Firmware

| Name                                               | Image                                                                                                               | Description                 | Version Format           |
| -------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------- | --------------------------- | ------------------------ |
| [amd-ucode](firmware/amd-ucode/)                   | [ghcr.io/siderolabs/amd-ucode](https://github.com/siderolabs/extensions/pkgs/container/amd-ucode)                   | AMD CPU microcode updates   | `linux firmware version` |
| [amdgpu-firmware](firmware/amdgpu-firmware/)       | [ghcr.io/siderolabs/amdgpu-firmware](https://github.com/siderolabs/extensions/pkgs/container/amdgpu-firmware)       | AMD GPU firmware            | `linux firmware version` |
| [bnx2-bnx2x](firmware/bnx2-bnx2x/)                 | [ghcr.io/siderolabs/bnx2-bnx2x](https://github.com/siderolabs/extensions/pkgs/container/bnx2-bnx2x)                 | Broadcom NetXtreme firmware | `linux firmware version` |
| [chelsio-firmware](firmware/chelsio-firmware/)     | [ghcr.io/siderolabs/chelsio-firmware](https://github.com/siderolabs/extensions/pkgs/container/chelsio-firmware)     | Chelsio NIC firmware        | `linux firmware version` |
| [i915-ucode](firmware/i915-ucode/)                 | [ghcr.io/siderolabs/i915-ucode](https://github.com/siderolabs/extensions/pkgs/container/i915-ucode)                 | Intel GPU firmware          | `linux firmware version` |
| [intel-ice-firmware](firmware/intel-ice-firmware/) | [ghcr.io/siderolabs/intel-ice-firmware](https://github.com/siderolabs/extensions/pkgs/container/intel-ice-firmware) | Intel ICE NIC firmware      | `linux firmware version` |
| [intel-ucode](firmware/intel-ucode/)               | [ghcr.io/siderolabs/intel-ucode](https://github.com/siderolabs/extensions/pkgs/container/intel-ucode)               | Intel CPU microcode updates | `upstream version`       |
| [qlogic-firmware](firmware/qlogic-firmware/)       | [ghcr.io/siderolabs/qlogic-firmware](https://github.com/siderolabs/extensions/pkgs/container/qlogic-firmware)       | Qlogic firmware             | `linux firmware version` |

### Drivers

| Name                                 | Image                                                                                                                                       | Description                          | Version Format                                        |
| ------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------ | ----------------------------------------------------- |
| [chelsio](drivers/chelsio/)          | [ghcr.io/siderolabs/chelsio-drivers](https://github.com/siderolabs/extensions/pkgs/container/chelsio-drivers)                               | Chelsio NIC drivers                  | `talos version`                                       |
| [gasket](drivers/gasket/)            | [ghcr.io/siderolabs/gasket-driver](https://github.com/siderolabs/extensions/pkgs/container/gasket-driver)                                   | Driver for Google Coral PCIe devices | `gasket driver upstream short commit`-`talos version` |
| [nvidia](nvidia-gpu/nvidia-modules/) | [ghcr.io/siderolabs/nvidia-open-gpu-kernel-modules](https://github.com/siderolabs/extensions/pkgs/container/nvidia-open-gpu-kernel-modules) | NVIDIA OSS Driver                    | `nvidia driver upstream version`-`talos version`      |
| [thunderbolt](drivers/thunderbolt/)  | [ghcr.io/siderolabs/thunderbolt](https://github.com/siderolabs/extensions/pkgs/container/thunderbolt)                                       | Thunderbolt drivers                  | `talos version`                                       |
| [usb-modem](drivers/usb-modem/)      | [ghcr.io/siderolabs/usb-modem-drivers](https://github.com/siderolabs/extensions/pkgs/container/usb-modem-drivers)                           | USB Modem drivers                    | `talos version`                                       |

### Miscellaneous

| Name                            | Image                                                                                             | Description                        | Version Format     |
| ------------------------------- | ------------------------------------------------------------------------------------------------- | ---------------------------------- | ------------------ |
| [binfmt-misc](misc/binfmt-misc) | [ghcr.io/siderolabs/binfmt-misc](https://github.com/siderolabs/extensions/pkgs/container/binfmt-misc) | Miscellaneous Binary Format | `talos version`                                       |

### Network

| Name                            | Image                                                                                             | Description                        | Version Format     |
| ------------------------------- | ------------------------------------------------------------------------------------------------- | ---------------------------------- | ------------------ |
| [tailscale](network/tailscale/) | [ghcr.io/siderolabs/tailscale](https://github.com/siderolabs/extensions/pkgs/container/tailscale) | [Tailscale](https://tailscale.com) | `upstream version` |

### Storage

| Name                                | Image                                                                                                 | Description         | Version Format                     |
| ----------------------------------- | ----------------------------------------------------------------------------------------------------- | ------------------- | ---------------------------------- |
| [iscsi-tools](storage/iscsi-tools/) | [ghcr.io/siderolabs/iscsi-tools](https://github.com/siderolabs/extensions/pkgs/container/iscsi-tools) | Open iSCSI tools    | `v0.1.0`                           |
| [mdadm](storage/mdadm/)             | [ghcr.io/siderolabs/mdadm](https://github.com/siderolabs/extensions/pkgs/container/mdadm) | manage MD devices tool | `upstream version`-`talos version`                    |
| [drbd](storage/drbd/)               | [ghcr.io/siderolabs/drbd](https://github.com/siderolabs/extensions/pkgs/container/drbd)               | DRBD driver module  | `upstream version`-`talos version` |
| [zfs](storage/zfs/)                 | [ghcr.io/siderolabs/zfs](https://github.com/siderolabs/extensions/pkgs/container/zfs)                 | ZFS driver module   | `upstream version`-`talos version` |
| [btrfs](storage/btrfs/)             | [ghcr.io/siderolabs/btrfs](https://github.com/siderolabs/extensions/pkgs/container/btrfs)             | BTRFS driver module | `talos version`                    |

### Power

| Name                            | Image                                                                                                     | Description                                                    | Version Format                     |
| ------------------------------- | --------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------- | ---------------------------------- |
| [nut-client](power/nut-client/) | [ghcr.io/siderolabs/nut-client](https://github.com/siderolabs/talos-extensions/pkgs/container/nut-client) | [Network UPS Tools](https://networkupstools.org) upsmon client | `upstream version`-`talos version` |

### Guest Agents

| Name                                                   | Image                                                                                                                     | Description                                                            | Version Format     |
| ------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------- | ------------------ |
| [qemu-guest-agent](guest-agents/qemu-guest-agent/)     | [ghcr.io/siderolabs/qemu-guest-agent](https://github.com/siderolabs/talos-extensions/pkgs/container/qemu-guest-agent)     | [QEMU Guest Agent](https://wiki.qemu.org/Features/GuestAgent)          | `upstream version` |
| [xe-guest-utilities](guest-agents/xe-guest-utilities/) | [ghcr.io/siderolabs/xe-guest-utilities](https://github.com/siderolabs/talos-extensions/pkgs/container/xe-guest-utilities) | [xe-guest-utilities](https://github.com/xenserver/xe-guest-utilitiest) | `upstream version` |

### NVIDIA GPU

| Name                                                             | Description                                                                                                                        | Version Format                     |
| ---------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------- |
| [nvidia-container-toolkit](nvidia-gpu/nvidia-container-toolkit/) | Tools to run [NVIDIA GPU workloads](https://docs.nvidia.com/datacenter/cloud-native/container-toolkit/overview.html) in containers | `driver version`-`toolkit version` |
| [nvidia-fabricmanager](nvidia-gpu/nvidia-fabricmanager/)         | [NVIDIA fabric manager](https://docs.nvidia.com/datacenter/tesla/pdf/fabric-manager-user-guide.pdf) support for GPU workloads      | `driver version`                   |
| [nvidia-open-gpu-kernel-modules](nvidia-gpu/nvidia-modules/)     | NVIDIA driver kernel modules                                                                                                       | `driver version`-`talos version`   |

#### Tools

| Name                                  | Description                        | Version Format  |
| ------------------------------------- | ---------------------------------- | --------------- |
| [util-linux-tools](tools/util-linux/) | Util Linux tools (only fstrim now) | `talos version` |

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
- `/lib/modules/`
- `/lib64/ld-linux-x86-64.so.2`
- `/usr/etc/udev/rules.d/`
- `/usr/local/`
- `/usr/share/glvnd/`
- `/usr/share/egl/`
- `/etc/vulkan/`

## Dependency Diagram

![Dependency Diagram](/deps.png)
