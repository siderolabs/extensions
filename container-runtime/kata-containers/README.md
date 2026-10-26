# kata-containers extension

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

### Pod annotations

containerd forwards to Kata only the `io.katacontainers.config.hypervisor.*`
annotations the handler's Kata configuration enables (the same allowlist
kata-deploy uses); for both handlers these are `enable_iommu`, `kernel_params`
and `kernel_verity_params`. Any other `io.katacontainers.*` pod annotation is
dropped, so pods cannot override other runtime or agent settings.

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
