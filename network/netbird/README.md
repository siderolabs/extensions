# Netbird

Adds https://netbird.io network interfaces as system extensions.
This means you can access your talos nodes from machines you have configured
with netbird

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

Configure the extension via `ExtensionServiceConfig` document.

```yaml
---
apiVersion: v1alpha1
kind: ExtensionServiceConfig
name: netbird
environment:
  - NB_SETUP_KEYS=<peer setup key>
```

or if you are selfhosting it with something like :

```yaml
---
apiVersion: v1alpha1
kind: ExtensionServiceConfig
name: netbird
environment:
  - NB_SETUP_KEYS=<peer setup key>
  - NB_MANAGEMENT_URL=https://netbird.selfhosted:443
  - NB_ADMIN_URL=https://netbird.selfhosted:443
```

Then apply the patch to your node's MachineConfigs
```bash
talosctl patch mc -p @netbird-config.yaml
```

You will then be able to verify that it is in place with the following command
```bash
talosctl get extensionserviceconfigs

NODE          NAMESPACE   TYPE                     ID           VERSION
mynode   runtime     ExtensionServiceConfig   netbird   1
```
## Configuration

For more configuration options please reffer to docs https://docs.netbird.io/