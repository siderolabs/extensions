# NVIDIA Container toolkit extension


## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).


## Usage

The following NVIDIA modules needs to be loaded, so add this to the talos config:

```yaml
machine:
  kernel:
    modules:
      - name: nvidia
      - name: nvidia_uvm
      - name: nvidia_drm
      - name: nvidia_modeset
```

`nvidia-container-cli` loads BPF programs and requires relaxed KSPP setting for [bpf_jit_harden](https://sysctl-explorer.net/net/core/bpf_jit_harden/), so Talos default setting
should be overridden:

```yaml
machine:
  sysctls:
    net.core.bpf_jit_harden: 1
```

> Warning! This disables [KSPP best practices](https://kernsec.org/wiki/index.php/Kernel_Self_Protection_Project/Recommended_Settings#sysctls) setting.

## Testing

Apply the following manifest to create a runtime class that uses the extension:

```yaml
---
apiVersion: node.k8s.io/v1
kind: RuntimeClass
metadata:
  name: nvidia
handler: nvidia
```

Install the NVIDIA device plugin:

```bash
helm repo add nvdp https://nvidia.github.io/k8s-device-plugin
helm repo update
helm install nvidia-device-plugin nvdp/nvidia-device-plugin --version=0.14.1 --set=runtimeClassName=nvidia
```

Apply the following manifest to run CUDA pod via nvidia runtime:

```yaml
---
apiVersion: v1
kind: Pod
metadata:
  name: gpu-operator-test
spec:
  restartPolicy: OnFailure
  runtimeClassName: nvidia
  containers:
  - name: cuda-vector-add
    image: "nvidia/samples:vectoradd-cuda11.6.0"
    resources:
      limits:
         nvidia.com/gpu: 1
```


The status can be viewed by running:

```bash
❯ kubectl get pods
NAME                READY   STATUS      RESTARTS   AGE
gpu-operator-test   0/1     Completed   0          13s
```

```bash
❯ kubectl logs gpu-operator-test
[Vector addition of 50000 elements]
Copy input data from the host memory to the CUDA device
CUDA kernel launch with 196 blocks of 256 threads
Copy output data from the CUDA device to the host memory
Test PASSED
Done
```
