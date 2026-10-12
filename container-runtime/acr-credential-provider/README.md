# Kubelet ACR Credential Provider extension

This extension provides the [acr-credential-provider](https://github.com/kubernetes-sigs/cloud-provider-azure/tree/master/cmd/acr-credential-provider) binary,
which can be executed by Kubelet to provide a short-lived token for pulling container images
from Microsoft Azure Container Registry (ACR).

## Installation

This extension requires Talos v1.15 or later, as it relies on the `configFiles`
field of the `KubeCredentialProviderConfig` document.

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

The `acr-credential-provider` requires an Azure cloud provider configuration file
which defines how to authenticate against Azure (e.g. managed identity or service principal).

Configure `acr-credential-provider` as a Kubelet image credential provider and
provide the configuration file by patching the machine config with the following:

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

## More Information

- <https://cloud-provider-azure.sigs.k8s.io/>
- <https://kubernetes.io/docs/tasks/administer-cluster/kubelet-credential-provider/>
- <https://kubernetes.io/docs/reference/config-api/kubelet-credentialprovider.v1/>
