# Spin extension

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

In order to create the Wasm workload, a runtimeclass needs to be created.

```yaml
apiVersion: node.k8s.io/v1
kind: RuntimeClass
metadata:
  name: wasmtime-spin-v2
handler: spin
```

## Testing

Apply the following manifest to run sample pod using spin:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: spin-test
spec:
  containers:
  - command:
    - /
    image: ghcr.io/spinkube/containerd-shim-spin/examples/spin-rust-hello
    name: spin-test
  runtimeClassName: wasmtime-spin-v2
```

The pod should run without any errors:

```bash
$ kubectl get pods
NAME            READY   STATUS      RESTARTS   AGE
spin-test   1/1     Running   0          28s
```
