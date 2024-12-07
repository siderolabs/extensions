# Cloudflare Tunnel

Cloudflare Tunnel securely connects resources to Cloudflare without a public IP. A lightweight daemon (cloudflared) creates outbound-only connections to Cloudflare, allowing safe access to services like HTTP, SSH, remote desktops, and other protocols.

More info: https://github.com/cloudflare/cloudflared/

## Installation

Cloudflared system extension can be installed by customising boot assets or after installation with the `installer`

You can use the following schematic file:
```yaml
# cloudflared-ext.yaml
customization:
  systemExtensions:
    officialExtensions:
      - siderolabs/cloudflared
```

Check documentation for install:
* https://www.talos.dev/latest/talos-guides/configuration/system-extensions/
* https://www.talos.dev/latest/talos-guides/install/boot-assets/

## Usage

Configure the extension via `ExtensionServiceConfig` document.

```yaml
# cloudflared-config.yaml
---
apiVersion: v1alpha1
kind: ExtensionServiceConfig
name: cloudflared
environment:
  - TUNNEL_TOKEN=<your_token>
  - TUNNEL_METRICS=localhost:2000
  - TUNNEL_EDGE_IP_VERSION=auto   # if your node is only configured for IPv6
```

Then apply the patch to your node's MachineConfigs
```bash
talosctl patch mc -p @cloudflared-config.yaml
```

You will then be able to verify that it is in place with the following command
```bash
talosctl get extensionserviceconfigs

NODE     NAMESPACE   TYPE                     ID            VERSION
mynode   runtime     ExtensionServiceConfig   cloudflared   1
```

## Configuration

See all run parameters here (use environment variables): https://developers.cloudflare.com/cloudflare-one/connections/connect-networks/configure-tunnels/tunnel-run-parameters/
