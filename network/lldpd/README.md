# LLDPD

Adds https://lldpd.github.io/ as system extensions.
This means a lldpd server is started that sends/receives LLDP messages.

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

Configure the extension via `ExtensionServiceConfig` document. You can add any lldpd related configuration and these will be executed at the LLDPD server startup.

```yaml
---
apiVersion: v1alpha1
kind: ExtensionServiceConfig
name: lldpd
configFiles:
  - content: |
      configure lldp portidsubtype ifname
      unconfigure lldp management-addresses-advertisements
      unconfigure lldp capabilities-advertisements
      configure system description "Talos Node"
    mountPath: /usr/local/etc/lldpd/lldpd.conf
```

Then apply the patch to your node's MachineConfigs
```bash
talosctl patch mc -p @lldpd-config.yaml
```

You will then be able to verify that it is in place with the following command
```bash
talosctl get extensionserviceconfigs

NODE          NAMESPACE   TYPE                     ID           VERSION
mynode   runtime     ExtensionServiceConfig   lldpd           1
```
