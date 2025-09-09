# Beszel Agent
Runs on each system you want to monitor and communicates system metrics to the hub.
More details: https://github.com/henrygd/beszel

## Installation

Beszel-agent system extension can be installed by customising boot assets or after installation with the `installer`

You can use the following schematic file:
```yaml
# beszel-agent-ext.yaml
customization:
  systemExtensions:
    officialExtensions:
      - siderolabs/beszel-agent
```

Check documentation for install:
* https://www.talos.dev/latest/talos-guides/configuration/system-extensions/
* https://www.talos.dev/latest/talos-guides/install/boot-assets/

## Usage

Configure the extension via `ExtensionServiceConfig` document.

```yaml
# beszel-agent-config.yaml
---
apiVersion: v1alpha1
kind: ExtensionServiceConfig
name: beszel-agent
environment:
  - TOKEN=<token from hub optional>
  - KEY=<pub key from hub>
  - HUB_URL=<hub address optional>

```

Then apply the patch to your node's MachineConfigs
```bash
talosctl patch mc -p @beszel-agent-config.yaml
```

You will then be able to verify that it is in place with the following command
```bash
talosctl get extensionserviceconfigs

NODE     NAMESPACE   TYPE                     ID            VERSION
mynode   runtime     ExtensionServiceConfig   beszel-agent   3
```

## Configuration

See all run parameters here (use environment variables):
https://beszel.dev/guide/environment-variables#agent