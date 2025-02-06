# containerd-image-verifier-sigstore extension

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

```yaml
machine:
  files:
    - content: |
        apiVersion: policy.sigstore.dev/v1alpha1
        kind: ClusterImagePolicy
        metadata:
          name: system
        spec:
          images:
          - glob: "**"
          authorities:
          - keyless:
              url: https://fulcio.sigstore.dev
              identities:
              - issuer: https://accounts.google.com
                subject: k8s-infra-gcr-promoter@k8s-artifacts-prod.iam.gserviceaccount.com
            ctlog:
              url: https://rekor.sigstore.dev
      path: /var/local/etc/containers/sigstore/kubernetes.yaml
      op: create
    - content: |
        apiVersion: policy.sigstore.dev/v1alpha1
        kind: ClusterImagePolicy
        metadata:
          name: system
        spec:
          images:
          - glob: "**"
          authorities:
          - keyless:
              identities:
              - issuer: https://accounts.google.com
                subjectRegExp: "@siderolabs\.com$"
      path: /var/local/etc/containers/sigstore/siderolabs.yaml
      op: create
```

**Important note: add all other identities and keys within the ClusterImagePolicy above for target container images**
