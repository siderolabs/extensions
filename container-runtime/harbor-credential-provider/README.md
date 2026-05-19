# Kubelet Harbor Credential Provider extension

This extension provides the `harbor-credential-provider` binary (built from
[container-registry/harbor-workload-identity-federation](https://github.com/container-registry/harbor-workload-identity-federation/tree/main/cmd/credential-provider-harbor)),
which can be executed by Kubelet to obtain short-lived credentials for pulling container images
from a [Harbor](https://goharbor.io/) / [8gears Container Registry](https://8gears.container-registry.com/)
instance configured for [Workload Identity Federation](https://github.com/container-registry/harbor-workload-identity-federation).

Instead of using static robot account secrets, the provider exchanges a Kubernetes ServiceAccount
token (issued by the kubelet) for Basic Auth credentials (`jwt:<SA-token>`) that Harbor validates
against a configured Federated Identity Provider (FedIDP).

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

You also need to configure `harbor-credential-provider` as a Kubelet image
credential provider. For this you'll need to patch the machine config with the following
(replacing the registry hostname with the value configured in Harbor as the FedIDP audience):

```yaml
machine:
  kubelet:
    credentialProviderConfig:
      kind: CredentialProviderConfig
      apiVersion: kubelet.config.k8s.io/v1
      providers:
        - name: harbor-credential-provider
          apiVersion: credentialprovider.kubelet.k8s.io/v1
          tokenAttributes:
            requireServiceAccount: true
            serviceAccountTokenAudience: "harbor.example.com"
            cacheType: Token
          matchImages:
            - "harbor.example.com"
          defaultCacheDuration: "1h"
          args:
            - "--username=jwt"
```

`defaultCacheDuration` is required by the kubelet config API. The binary itself
returns a zero cache duration because it passes the service account token through
as the registry password.

Kubernetes 1.34+ is required for ServiceAccount-token credential provider integration,
which on Talos means v1.11.0 or later (see the [Talos v1.11 support matrix](https://docs.siderolabs.com/talos/v1.11/getting-started/support-matrix)).

## More Information

- <https://github.com/container-registry/harbor-workload-identity-federation>
- <https://kubernetes.io/docs/tasks/administer-cluster/kubelet-credential-provider/>
- <https://kubernetes.io/docs/reference/config-api/kubelet-credentialprovider.v1/>
- <https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/4412-projected-service-account-tokens-for-kubelet-image-credential-providers>
