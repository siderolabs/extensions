# youki extension

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

Apply the following manifest to run an nginx pod using the Youki Runtime:

```yaml
apiVersion: node.k8s.io/v1
kind: RuntimeClass
metadata:
  name: youki
handler: youki
```

## Testing

Apply the following manifest to run nginx pod via youki:

```yaml

---
apiVersion: v1
kind: Pod
metadata:
  name: nginx-youki
spec:
  runtimeClassName: youki
  containers:
  - name: nginx
    image: nginx
```

The pod should be up and running:

```bash
$ kubectl get pods
NAME           READY   STATUS    RESTARTS   AGE
nginx-youki   1/1     Running   0          40s
```
