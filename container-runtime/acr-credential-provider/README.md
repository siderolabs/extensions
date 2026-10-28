# Kubelet ACR Credential Provider extension

This extension provides the [acr-credential-provider](https://github.com/kubernetes-sigs/cloud-provider-azure/tree/master/cmd/acr-credential-provider) binary,
which can be executed by Kubelet to provide a short-lived token for pulling container images
from Microsoft Azure Container Registry (ACR).

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

The `acr-credential-provider` requires an Azure cloud provider configuration file
which defines how to authenticate against Azure (e.g. managed identity or service principal).

Configure `acr-credential-provider` as a Kubelet image credential provider and
provide the configuration file by patching the machine config with the following
(Talos v1.15+):

```yaml
apiVersion: v1alpha1
kind: KubeCredentialProviderConfig
configuration:
  apiVersion: kubelet.config.k8s.io/v1
  kind: CredentialProviderConfig
  providers:
    - name: acr-credential-provider
      matchImages:
        - "*.azurecr.io"
        - "*.azurecr.cn"
        - "*.azurecr.de"
        - "*.azurecr.us"
      defaultCacheDuration: "10m"
      apiVersion: credentialprovider.kubelet.k8s.io/v1
      args:
        - /etc/kubernetes/credentialproviders/azure.json
configFiles:
  - name: azure.json
    content: |
      {
        "cloud": "AzurePublicCloud",
        "tenantId": "<tenant-id>",
        "useManagedIdentityExtension": true,
        "userAssignedIdentityID": "<managed-identity-client-id>"
      }
```

Each entry of `configFiles` is written to the `/etc/kubernetes/credentialproviders`
directory with the permissions `0600`.

### Older Talos versions

Talos versions older than v1.15 don't support `configFiles`. Create the
configuration file via `machine.files` instead and reference it in the `args`:

```yaml
machine:
  files:
    - path: /var/lib/kubelet/azure.json
      permissions: 0o600
      op: create
      content: |
        {
          "cloud": "AzurePublicCloud",
          "tenantId": "<tenant-id>",
          "useManagedIdentityExtension": true,
          "userAssignedIdentityID": "<managed-identity-client-id>"
        }
```

On Talos v1.14, use the `KubeCredentialProviderConfig` document above without
`configFiles` and with `args` set to `/var/lib/kubelet/azure.json`. On Talos
versions older than v1.14, use the `machine.kubelet.credentialProviderConfig`
field instead:

```yaml
machine:
  kubelet:
    credentialProviderConfig:
      apiVersion: kubelet.config.k8s.io/v1
      kind: CredentialProviderConfig
      providers:
        - name: acr-credential-provider
          matchImages:
            - "*.azurecr.io"
            - "*.azurecr.cn"
            - "*.azurecr.de"
            - "*.azurecr.us"
          defaultCacheDuration: "10m"
          apiVersion: credentialprovider.kubelet.k8s.io/v1
          args:
            - /var/lib/kubelet/azure.json
```

## More Information

- <https://cloud-provider-azure.sigs.k8s.io/>
- <https://kubernetes.io/docs/tasks/administer-cluster/kubelet-credential-provider/>
- <https://kubernetes.io/docs/reference/config-api/kubelet-credentialprovider.v1/>
