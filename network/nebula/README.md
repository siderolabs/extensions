# Nebula

https://github.com/slackhq/nebula

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

Configure the extension via `ExtensionServiceConfig` document. You can add any nebula related configuration and these will be executed at Nebula startup.

```yaml
---
apiVersion: v1alpha1
kind: ExtensionServiceConfig
name: nebula
configFiles:
  - content: |
        pki:
        ca: /usr/local/etc/nebula/ca.crt
        cert: /usr/local/etc/nebula/node.crt
        key: /usr/local//etc/nebula/node.key
    mountPath: /usr/local/etc/nebula/config.yml
  - content: |
        -----BEGIN NEBULA CERTIFICATE-----
        -----END NEBULA CERTIFICATE-----
    mountPath: /usr/local/etc/nebula/ca.crt
   - content: |
        -----BEGIN NEBULA CERTIFICATE-----
        -----END NEBULA CERTIFICATE-----
     mountPath: /usr/local/etc/nebula/node.crt
   - content: |
     mountPath: /usr/local/etc/nebula/node.key

```

Then apply the patch to your node's MachineConfigs
```bash
talosctl patch mc -p @nebula-config.yaml
```

You will then be able to verify that it is in place with the following command
```bash
talosctl get extensionserviceconfigs

NODE          NAMESPACE   TYPE                     ID           VERSION
mynode   runtime     ExtensionServiceConfig   nebula            1
```
