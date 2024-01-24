# WasmEdge extension

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

In order to create the Wasm workload, a runtimeclass needs to be created.

```yaml
apiVersion: node.k8s.io/v1
kind: RuntimeClass
metadata:
  name: wasmedge
handler: wasmedge
```

## Testing

Apply the following manifest to run sample pod using wasmedge:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: wasmedge-test
spec:
  restartPolicy: Never
  runtimeClassName: wasmedge
  containers:
  - name: wasmedge-test
    image: wasmedge/example-wasi:latest
```

The pod should run without any errors:

```bash
$ kubectl get pods
NAME            READY   STATUS      RESTARTS   AGE
wasmedge-test   0/1     Completed   0          28s
```
