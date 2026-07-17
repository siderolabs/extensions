# kata-containers-nvidia-gpu-snp extension

NVIDIA GPU add-on for the [kata-containers-snp](../kata-containers-snp)
extension. Registers one containerd runtime handler:

| Handler                    | Purpose                                                    | Requires                              |
| -------------------------- | ---------------------------------------------------------- | ------------------------------------- |
| `kata-qemu-nvidia-gpu-snp` | Confidential pods in SEV-SNP VMs with NVIDIA GPU VFIO passthrough | AMD SEV-SNP host, NVIDIA CC-capable GPU |

The NVIDIA driver and NVRC (NVIDIA's init) run **inside the encrypted
guest** — the host never loads a GPU driver. The guest kernel modules are
built from NVIDIA's open (MIT/GPLv2) kernel-module tree; the guest image also
contains NVIDIA's redistributable proprietary userspace (same class of content
as the nvidia-container-toolkit extension).

Ships:

- the GPU+SNP Kata configuration (`configuration-qemu-nvidia-gpu-snp.toml`,
  generated at build time from the release's own default with the same
  recipe as kata-containers-snp — path rewrite, guest-pull,
  `create_container_timeout`, though upstream's GPU default already ships
  1200; the dm-verity parameters come from the release itself and every
  generated field is asserted by the build),
- the NVIDIA guest kernel and the NVIDIA confidential guest image (~533 MiB;
  one of the largest objects in the catalog),
- `qemu-system-x86_64-snp-experimental` and its datadir, stripped and
  pruned the same way the kata-containers extension treats its QEMU — this
  is the only handler whose config names the snp-experimental binary, so it
  lives here (the siblings' handlers run on kata-containers' standard
  QEMU),
- the NVIDIA Driver License Agreement
  (`/usr/local/share/doc/NVIDIA-driver/LICENSE`) — the driver license
  requires the agreement text to accompany redistributed driver binaries,
  and the guest image itself ships none (the same mechanism the
  nvidia-container-toolkit extension applies to its host driver stack,
  `NVIDIA_GLX-1.0/LICENSE`).

amd64 only.

## Requirements

- **kata-containers and kata-containers-snp extensions from the same
  release** on the same node — kata-containers provides
  `containerd-shim-kata-v2` and `virtiofsd`; kata-containers-snp provides
  the `AMDSEV.fd` OVMF firmware this handler's config resolves to.
  Extensions cannot declare dependencies on each other; all three are
  version-paired through the shared `KATA_CONTAINERS_VERSION` pin. A
  missing sibling surfaces at pod start: containerd cannot resolve the
  shim, or QEMU cannot load the firmware.
- **Host kernel with `CONFIG_IOMMUFD` (`=y` or `=m`) and
  `CONFIG_VFIO_DEVICE_CDEV=y`** — Kata requires the VFIO IOMMUFD cdev
  interface for confidential guests. This sets the compatibility floor:
  Talos >= v1.13.7, which backported them from the v1.14 kernel (see
  `manifest.yaml.tmpl`; the encoded value is a placeholder until the repo's
  Talos pin allows the real floor).
- **GPU Confidential Computing mode ON** (persistent per-GPU setting, e.g.
  via [NVIDIA gpu-admin-tools](https://github.com/NVIDIA/gpu-admin-tools)).
  A CC-capable GPU with CC mode off makes confidential guests stall until the
  sandbox times out, with no useful error.
- SEV-SNP enabled in host firmware, and IOMMU/ACS isolation for passthrough.

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

Bind the GPU to vfio-pci with `PCIDriverRebindConfig`, keeping a
parameter-free `vfio_pci` entry in `machine.kernel.modules` so the driver
is registered (the rebind writes `driver_override` and pokes
`drivers_probe`, which only works for a loaded driver — and unlike a
`vfio_pci ids=` module parameter, the rebind binds the device no matter
what loaded the module first):

```yaml
machine:
  kernel:
    modules:
      - name: vfio_pci
---
apiVersion: v1alpha1
kind: PCIDriverRebindConfig
name: 0000:c1:00.0 # the GPU's PCI address
targetDriver: vfio-pci
```

The handler consumes a host CDI spec (kind `nvidia.com/pgpu`) naming the
GPU's IOMMUFD cdev (`/dev/vfio/devices/vfioN`). The NVIDIA GPU Operator
provides it: deploy the operator with `sandboxWorkloads.enabled=true`,
leave `kataManager` disabled (this extension ships the runtime), and label
the node `nvidia.com/gpu.workload.config=vm-passthrough` —
its sandbox device plugin writes `/var/run/cdi/nvidia.yaml` (cdev device
nodes on IOMMUFD hosts; the legacy `/dev/vfio/<group>` nodes it falls back
to on older kernels make Kata pick the legacy VFIO backend and fail
confidential guests with `ConfidentialGuest needs IOMMUFD`) and advertises
the `nvidia.com/pgpu` resource. Set `vfioManager.enabled=false` — the
`PCIDriverRebindConfig` above already binds the GPU. Verify with
`kubectl get nodes -o jsonpath='{.items[*].status.allocatable.nvidia\.com/pgpu}'`.

## Testing

Apply the following manifest to run a pod in an SEV-SNP confidential VM
with a passed-through GPU (the GPU Operator setup above must be running):

```yaml
apiVersion: node.k8s.io/v1
kind: RuntimeClass
metadata:
  name: kata-qemu-nvidia-gpu-snp
handler: kata-qemu-nvidia-gpu-snp
overhead:
    podFixed:
        memory: "2Gi"
        cpu: "1"
---
apiVersion: v1
kind: Pod
metadata:
  name: gpu-snp
spec:
  runtimeClassName: kata-qemu-nvidia-gpu-snp
  containers:
  - name: cuda
    image: nvcr.io/nvidia/cuda:12.5.0-base-ubuntu22.04
    command: ["sleep", "infinity"]
    resources:
      limits:
        nvidia.com/pgpu: "1"
```

The GPU should be visible inside the guest:

```bash
$ kubectl exec gpu-snp -- nvidia-smi
```

Guest-pull unpacks the container image into guest tmpfs — size the VM for the
image via pod annotations (e.g.
`io.katacontainers.config.hypervisor.default_memory: "16384"`); the shipped
`create_container_timeout = 1200` covers multi-minute in-guest pulls.
