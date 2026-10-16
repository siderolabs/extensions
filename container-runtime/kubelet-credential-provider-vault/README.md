# Kubelet Vault Credential Provider extension

This extension provides the `kubelet-credential-provider-vault` binary (built from
[simonostendorf/kubelet-credential-provider-vault](https://github.com/simonostendorf/kubelet-credential-provider-vault)),
which can be executed by Kubelet to read image registry credentials from
[HashiCorp Vault](https://www.hashicorp.com/products/vault) or
[OpenBao](https://openbao.org/).

Instead of an image pull Secret in every namespace, the provider authenticates against Vault with
the Kubernetes ServiceAccount token the kubelet issues for the pod, using Vault's
[Kubernetes auth method](https://developer.hashicorp.com/vault/docs/auth/kubernetes), and reads a
kv v2 secret holding `username` and `password`.

This works with any registry whose credentials are kept in Vault. The registry itself needs no
federation support, unlike `harbor-credential-provider`, which exchanges the token with Harbor
directly.

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

You also need to configure `kubelet-credential-provider-vault` as a Kubelet image credential
provider. For this you'll need to patch the machine config with the following:

```yaml
apiVersion: v1alpha1
kind: KubeCredentialProviderConfig
configuration:
  apiVersion: kubelet.config.k8s.io/v1
  kind: CredentialProviderConfig
  providers:
    - name: kubelet-credential-provider-vault
      apiVersion: credentialprovider.kubelet.k8s.io/v1
      matchImages:
        - registry.example.com
      defaultCacheDuration: 12h
      args:
        - --vault-addr=https://vault.example.com
        - --vault-auth-method=kubernetes
        - --vault-auth-mount=kubernetes
        - --vault-auth-role=image-pull
        - --vault-secret-mount=registry
        - --vault-secret-path=credentials
      tokenAttributes:
        requireServiceAccount: true
        serviceAccountTokenAudience: vault-image-pull
        cacheType: ServiceAccount
```

The Vault role has to accept the audience configured as `serviceAccountTokenAudience`, and the
policy attached to it needs `read` on the kv path given by `--vault-secret-mount` and
`--vault-secret-path`.

The kubelet issues the token for whichever ServiceAccount the pod runs as, so the role cannot bind a
single ServiceAccount name. Binding the namespaces that may pull, together with the audience, is
what keeps it narrow: an ordinary projected token carries a different audience.

`cacheType: ServiceAccount` caches the credential per ServiceAccount, which fits a credential that
does not depend on the individual pod. Use `Token` where the credential's lifetime is tied to the
token.

Kubernetes 1.34+ is required for ServiceAccount-token credential provider integration,
which on Talos means v1.11.0 or later (see the [Talos v1.11 support matrix](https://docs.siderolabs.com/talos/v1.11/getting-started/support-matrix)).

## More Information

- <https://github.com/simonostendorf/kubelet-credential-provider-vault>
- <https://kubernetes.io/docs/tasks/administer-cluster/kubelet-credential-provider/>
- <https://kubernetes.io/docs/reference/config-api/kubelet-credentialprovider.v1/>
- <https://github.com/kubernetes/enhancements/tree/master/keps/sig-auth/4412-projected-service-account-tokens-for-kubelet-image-credential-providers>
