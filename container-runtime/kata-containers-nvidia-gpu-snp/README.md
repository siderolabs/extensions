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
  interface for confidential guests, and rejects anything else with
  `ConfidentialGuest needs IOMMUFD`. Talos enables both from **v1.13.7**
  and **v1.14.0-beta.0** onward (siderolabs/pkgs#1608, backported in
  \#1618); `CONFIG_IOMMUFD=m`, so `iommufd.ko` must also be in the
  module allowlist, as it is from those releases.
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

The operator's sandbox operands run privileged, and Talos enforces
PodSecurity `baseline` on every namespace except `kube-system`, so label
the operator's namespace before installing — otherwise the DaemonSets are
rejected at admission and only the namespace events say why:

```bash
$ kubectl create ns gpu-operator
$ kubectl label ns gpu-operator pod-security.kubernetes.io/enforce=privileged
```

The handler consumes a host CDI spec (kind `nvidia.com/pgpu`) naming the
GPU's IOMMUFD cdev (`/dev/vfio/devices/vfioN`). The NVIDIA GPU Operator
writes it, but only in its Kata sandbox mode, which needs gpu-operator
v26.3.0 or newer: `sandboxWorkloads.mode` and `kataSandboxDevicePlugin`
do not exist before that, and the chart ships no `values.schema.json`, so
`--set sandboxWorkloads.mode=kata` on an older chart is accepted in
silence and does nothing. Deploy with `sandboxWorkloads.enabled=true`
**and `sandboxWorkloads.mode=kata`** — the default is `kubevirt`, which
deploys a different device plugin and leaves no spec behind. Leave
`kataManager` disabled (this extension ships the runtime), set
`vfioManager.enabled=false` (the `PCIDriverRebindConfig` above already
binds the GPU), and label the node
`nvidia.com/gpu.workload.config=vm-passthrough`. The Kata sandbox device
plugin — `kataSandboxDevicePlugin`, not `sandboxDevicePlugin`, which is
the KubeVirt one — then writes `/var/run/cdi/nvidia.com-pgpu.yaml` and
advertises the `nvidia.com/pgpu` resource. On a host without the cdev
interface the plugin falls back to the legacy `/dev/vfio/<group>` nodes,
which Kata rejects for confidential guests.

CDI generation needs plugin v0.0.2 or newer, which is exactly what
v26.3.0 defaults to; later releases default higher. Verify the spec file,
not the resource: the Kata plugin advertises `nvidia.com/pgpu` whether or
not it wrote a spec, so a node with no spec still reports capacity and
fails only at pod start, with
`CDI device injection failed: unresolvable CDI devices nvidia.com/pgpu=...`.

```bash
$ talosctl -n <node-ip> ls /var/run/cdi   # expect nvidia.com-pgpu.yaml
```

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
        memory: "10Gi"
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

The overhead values match upstream kata-deploy, which budgets 10240Mi for
this handler — more than the guest's own `default_memory = 8192`, which
for a confidential guest is fully pinned host RAM.

Guest-pull unpacks the container image into guest tmpfs — size the VM for the
image via pod annotations (e.g.
`io.katacontainers.config.hypervisor.default_memory: "16384"`). The pull
budget is the lesser of kubelet's `runtimeRequestTimeout` and the shipped
`create_container_timeout = 1200`, and kubelet defaults to 2 m, so raise it
too for images that take longer:

```yaml
machine:
  kubelet:
    extraConfig:
      runtimeRequestTimeout: 20m
```

Because the pull happens inside the guest, `imagePullSecrets` from the pod
spec are not used; credentials for a private registry have to reach the
guest through a KBS (`agent.image_registry_auth` / `agent.aa_kbc_params`).
