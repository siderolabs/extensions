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

> Warning! This disables [KSPP best practices](https://kspp.github.io/Recommended_Settings#sysctls) setting.

## Testing

### default

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

### With platform KVM (requires nested virtualization or bare-metal)

Apply the following manifest to run nginx pod via gVisor with kvm platform:

```yaml
apiVersion: node.k8s.io/v1
kind: RuntimeClass
metadata:
  name: gvisor
handler: runsc-kvm
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

**Hint**: You can run both runtime handlers in parallel by naming them different
