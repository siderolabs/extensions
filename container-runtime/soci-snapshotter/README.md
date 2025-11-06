# AWS SOCI Snapshotter extension

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Pulling from Privte Registries

To pull from private registries an additional step is required. You must configure the Kubelet to use the SOCI snapshotter as an image service proxy. This is explained in more detail in the [SOCI docs](https://github.com/awslabs/soci-snapshotter/blob/main/docs/registry-authentication.md#kubernetes-cri-credentials). An example config patch:

```yaml
machine:
  kubelet:
    extraConfig:
      imageServiceEndpoint: unix:///var/run/soci-snapshotter/soci-snapshotter-grpc.sock
```
