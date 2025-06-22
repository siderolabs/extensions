# Newt client

Newt is a fully user space WireGuard tunnel client and TCP/UDP proxy, designed to securely expose private resources controlled by Pangolin.
By using Newt, you don't need to manage complex WireGuard tunnels and NATing.
More info: https://github.com/fosrl/newt

## Installation

Newt system extension can be installed by customising boot assets or after installation with the `installer`

You can use the following schematic file:
```yaml
# newt-ext.yaml
customization:
  systemExtensions:
    officialExtensions:
      - siderolabs/newt
```

Check documentation for install:
* https://www.talos.dev/latest/talos-guides/configuration/system-extensions/
* https://www.talos.dev/latest/talos-guides/install/boot-assets/

## Usage

Configure the extension via `ExtensionServiceConfig` document.

```yaml
# newt-config.yaml
---
apiVersion: v1alpha1
kind: ExtensionServiceConfig
name: newt
environment:
  - PANGOLIN_ENDPOINT=https://example.com
  - NEWT_ID=2ix2t8xk22ubpfy
  - NEWT_SECRET=nnisrfsdfc7prqsp9ewo1dvtvci50j5uiqotez00dgap0ii2
```

Then apply the patch to your node's MachineConfigs
```bash
talosctl patch mc -p @newt-config.yaml
```

You will then be able to verify that it is in place with the following command
```bash
talosctl get extensionserviceconfigs

NODE     NAMESPACE   TYPE                     ID            VERSION
mynode   runtime     ExtensionServiceConfig   newt   1
```

## Configuration

See all run parameters here (use environment variables): https://docs.fossorial.io/Newt/overview#cli-args
