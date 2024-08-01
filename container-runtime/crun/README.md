# crun extension

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

In order to create the Wasm workload, a runtimeclass needs to be created.

```yaml
apiVersion: node.k8s.io/v1
kind: RuntimeClass
metadata:
  name: crun
handler: crun
```

## Testing

Apply the following manifest to run nginx pod via crun:

```yaml

---
apiVersion: v1
kind: Pod
metadata:
  name: nginx-crun
spec:
  runtimeClassName: crun
  containers:
  - name: nginx
    image: nginx
```

The pod should be up and running:

```bash
$ kubectl get pods
NAME           READY   STATUS    RESTARTS   AGE
nginx-crun   1/1     Running   0          40s
```
