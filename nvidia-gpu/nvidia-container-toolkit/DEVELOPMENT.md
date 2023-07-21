# development

This document is intended as a guide to updating the `nvidia-container-toolkit` dependencies.

## Components

### [nvidia-container-cli](./nvidia-container-cli/)

`nvidia-container-cli` is called by the `nvidia-container-runtime` to setup the required NVIDIA library mounts and NVIDIA device files for a workload container

### [nvidia-container-runtime](./nvidia-container-runtime/)

`nvidia-container-runtime` is the runtime used by `containerd` to run workload containers. It's mostly a wrapper around `runc`

It also ships a tool called `nvidia-container-runtime-hook` which is used to setup OCI hooks.

### [glibc](./glibc/)

`nvidia-container-cli` is fully dependent on `glibc` to be able to access the NVIDIA shared objects.

## Updating the nvidia driver version

- Update the driver version in `pkgs` repo [here](https://github.com/siderolabs/pkgs/blob/master/nonfree/kmod-nvidia/pkg.yaml)
- Update the driver version [here](../vars.yaml)
- Update the version checksums [here](./nvidia-pkgs/pkg.yaml)

## Updating the nvidia-container-toolkit version

- Update the `libnvidia-container` version checksums and `REVISION` [here](./nvidia-container-cli/pkg.yaml)
- Update the `container-toolkit` version checksums and `GIT_COMMIT` [here](./nvidia-container-runtime/pkg.yaml)
