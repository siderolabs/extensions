# Kubelet ECR Credential Provider extension

This extension provides the [ecr-credential-provider](https://github.com/kubernetes/cloud-provider-aws/tree/master/cmd/ecr-credential-provider) binary,
which can be executed by Kubelet to provide a short-lived token for pulling container images
from Amazon Web Services' Elastic Container Registry (ECR).

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

You also need to configure `ecr-credential-provider` as a Kubelet image
credential provider. For this you'll need patch the machine config with the following:

```yaml
machine:
  kubelet:
    credentialProviderConfig:
      apiVersion: kubelet.config.k8s.io/v1
      kind: CredentialProviderConfig
      providers:
        - name: ecr-credential-provider
          matchImages:
            - "*.dkr.ecr.*.amazonaws.com"
            - "*.dkr.ecr.*.amazonaws.com.cn"
            - "*.dkr.ecr-fips.*.amazonaws.com"
            - "*.dkr.ecr.us-iso-east-1.c2s.ic.gov"
            - "*.dkr.ecr.us-isob-east-1.sc2s.sgov.gov"
          defaultCacheDuration: "12h"
          apiVersion: credentialprovider.kubelet.k8s.io/v1
```

## More Information

- <https://cloud-provider-aws.sigs.k8s.io/credential_provider/>
- <https://kubernetes.io/docs/tasks/administer-cluster/kubelet-credential-provider/>
- <https://kubernetes.io/docs/reference/config-api/kubelet-credentialprovider.v1/>
