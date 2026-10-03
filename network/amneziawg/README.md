# AmneziaWG

Adds [AmneziaWG](https://docs.amnezia.org/documentation/amnezia-wg/) VPN interfaces as system extensions.
AmneziaWG is a WireGuard-based VPN protocol with DPI bypass capabilities through traffic obfuscation (header randomization, packet padding, jitter).

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

Configure the extension via `ExtensionServiceConfig` document.

```yaml
---
apiVersion: v1alpha1
kind: ExtensionServiceConfig
name: amneziawg
configFiles:
  - content: |
      [Interface]
      PrivateKey = <your-private-key>
      ListenPort = 51820
      Address = 10.0.0.1/24
      Jc = 5
      Jmin = 50
      Jmax = 1000
      S1 = 52
      S2 = 52
      H1 = 12345678
      H2 = 87654321
      H3 = 11223344
      H4 = 44332211

      [Peer]
      PublicKey = <peer-public-key>
      Endpoint = vpn.example.com:51820
      AllowedIPs = 0.0.0.0/0
      PersistentKeepalive = 25
    mountPath: /etc/amneziawg/awg0.conf
```

Then apply the patch to your node's MachineConfigs
```bash
talosctl patch mc -p @amneziawg-config.yaml
```

You will then be able to verify that it is in place with the following command
```bash
talosctl get extensionserviceconfigs

NODE          NAMESPACE   TYPE                     ID           VERSION
mynode   runtime     ExtensionServiceConfig   amneziawg    1
```

## Configuration

The config file follows the standard WireGuard format with additional AmneziaWG-specific parameters for DPI obfuscation:

| Parameter | Description |
|-----------|-------------|
| `Jc` | Junk packet count |
| `Jmin` | Minimum junk packet size |
| `Jmax` | Maximum junk packet size |
| `S1` | Init packet padding size |
| `S2` | Response packet padding size |
| `H1` | Init packet header value |
| `H2` | Response packet header value |
| `H3` | Cookie packet header value |
| `H4` | Transport packet header value |

All standard WireGuard fields (`PrivateKey`, `ListenPort`, `Address`, `PublicKey`, `Endpoint`, `AllowedIPs`, `PersistentKeepalive`, `MTU`, `PostUp`, `PostDown`) are also supported.

### Environment overrides

Environment variables can be set via `ExtensionServiceConfig`:

```yaml
---
apiVersion: v1alpha1
kind: ExtensionServiceConfig
name: amneziawg
environment:
  - AWG_INTERFACE=awg0
  - AWG_CONFIG=/etc/amneziawg/awg0.conf
  - AWG_LOG_LEVEL=verbose
configFiles:
  - content: |
      ...
    mountPath: /etc/amneziawg/awg0.conf
```

| Variable | Default | Description |
|----------|---------|-------------|
| `AWG_INTERFACE` | `awg0` | Network interface name |
| `AWG_CONFIG` | `/etc/amneziawg/awg0.conf` | Path to config file |
| `AWG_LOG_LEVEL` | `info` | Log level (`info` or `verbose`) |

For more information see https://docs.amnezia.org/documentation/amnezia-wg/
