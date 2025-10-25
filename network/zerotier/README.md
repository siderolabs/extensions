# ZeroTier

Adds https://zerotier.com network interfaces as system extensions.
This means you can access your Talos nodes from machines you have configured
with ZeroTier, creating a secure overlay network.

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

Configure the extension via `ExtensionServiceConfig` document.

```yaml
---
apiVersion: v1alpha1
kind: ExtensionServiceConfig
name: zerotier
environment:
  - ZEROTIER_NETWORK=<your network id>
```

Then apply the patch to your node's MachineConfigs

```bash
talosctl patch mc -p @zerotier-config.yaml
```

You can then verify that it is in place with the following command

```bash
talosctl get extensionserviceconfigs

NODE          NAMESPACE   TYPE                     ID         VERSION
mynode        runtime     ExtensionServiceConfig   zerotier   1
```

## Configuration

The extension can be configured through environment variables:

- `ZEROTIER_NETWORK`: The network ID to join (required, you can also specify multiple network by separting them with ",")
- `ZEROTIER_IDENTITY_SECRET`: Optional pre-existing identity to use (format: "address:0:public:private")
- `ZEROTIER_PLANET`: Optional pre-existing planet file encoded in base64

### Using an existing identity

If you want to maintain the same ZeroTier identity across rebuilds or different nodes, you can specify an existing identity:

```yaml
---
apiVersion: v1alpha1
kind: ExtensionServiceConfig
name: zerotier
environment:
  - ZEROTIER_NETWORK=<your network id>
  - ZEROTIER_IDENTITY_SECRET=<identity string>
```

### Join multiple network

If you want to join multiple zerotier network, you can use the following format:

```yaml
---
apiVersion: v1alpha1
kind: ExtensionServiceConfig
name: zerotier
environment:
  - ZEROTIER_NETWORK=<your network id 1>,<your network id 2>,<your network id 3>
```

If no identity is provided, a new one will be generated automatically. (You may need to authorize this node in your Zerotier network according to your network policies before it will recieve an IP address).

### Using an custom planet file

If you want to specify custom planet file from a hosted planet, you can specify an custom planet:

```yaml
---
apiVersion: v1alpha1
kind: ExtensionServiceConfig
name: zerotier
environment:
  - ZEROTIER_NETWORK=<your network id>
  - ZEROTIER_IDENTITY_SECRET=<identity string>
  - ZEROTIER_PLANET=<base64 encoded planet file>
```

If no planet is provided, the public planet file from ZeroTier will be used.
