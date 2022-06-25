# development

This document is intended as a guide to updating the `nvidia-container-toolkit` dependencies.

## Components

### [nvidia-container-cli](./nvidia-container-cli/)

`nvidia-container-cli` is called by the `nvidia-container-runtime` to setup the required NVIDIA library mounts and NVIDIA device files for a workload container

### [nvidia-container-runtime](./nvidia-container-runtime/)

`nvidia-container-runtime` is the runtime used by `containerd` to run workload containers. It's mostly a wrapper around `runc`

It also ships a tool called `nvidia-container-runtime-hook` which is used to setup OCI hooks, it's a symlink to `nvidia-container-toolkit`, which eventually calls `nvidia-container-cli`

### [nvidia-device-create](./nvidia-device-create/)

This is used to create the required NVIDIA device files under `/dev`. This required udev rules.

### [glibc](./glibc/)

`nvidia-container-cli` is fully dependent on `glibc` to be able to access the NVIDIA shared objects.

## Updating the nvidia driver version

- Update the driver version in `pkgs` repo [here](https://github.com/siderolabs/pkgs/blob/master/nonfree/kmod-nvidia/pkg.yaml)
- Update the driver version [here](../vars.yaml)
- Update the version checksums [here](./nvidia-pkgs/pkg.yaml)

## Updating the nvidia-container-toolkit version

- Update the `libnvidia-container` version [here](./nvidia-container-cli/pkg.yaml)
- Update the `container-toolkit` version [here](./nvidia-container-runtime/pkg.yaml)

Make sure to also update the `nvidia-device-create` [here](./nvidia-device-create/pkg.yaml)

### Patches

- [nvidia-container-cli](./nvidia-container-cli/patches/libnvidia-container/)
    - `common.h.patch` - use custom glibc interpreter path
    - `Makefile.patch` - build statically linked with `libcap` and `libseccomp`
    - `nvc_ldcache.c.patch` - use the standard `ld.so.cache` path inside the container
- [container-runtime](./nvidia-container-runtime/patches/nvidia-container-runtime/)
    - `main.go.patch` - use custom path for the nvidia-container-runtime config
- [container-runtime](./nvidia-container-runtime/patches/nvidia-container-toolkit/)
    - `hook_config.go.patch` - use custom path for the nvidia-container-runtime config
- [nvidia-device-create](./nvidia-device-create/patches/nvidia-graphics-drivers-build/)
    - Makefile.patch - build statically linked with `libpciaccess`
