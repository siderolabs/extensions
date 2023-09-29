# gVisor extension

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

gVisor requires unprivileged user namespace creation, so Talos default setting
should be overridden:

```yaml
machine:
  sysctls:
    user.max_user_namespaces: "11255"
```

> Warning! This disables [KSPP best practices](https://kernsec.org/wiki/index.php/Kernel_Self_Protection_Project/Recommended_Settings#sysctls) setting.

## Testing

Apply the following manifest to run nginx pod via gVisor:

```yaml
apiVersion: node.k8s.io/v1
kind: RuntimeClass
metadata:
  name: gvisor
handler: runsc
---
apiVersion: v1
kind: Pod
metadata:
  name: nginx-gvisor
spec:
  runtimeClassName: gvisor
  containers:
  - name: nginx
    image: nginx
```

The pod should be up and running:

```bash
$ kubectl get pods
NAME           READY   STATUS    RESTARTS   AGE
nginx-gvisor   1/1     Running   0          40s
```
