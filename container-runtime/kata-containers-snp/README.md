# kata-containers-snp extension

Kata Containers with AMD SEV-SNP confidential VMs ([Confidential
Containers](https://confidentialcontainers.org/)). Registers two containerd
runtime handlers:

| Handler              | Purpose                                     | Requires            |
| -------------------- | ------------------------------------------- | ------------------- |
| `kata-qemu-snp`      | Confidential pods in SEV-SNP encrypted VMs  | AMD SEV-SNP host    |
| `kata-qemu-coco-dev` | Same guest stack without a TEE (dev/test)   | any KVM-capable host |

This extension is a layer on top of the
[kata-containers](../kata-containers) extension, which provides
`containerd-shim-kata-v2`, `virtiofsd`, and the standard QEMU with its
datadir. This one adds only what SEV-SNP needs: the OVMF build with SNP
support (`AMDSEV.fd`), the confidential guest kernel + dm-verity guest
image, and the two handler configs. SEV-SNP in Kata is QEMU-only: the Cloud
Hypervisor driver rejects it (`SEV-SNP protection is not supported by Cloud
Hypervisor`, `virtcontainers/clh.go`).

Both handlers use guest-pull (`experimental_force_guest_pull`): container
images are pulled and unpacked inside the guest VM, never shared from the
host — for `kata-qemu-snp` this means the host never sees image contents.
Guest-pull via the config (rather than kata-deploy's default nydus
snapshotter) is deliberate: Talos cannot run the nydus-snapshotter host
daemon, and the config flag alone activates guest-pull. Because the image is
pulled and unpacked inside the guest during CreateContainerRequest,
`create_container_timeout` is raised from Kata's 60 s default to 1200 s
(multi-GB images time out otherwise).

The handler configs are generated at build time from the release's own
defaults with exactly those changes (path rewrite, guest-pull, timeout) —
everything else, including the guest image's dm-verity parameters, comes
from the release itself, so a Kata version bump needs no config maintenance.
The build asserts every generated field, so an upstream rename fails the
build instead of silently dropping an edit.

amd64 only.

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

The [kata-containers](../kata-containers) extension must be installed on
the same node — it provides the shim, `virtiofsd`, and the QEMU these
handlers run on. The two extensions share no files and are version-paired
through the catalog's single `KATA_CONTAINERS_VERSION` pin.

The host needs SEV-SNP enabled in firmware (SME/SNP BIOS options) for
`kata-qemu-snp`; `kata-qemu-coco-dev` runs on any KVM host.

The compatibility floor is `>= v1.12.0` because host-side SEV-SNP kernel
support entered the Talos kernel in v1.12 (siderolabs/pkgs#1396) — on older
releases `kata-qemu-snp` guests cannot launch.

## Testing

Apply the following manifest to run a pod inside an SEV-SNP confidential VM:

```yaml
apiVersion: node.k8s.io/v1
kind: RuntimeClass
metadata:
  name: kata-qemu-snp
handler: kata-qemu-snp
overhead:
    podFixed:
        memory: "2Gi"
        cpu: "1"
---
apiVersion: v1
kind: Pod
metadata:
  name: nginx-snp
spec:
  runtimeClassName: kata-qemu-snp
  containers:
  - name: nginx
    image: nginx
```

The pod should be up and running, with SEV-SNP active inside the guest:

```bash
$ kubectl get pods
NAME        READY   STATUS    RESTARTS   AGE
nginx-snp   1/1     Running   0          40s
$ kubectl exec nginx-snp -- dmesg | grep -i sev
[    0.213004] Memory Encryption Features active: AMD SEV SEV-ES SEV-SNP
```

On a host without SEV-SNP, use the `kata-qemu-coco-dev` handler instead (same
guest stack, no attestation or memory encryption).

The overhead values match upstream kata-deploy: SNP guest memory
(`default_memory = 2048`) is fully pinned host RAM.
