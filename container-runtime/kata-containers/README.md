# kata-containers extension

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

The extension registers two containerd runtime handlers, selected per-workload
with a `RuntimeClass`:

| handler     | hypervisor          | notes                                                                  |
| ----------- | ------------------- | ---------------------------------------------------------------------- |
| `kata`      | Cloud Hypervisor    | fast boot, low overhead                                                |
| `kata-qemu` | QEMU                | on x86-64, exposes the host CPU virt flag to the guest for nested VMs  |

Both handlers share the same guest kernel and rootfs image and differ only in
the VMM binary and its configuration.

## Testing

Apply the following manifest to run nginx pod using Kata Containers:

```yaml
apiVersion: node.k8s.io/v1
kind: RuntimeClass
metadata:
  name: kata
handler: kata
overhead:
    podFixed:
        memory: "130Mi"
        cpu: "250m"
---
apiVersion: v1
kind: Pod
metadata:
  name: nginx-kata
spec:
  runtimeClassName: kata
  containers:
  - name: nginx
    image: nginx
```

The pod should be up and running:

```bash
$ kubectl get pods
NAME           READY   STATUS    RESTARTS   AGE
nginx-kata     1/1     Running   0          40s
```

To use the QEMU hypervisor instead, select the `kata-qemu` handler:

```yaml
apiVersion: node.k8s.io/v1
kind: RuntimeClass
metadata:
  name: kata-qemu
handler: kata-qemu
overhead:
    podFixed:
        memory: "160Mi"
        cpu: "250m"
---
apiVersion: v1
kind: Pod
metadata:
  name: nginx-kata-qemu
spec:
  runtimeClassName: kata-qemu
  containers:
  - name: nginx
    image: nginx
```

On x86-64, confirm the guest sees the CPU virtualization flag (required for
nested VMs; not available on arm64, where the guest boots without EL2):

```bash
$ kubectl exec nginx-kata-qemu -- grep -m1 -o -E 'vmx|svm' /proc/cpuinfo
vmx
```
