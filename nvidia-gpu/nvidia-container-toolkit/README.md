# NVIDIA driver userspace extension

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

Install this extension together with the matching NVIDIA kernel module extension.

## Usage

Load the NVIDIA modules in the Talos machine configuration:

```yaml
---
apiVersion: v1alpha1
kind: KernelModuleConfig
name: nvidia
---
apiVersion: v1alpha1
kind: KernelModuleConfig
name: nvidia_uvm
---
apiVersion: v1alpha1
kind: KernelModuleConfig
name: nvidia_drm
---
apiVersion: v1alpha1
kind: KernelModuleConfig
name: nvidia_modeset
```

Install NVIDIA GPU Operator with its driver disabled. The operator installs the container toolkit under writable `/var` and registers it with containerd through NRI. No NVIDIA `RuntimeClass` is required.

```bash
helm install gpu-operator gpu-operator \
  --repo https://helm.ngc.nvidia.com/nvidia \
  --namespace gpu-operator \
  --create-namespace \
  --set driver.enabled=false \
  --set toolkit.enabled=true \
  --set toolkit.installDir=/var/lib/nvidia \
  --set cdi.enabled=true \
  --set cdi.nriPluginEnabled=true \
  --set hostPaths.driverInstallDir=/usr/local
```

## Testing

Run a normal CUDA workload requesting a GPU. Do not set `runtimeClassName`.

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: cuda-test
spec:
  restartPolicy: OnFailure
  containers:
    - name: cuda-test
      image: nvcr.io/nvidia/k8s/cuda-sample:vectoradd-cuda12.5.0
      resources:
        limits:
          nvidia.com/gpu: 1
```

A successful run ends with `Test PASSED` in the pod logs.
