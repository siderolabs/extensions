# Kubelet ACR Credential Provider extension

This extension provides the [acr-credential-provider](https://github.com/kubernetes-sigs/cloud-provider-azure/tree/master/cmd/acr-credential-provider) binary,
which can be executed by Kubelet to provide a short-lived token for pulling container images
from Microsoft Azure Container Registry (ACR).

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

The `acr-credential-provider` requires an Azure cloud provider configuration file
which defines how to authenticate against Azure (e.g. managed identity or service principal).
The file must be placed in a location accessible to Kubelet, e.g. `/var/lib/kubelet/azure.json`.

> **Note:** `machine.files` is deprecated, but its replacement (`EtcFileConfig`)
> can't write files to any location visible to the Kubelet, so `machine.files`
> remains the only way to provide the configuration file.

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

You also need to configure `acr-credential-provider` as a Kubelet image
credential provider. For this you'll need patch the machine config with the following:

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
        - /var/lib/kubelet/azure.json
```

For Talos versions older than v1.14, use the `machine.kubelet.credentialProviderConfig`
field instead:

```yaml
machine:
  kubelet:
    credentialProviderConfig:
      apiVersion: kubelet.config.k8s.io/v1
      kind: CredentialProviderConfig
      providers:
        - name: acr-credential-provider
          # ... same as above
```

## More Information

- <https://cloud-provider-azure.sigs.k8s.io/>
- <https://kubernetes.io/docs/tasks/administer-cluster/kubelet-credential-provider/>
- <https://kubernetes.io/docs/reference/config-api/kubelet-credentialprovider.v1/>
