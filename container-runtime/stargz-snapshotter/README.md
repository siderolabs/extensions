# Stargz Snapshotter extension

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Pulling from Private Registries

To pull from private registries an additional step is required. You must configure the Kubelet to use the Stargz snapshotter as an image service proxy. This is explained in more detail in the [Stargz Snapshotter docs](https://github.com/containerd/stargz-snapshotter/blob/main/docs/overview.md#cri-based-authentication). An example config patch:

```yaml
machine:
  kubelet:
    extraConfig:
      imageServiceEndpoint: unix:///var/run/containerd-stargz-grpc/containerd-stargz-grpc.sock
```

## Enabling on a Node with Existing Images

Talos sets `discard_unpacked_layers = true` for the CRI plugin, so layer blobs of images pulled before the extension was installed were deleted after unpacking into the `overlayfs` snapshotter.
After switching to the Stargz snapshotter, such images are still reported as present, but containerd cannot unpack them into the new snapshotter and container creation fails with an error like:

```text
failed to extract layer (...) to stargz as "extract-...": failed to get reader from content store: content digest sha256:...: not found
```

To fix this, remove the affected images so that they are pulled again, or wipe the `EPHEMERAL` partition after installing the extension.
