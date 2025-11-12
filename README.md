# Talos Linux System Extensions

This repo serves as a central place for publishing extensions to Talos Linux.
Extensions enable additional functionality beyond the default Talos Linux capabilities.
Things like gVisor, GPU support, etc. are good candidates for extensions.

## Using Extensions

Extensions in this repo are published as container images.
These images can be added to Talos Linux [boot asset](https://www.talos.dev/latest/talos-guides/install/boot-assets/) to produce a final boot asset containing a base Talos `initramfs` and
a set of [system extensions](https://www.talos.dev/latest/talos-guides/configuration/system-extensions/) appended to it.

The extension image is composed of a `manifest.yaml` file that provides information and compatibility information, as well as a `rootfs` that contains things like compiled binaries that are bind-mounted into the system.

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

### Official Extension Tiers

Talos Linux provides several official system extensions, which are split into the following
tiers based on support level:

| Tier | :green_square: core | :yellow_square: extra | :white_large_square: contrib |
| --- | --- | --- | --- |
| Description | Extensions fully supported by Sidero Labs | Some level of support, might vary per extension | Supported by the community |
| Supported by Sidero Labs | 🟢 | ✔️ (best effort) | ❌ |
| Support Channel | GitHub [Issues](https://github.com/siderolabs/extensions/issues), [Discussions](https://github.com/siderolabs/extensions/discussions), [Sidero Labs commercial support](https://www.siderolabs.com/support/) | GitHub [Issues](https://github.com/siderolabs/extensions/issues) and [Discussions](https://github.com/siderolabs/extensions/discussions) | GitHub [Discussions in “contrib” section](https://github.com/siderolabs/extensions/discussions/categories/contrib) |
| Updates managed by Sidero Labs | 🟢 | 🟢 | ✔️ (best effort) |
| Documentation | 🟢 | ✔️ (best effort) | ❌ |
| Automated tests | 🟢 (or no automated tests required, e.g. firmware) | ✔️ (best effort) | ❌ |
| SBOMs | 🟢 (or not required, e.g. firmware) | ✔️ (best effort) | ❌ (community might provide some, but not required) |
| CVE Scan | 🟢 | ✔️ (scan is done, but CVEs don’t block the release) | ❌ |
| Compatibility/Build issues | 🟢 | ✔️ (best effort) | ❌ (extension will be disabled if it fails to build) |

<!-- ### BEGIN GENERATED CONTENT -->
### Container Runtimes

| Name | Tier | Image | Version | Description |
| ---- | ---- | ----- | ------- | ----------- |
| [crun](container-runtime/crun) | :yellow_square: extra | [ghcr.io/siderolabs/crun](https://github.com/siderolabs/extensions/pkgs/container/crun) | `1.25` |  This system extension provides crun using containerd's runtime handler. |
| [ecr-credential-provider](container-runtime/ecr-credential-provider) | :yellow_square: extra | [ghcr.io/siderolabs/ecr-credential-provider](https://github.com/siderolabs/extensions/pkgs/container/ecr-credential-provider) | `v1.34.1` |  This system extension provides a binary which implements Kubelet's CredentialProvider API to authenticate against AWS' Elastic Container Registry and pull images. |
| [gvisor](container-runtime/gvisor) | :green_square: core | [ghcr.io/siderolabs/gvisor](https://github.com/siderolabs/extensions/pkgs/container/gvisor) | `20251103.0` |  This system extension provides gVisor using containerd's runtime handler. |
| [gvisor-debug](container-runtime/gvisor-debug) | :yellow_square: extra | [ghcr.io/siderolabs/gvisor-debug](https://github.com/siderolabs/extensions/pkgs/container/gvisor-debug) | `v1.0.0` |  This system extension enables gVisor debug logging. |
| [kata-containers](container-runtime/kata-containers) | :yellow_square: extra | [ghcr.io/siderolabs/kata-containers](https://github.com/siderolabs/extensions/pkgs/container/kata-containers) | `3.22.0` |  This system extension provides kata-container using containerd's runtime handler. |
| [spin](container-runtime/spin) | :yellow_square: extra | [ghcr.io/siderolabs/spin](https://github.com/siderolabs/extensions/pkgs/container/spin) | `v0.22.0` |  This system extension provides support for spin runtime (WebAssembly) containers. |
| [stargz-snapshotter](container-runtime/stargz-snapshotter) | :green_square: core | [ghcr.io/siderolabs/stargz-snapshotter](https://github.com/siderolabs/extensions/pkgs/container/stargz-snapshotter) | `v0.18.1` |  This system extension provides Stargz Snapshotter using containerd's runtime handler. |
| [wasmedge](container-runtime/wasmedge) | :yellow_square: extra | [ghcr.io/siderolabs/wasmedge](https://github.com/siderolabs/extensions/pkgs/container/wasmedge) | `v0.6.0` |  This system extension provides support for WasmEdge runtime (WebAssembly) containers. |
| [youki](container-runtime/youki) | :white_large_square: contrib | [ghcr.io/siderolabs/youki](https://github.com/siderolabs/extensions/pkgs/container/youki) | `0.5.7` |  This system extension provides youki using containerd's runtime handler. |

### Firmware

| Name | Tier | Image | Version | Description |
| ---- | ---- | ----- | ------- | ----------- |
| [amd-ucode](firmware/amd-ucode) | :green_square: core | [ghcr.io/siderolabs/amd-ucode](https://github.com/siderolabs/extensions/pkgs/container/amd-ucode) | `20251111` |  This system extension provides AMD microcode binaries. |
| [bnx2-bnx2x](firmware/bnx2-bnx2x) | :green_square: core | [ghcr.io/siderolabs/bnx2-bnx2x](https://github.com/siderolabs/extensions/pkgs/container/bnx2-bnx2x) | `20251111` | This system extension provides bnx2 and bnx2x binaries. |
| [chelsio-firmware](firmware/chelsio) | :white_large_square: contrib | [ghcr.io/siderolabs/chelsio-firmware](https://github.com/siderolabs/extensions/pkgs/container/chelsio-firmware) | `20251111` |  This system extension provides Chelsio NIC firmware binaries. |
| [intel-ice-firmware](firmware/intel-ice-firmware) | :green_square: core | [ghcr.io/siderolabs/intel-ice-firmware](https://github.com/siderolabs/extensions/pkgs/container/intel-ice-firmware) | `20251111` |  This system extension provides Intel Ice firmware binaries. |
| [intel-ucode](firmware/intel-ucode) | :green_square: core | [ghcr.io/siderolabs/intel-ucode](https://github.com/siderolabs/extensions/pkgs/container/intel-ucode) | `20251111` |  This system extension provides Intel microcode binaries. |
| [qlogic-firmware](firmware/qlogic-firmware) | :green_square: core | [ghcr.io/siderolabs/qlogic-firmware](https://github.com/siderolabs/extensions/pkgs/container/qlogic-firmware) | `20251111` |  This system extension provides firmware for QLogic devices. |
| [realtek-firmware](firmware/realtek-firmware) | :green_square: core | [ghcr.io/siderolabs/realtek-firmware](https://github.com/siderolabs/extensions/pkgs/container/realtek-firmware) | `20251111` |  This system extension provides realtek firmware binaries. |
| [revpi-firmware](firmware/revpi-firmware) | :white_large_square: contrib | [ghcr.io/siderolabs/revpi-firmware](https://github.com/siderolabs/extensions/pkgs/container/revpi-firmware) | `v1.0.0` |  This system extension provides tools e.g. udev rules for the RevolutionPi platform. |

### Direct Rendering Manager (DRM)

| Name | Tier | Image | Version | Description |
| ---- | ---- | ----- | ------- | ----------- |
| [amdgpu](drm/amdgpu) | :green_square: core | [ghcr.io/siderolabs/amdgpu](https://github.com/siderolabs/extensions/pkgs/container/amdgpu) | `20251111-VERSION` |  This system extension provides AMDGPU firmware binaries and kernel modules. |
| [i915](drm/i915) | :green_square: core | [ghcr.io/siderolabs/i915](https://github.com/siderolabs/extensions/pkgs/container/i915) | `20251111-VERSION` |  This system extension provides Intel GPU microcode binaries and kernel modules. |
| [panfrost](drm/panfrost) | :white_large_square: contrib | [ghcr.io/siderolabs/panfrost](https://github.com/siderolabs/extensions/pkgs/container/panfrost) | `20251111-VERSION` |  This system extension provides ARM Mali Midgard, Bifrost, and Valhall firmware binaries and kernel modules. |
| [vc4](drm/vc4) | :yellow_square: extra | [ghcr.io/siderolabs/vc4](https://github.com/siderolabs/extensions/pkgs/container/vc4) | `VERSION` |  This system extension provides kernel modules for Broadcom VideoCore GPU. |
| [xe](drm/xe) | :green_square: core | [ghcr.io/siderolabs/xe](https://github.com/siderolabs/extensions/pkgs/container/xe) | `20251111-VERSION` |  This system extension provides Intel GPU microcode binaries and kernel modules. |

### Drivers

| Name | Tier | Image | Version | Description |
| ---- | ---- | ----- | ------- | ----------- |
| [amazon-ena](drivers/amazon-ena) | :green_square: core | [ghcr.io/siderolabs/amazon-ena](https://github.com/siderolabs/extensions/pkgs/container/amazon-ena) | `2.16.0-VERSION` |  This system extension provides Amazon ENA kernel modules built against a specific Talos version. ENA is a networking interface designed to make good use of modern CPU features and system architectures. |
| [chelsio-drivers](drivers/chelsio) | :yellow_square: extra | [ghcr.io/siderolabs/chelsio-drivers](https://github.com/siderolabs/extensions/pkgs/container/chelsio-drivers) | `VERSION` |  This system extension provides Chelsio network drivers. |
| [gasket-driver](drivers/gasket) | :yellow_square: extra | [ghcr.io/siderolabs/gasket-driver](https://github.com/siderolabs/extensions/pkgs/container/gasket-driver) | `5815ee3-VERSION` |  This system extension provides google gasket driver kernel modules built against a specific Talos version. This driver is required for PCIe and M.2 Google Coral accelerators. There are 2 kernel modules ("gasket" and "apex") required to enable this driver. |
| [hailort](drivers/hailort) | :yellow_square: extra | [ghcr.io/siderolabs/hailort](https://github.com/siderolabs/extensions/pkgs/container/hailort) | `4.23.0` |  Driver for HailoRT family of AI hardware (eg. Hailo-8L) and is required for PCIe and M.2 Hailo accelerators. |
| [mei](drivers/mei) | :green_square: core | [ghcr.io/siderolabs/mei](https://github.com/siderolabs/extensions/pkgs/container/mei) | `VERSION` |  This system extension provides Intel Management Engine drivers kernel modules built against a specific Talos version. This driver enables the Intel Management Engine, a prerequisite for Intel Arc discrete GPUs. |
| [tenstorrent](drivers/tenstorrent) | :yellow_square: extra | [ghcr.io/siderolabs/tenstorrent](https://github.com/siderolabs/extensions/pkgs/container/tenstorrent) | `2.5.0` |  Driver for Tenstorrent AI processing hardware |
| [thunderbolt](drivers/thunderbolt) | :yellow_square: extra | [ghcr.io/siderolabs/thunderbolt](https://github.com/siderolabs/extensions/pkgs/container/thunderbolt) | `VERSION` |  This system extension provides Thunderbolt/USB4 drivers kernel modules built against a specific Talos version. It enables support for Thunderbolt/USB4 devices, including those used for networking. WARNING: This extension automatically authorizes all Thunderbolt devices during system boot, which poses potential security risks. Use at your own discretion. |
| [uinput](drivers/uinput) | :yellow_square: extra | [ghcr.io/siderolabs/uinput](https://github.com/siderolabs/extensions/pkgs/container/uinput) | `VERSION` |  This system extension provides the uinput kernel module built against a specific Talos version. This kernel module makes it possible to emulate input devices from userspace. By writing to /dev/uinput (or /dev/input/uinput) device, a process can create a virtual input device with specific capabilities. Once this virtual device is created, the process can send events through it, that will be delivered to userspace and in-kernel consumers. |
| [usb-audio-drivers](drivers/usb-audio-drivers) | :white_large_square: contrib | [ghcr.io/siderolabs/usb-audio-drivers](https://github.com/siderolabs/extensions/pkgs/container/usb-audio-drivers) | `VERSION` |  This system extension provides ALSA USB Audio kernel modules built against a specific Talos version. This enables USB audio interfaces (class-compliant sound cards, headsets, DACs) on Talos Linux. |
| [usb-modem-drivers](drivers/usb-modem) | :yellow_square: extra | [ghcr.io/siderolabs/usb-modem-drivers](https://github.com/siderolabs/extensions/pkgs/container/usb-modem-drivers) | `VERSION` |  This system extension provides USB modem drivers kernel modules built against a specific Talos version. This driver is required for USB modems to function. This extension includes all the drivers needed to operate any USB modem under Linux, but your device might not require all of them. Read your device's docs to learn which drivers you need, or just enable them all as a starting point. |
| [v4l-uvc-drivers](drivers/v4l-uvc) | :yellow_square: extra | [ghcr.io/siderolabs/v4l-uvc-drivers](https://github.com/siderolabs/extensions/pkgs/container/v4l-uvc-drivers) | `VERSION` |  This system extension provides the Video4Linux kernel modules required for USB Video Class devices built against a specific Talos version. This driver enables Video4Linux devices such as webcams. |
| [xdma-driver](drivers/xdma-driver) | :yellow_square: extra | [ghcr.io/siderolabs/xdma-driver](https://github.com/siderolabs/extensions/pkgs/container/xdma-driver) | `aefa9a1-VERSION` |  Xilinx DMA Driver |

### Digital Video Broadcasting (DVB)

| Name | Tier | Image | Version | Description |
| ---- | ---- | ----- | ------- | ----------- |
| [dvb-cx23885](dvb/cx23885) | :white_large_square: contrib | [ghcr.io/siderolabs/dvb-cx23885](https://github.com/siderolabs/extensions/pkgs/container/dvb-cx23885) | `VERSION` |  This system extension provides the dvb kernel modules required for Hauppage WinTV-quadHD PCIe tuner built against a specific Talos version. Includes the firmware required. |
| [dvb-m88ds3103](dvb/dvb-m88ds3103) | :white_large_square: contrib | [ghcr.io/siderolabs/dvb-m88ds3103](https://github.com/siderolabs/extensions/pkgs/container/dvb-m88ds3103) | `VERSION` |  This system extension provides the dvb-demod-m88ds3103.fw firmware for DVB-S/S2 PCIe cards like DVBSky S952. It is intended to be used as a dependency on existing DVB driver extension dvb-cx23885 that provides the necessary kernel modules. |

### Miscellaneous

| Name | Tier | Image | Version | Description |
| ---- | ---- | ----- | ------- | ----------- |
| [binfmt-misc](misc/binfmt-misc) | :yellow_square: extra | [ghcr.io/siderolabs/binfmt-misc](https://github.com/siderolabs/extensions/pkgs/container/binfmt-misc) | `VERSION` |  This system extension provides kernel module driver for binfmt-misc built against a specific Talos version. |
| [glibc](misc/glibc) | :green_square: core | [ghcr.io/siderolabs/glibc](https://github.com/siderolabs/extensions/pkgs/container/glibc) | `2.41` | This system extension provides glibc. |

### Network

| Name | Tier | Image | Version | Description |
| ---- | ---- | ----- | ------- | ----------- |
| [cloudflared](network/cloudflared) | :white_large_square: contrib | [ghcr.io/siderolabs/cloudflared](https://github.com/siderolabs/extensions/pkgs/container/cloudflared) | `2025.11.1` |  Cloudflare Tunnel securely connects resources to Cloudflare without a public IP. A lightweight daemon (cloudflared) creates outbound-only connections to Cloudflare, allowing safe access to services like HTTP, SSH, remote desktops, and other protocols. More info: https://github.com/cloudflare/cloudflared/ |
| [lldpd](network/lldpd) | :yellow_square: extra | [ghcr.io/siderolabs/lldpd](https://github.com/siderolabs/extensions/pkgs/container/lldpd) | `1.0.20` |  LLDP adds a LLDP discovery service to Talos. LLDP cli can be used to interface with the daemon. |
| [nebula](network/nebula) | :white_large_square: contrib | [ghcr.io/siderolabs/nebula](https://github.com/siderolabs/extensions/pkgs/container/nebula) | `1.9.7` |  A scalable overlay networking tool with a focus on performance, simplicity and security |
| [netbird](network/netbird) | :white_large_square: contrib | [ghcr.io/siderolabs/netbird](https://github.com/siderolabs/extensions/pkgs/container/netbird) | `0.59.12` |  NetBird combines a WireGuard®-based overlay network with Zero Trust Network Access, providing a unified open source platform for reliable and secure connectivity. |
| [newt](network/newt) | :white_large_square: contrib | [ghcr.io/siderolabs/newt](https://github.com/siderolabs/extensions/pkgs/container/newt) | `1.6.0` |  Newt is a fully user space WireGuard tunnel client and TCP/UDP proxy, designed to securely expose private resources controlled by Pangolin. By using Newt, you don't need to manage complex WireGuard tunnels and NATing. More info: https://github.com/fosrl/newt |
| [tailscale](network/tailscale) | :yellow_square: extra | [ghcr.io/siderolabs/tailscale](https://github.com/siderolabs/extensions/pkgs/container/tailscale) | `1.90.6` |  Tailscale connects your team's devices and development environments for easy access to remote resources. |
| [zerotier](network/zerotier) | :white_large_square: contrib | [ghcr.io/siderolabs/zerotier](https://github.com/siderolabs/extensions/pkgs/container/zerotier) | `1.16.0` |  Connect your Talos cluster into a zerotier network |

### Storage

| Name | Tier | Image | Version | Description |
| ---- | ---- | ----- | ------- | ----------- |
| [btrfs](storage/btrfs) | :yellow_square: extra | [ghcr.io/siderolabs/btrfs](https://github.com/siderolabs/extensions/pkgs/container/btrfs) | `VERSION` |  This system extension provides kernel module driver for BTRFS built against a specific Talos version. |
| [drbd](storage/drbd) | :yellow_square: extra | [ghcr.io/siderolabs/drbd](https://github.com/siderolabs/extensions/pkgs/container/drbd) | `9.2.15-VERSION` |  This system extension provides kernel module driver for DRBD built against a specific Talos version. |
| [fuse3](storage/fuse3) | :green_square: core | [ghcr.io/siderolabs/fuse3](https://github.com/siderolabs/extensions/pkgs/container/fuse3) | `3.17.4` |  This system extension provides fuse3 functionality. |
| [iscsi-tools](storage/iscsi-tools) | :green_square: core | [ghcr.io/siderolabs/iscsi-tools](https://github.com/siderolabs/extensions/pkgs/container/iscsi-tools) | `v0.2.0` |  This system extension provides iscsi-tools. |
| [mdadm](storage/mdadm) | :white_large_square: contrib | [ghcr.io/siderolabs/mdadm](https://github.com/siderolabs/extensions/pkgs/container/mdadm) | `v4.4` |  This system extension provides mdadm binary. |
| [nfs-utils](storage/nfs-utils) | :white_large_square: contrib | [ghcr.io/siderolabs/nfs-utils](https://github.com/siderolabs/extensions/pkgs/container/nfs-utils) | `v0.1.1` |  This system extension provides rpcbind and rpc.statd for NFSv3 file locking support.  rpcbind is a server that converts RPC program numbers into universal addresses. rpc.statd is the NSM (Network Status Monitor) service daemon that notifies NFS peers of restarts.  These services are required for NFSv3 mounts with file locking support. |
| [nfsd](storage/nfsd) | :yellow_square: extra | [ghcr.io/siderolabs/nfsd](https://github.com/siderolabs/extensions/pkgs/container/nfsd) | `VERSION` |  This system extension provides kernel module driver for NFSD built against a specific Talos version. |
| [nfsrahead](storage/nfsrahead) | :white_large_square: contrib | [ghcr.io/siderolabs/nfsrahead](https://github.com/siderolabs/extensions/pkgs/container/nfsrahead) | `2.8.3` |  This system extension provides nfsrahead, a tool to configure the readahead for NFS mounts. |
| [zfs](storage/zfs) | :yellow_square: extra | [ghcr.io/siderolabs/zfs](https://github.com/siderolabs/extensions/pkgs/container/zfs) | `2.4.0-rc2-VERSION` |  This system extension provides the ZFS kernel module, the ZFS utilities, and a service to import all ZFS pools on start and unmount all pools on stop. |

### Power

| Name | Tier | Image | Version | Description |
| ---- | ---- | ----- | ------- | ----------- |
| [nut-client](power/nut-client) | :white_large_square: contrib | [ghcr.io/siderolabs/nut-client](https://github.com/siderolabs/extensions/pkgs/container/nut-client) | `2.8.4` |  This system extension provides the network-ups-tools upsmon service. |

### Guest Agents

| Name | Tier | Image | Version | Description |
| ---- | ---- | ----- | ------- | ----------- |
| [metal-agent](guest-agents/metal-agent) | :green_square: core | [ghcr.io/siderolabs/metal-agent](https://github.com/siderolabs/extensions/pkgs/container/metal-agent) | `v0.1.3` |  This system extension provides talos-metal-agent |
| [qemu-guest-agent](guest-agents/qemu-guest-agent) | :yellow_square: extra | [ghcr.io/siderolabs/qemu-guest-agent](https://github.com/siderolabs/extensions/pkgs/container/qemu-guest-agent) | `10.1.2` |  This system extension provides the QEMU Guest Agent service. |
| [vmtoolsd-guest-agent](guest-agents/vmtoolsd-guest-agent) | :yellow_square: extra | [ghcr.io/siderolabs/vmtoolsd-guest-agent](https://github.com/siderolabs/extensions/pkgs/container/vmtoolsd-guest-agent) | `v1.4.0` |  This system extension provides talos-vmtoolsd |
| [xen-guest-agent](guest-agents/xen-guest-agent) | :yellow_square: extra | [ghcr.io/siderolabs/xen-guest-agent](https://github.com/siderolabs/extensions/pkgs/container/xen-guest-agent) | `0.4.0-g5c274e6` |  xen-guest-agent communicates information and metrics with the Xen host. |

### NVIDIA GPU

| Name | Tier | Image | Version | Description |
| ---- | ---- | ----- | ------- | ----------- |
| [nonfree-kmod-nvidia-lts](nvidia-gpu/nonfree/kmod-nvidia/lts) | :green_square: core | [ghcr.io/siderolabs/nonfree-kmod-nvidia-lts](https://github.com/siderolabs/extensions/pkgs/container/nonfree-kmod-nvidia-lts) | `580.95.05-VERSION` |  This system extension provides nvidia proprietary kernel modules built against a specific Talos version. |
| [nonfree-kmod-nvidia-production](nvidia-gpu/nonfree/kmod-nvidia/production) | :green_square: core | [ghcr.io/siderolabs/nonfree-kmod-nvidia-production](https://github.com/siderolabs/extensions/pkgs/container/nonfree-kmod-nvidia-production) | `570.195.03-VERSION` |  This system extension provides nvidia proprietary kernel modules built against a specific Talos version. |
| [nvidia-container-toolkit-lts](nvidia-gpu/nvidia-container-toolkit/lts) | :green_square: core | [ghcr.io/siderolabs/nvidia-container-toolkit-lts](https://github.com/siderolabs/extensions/pkgs/container/nvidia-container-toolkit-lts) | `580.95.05-v1.18.0` |  This system extension provides nvidia runtime and it's dependencies using NVIDIA's runtime handler. |
| [nvidia-container-toolkit-production](nvidia-gpu/nvidia-container-toolkit/production) | :green_square: core | [ghcr.io/siderolabs/nvidia-container-toolkit-production](https://github.com/siderolabs/extensions/pkgs/container/nvidia-container-toolkit-production) | `570.195.03-v1.18.0` |  This system extension provides nvidia runtime and it's dependencies using NVIDIA's runtime handler. |
| [nvidia-fabricmanager-lts](nvidia-gpu/nvidia-fabricmanager/lts) | :green_square: core | [ghcr.io/siderolabs/nvidia-fabricmanager-lts](https://github.com/siderolabs/extensions/pkgs/container/nvidia-fabricmanager-lts) | `580.95.05` |  This system extension provides the Nvidia fabricmanager for GPU's that need NVLink support. |
| [nvidia-fabricmanager-production](nvidia-gpu/nvidia-fabricmanager/production) | :green_square: core | [ghcr.io/siderolabs/nvidia-fabricmanager-production](https://github.com/siderolabs/extensions/pkgs/container/nvidia-fabricmanager-production) | `570.195.03` |  This system extension provides the Nvidia fabricmanager for GPU's that need NVLink support. |
| [nvidia-gdrdrv-device](nvidia-gpu/nvidia-gdrdrv-device) | :yellow_square: extra | [ghcr.io/siderolabs/nvidia-gdrdrv-device](https://github.com/siderolabs/extensions/pkgs/container/nvidia-gdrdrv-device) | `v2.5.1` |  This system extension provides NVIDIA GPUDirect RDMA (gdrcopy) device file. The gdrdrv kernel module enables low-latency GPU memory access for RDMA operations. |
| [nvidia-open-gpu-kernel-modules-lts](nvidia-gpu/nvidia-modules/lts) | :green_square: core | [ghcr.io/siderolabs/nvidia-open-gpu-kernel-modules-lts](https://github.com/siderolabs/extensions/pkgs/container/nvidia-open-gpu-kernel-modules-lts) | `580.95.05-VERSION` |  This system extension provides nvidia open source driver kernel modules built against a specific Talos version. |
| [nvidia-open-gpu-kernel-modules-production](nvidia-gpu/nvidia-modules/production) | :green_square: core | [ghcr.io/siderolabs/nvidia-open-gpu-kernel-modules-production](https://github.com/siderolabs/extensions/pkgs/container/nvidia-open-gpu-kernel-modules-production) | `570.195.03-VERSION` |  This system extension provides nvidia open source driver kernel modules built against a specific Talos version. |

### Tools

| Name | Tier | Image | Version | Description |
| ---- | ---- | ----- | ------- | ----------- |
| [ctr](tools/ctr) | :green_square: core | [ghcr.io/siderolabs/ctr](https://github.com/siderolabs/extensions/pkgs/container/ctr) | `v2.1.5` |  This extension provides ctr containerd helper binary |
| [nvme-cli](tools/nvme-cli) | :white_large_square: contrib | [ghcr.io/siderolabs/nvme-cli](https://github.com/siderolabs/extensions/pkgs/container/nvme-cli) | `v2.14` |  This system extension provides the NVMe command line interface. |
| [util-linux-tools](tools/util-linux) | :white_large_square: contrib | [ghcr.io/siderolabs/util-linux-tools](https://github.com/siderolabs/extensions/pkgs/container/util-linux-tools) | `2.41.2` |  This system extension provides a minimal util-linux package. |

<!-- ### END GENERATED CONTENT -->

## Building Extensions

In the current form, building extensions requires the use of our [bldr](https://github.com/siderolabs/bldr) tool.
It is highly recommended to take a look at an existing extension as a template for building your own.
The rough flow should look like the following:

- Create a `manifest.yaml` file that contains information about your system extension. See instructions below for this file.
- Create a `pkg.yaml` file that details the full flow of downloading, building, and installing your application.
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
- `/usr/lib/firmware/`
- `/usr/lib/modules/`
- `/usr/lib/ld-linux-x86-64.so.2`
- `/usr/lib/ld-linux-aarch64.so.1`
- `/usr/bin/ldconfig` (used by NVIDIA Container Toolkit)
- `/usr/lib/udev/rules.d/`
- `/usr/local/`
- `/usr/share/glvnd/`
- `/usr/share/egl/`
- `/etc/vulkan/`

## Dependency Diagram

![Dependency Diagram](/deps.svg)
