# OpenConnect

OpenConnect is a multi-protocol SSL VPN client supporting:

- Cisco AnyConnect
- Juniper Network Connect
- Palo Alto GlobalProtect
- Pulse Secure
- F5 BIG-IP
- Fortinet SSL VPN
- Array Networks SSL VPN

This extension allows you to connect your Talos nodes to enterprise VPN networks.

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

Configure the extension via `ExtensionServiceConfig` document.

### Basic Usage (Username/Password)

```yaml
---
apiVersion: v1alpha1
kind: ExtensionServiceConfig
name: openconnect
environment:
  - OPENCONNECT_SERVER=vpn.example.com
  - OPENCONNECT_USER=username
  - OPENCONNECT_PASSWORD=password
  - OPENCONNECT_SERVERCERT=pin-sha256:ABC123...
```

### Certificate Authentication

```yaml
---
apiVersion: v1alpha1
kind: ExtensionServiceConfig
name: openconnect
environment:
  - OPENCONNECT_SERVER=vpn.example.com
  - OPENCONNECT_CERTIFICATE=/var/lib/openconnect/client.crt
  - OPENCONNECT_PRIVATE_KEY=/var/lib/openconnect/client.key
  - OPENCONNECT_SERVERCERT=pin-sha256:ABC123...
```

### GlobalProtect Example

```yaml
---
apiVersion: v1alpha1
kind: ExtensionServiceConfig
name: openconnect
environment:
  - OPENCONNECT_SERVER=vpn.example.com
  - OPENCONNECT_PROTOCOL=gp
  - OPENCONNECT_USER=username
  - OPENCONNECT_PASSWORD=password
```

### Pulse Secure Example

```yaml
---
apiVersion: v1alpha1
kind: ExtensionServiceConfig
name: openconnect
environment:
  - OPENCONNECT_SERVER=vpn.example.com
  - OPENCONNECT_PROTOCOL=pulse
  - OPENCONNECT_USER=username
  - OPENCONNECT_PASSWORD=password
```

Then apply the patch to your node's MachineConfigs:

```bash
talosctl patch mc -p @openconnect-config.yaml
```

You can verify that it is in place with the following command:

```bash
talosctl get extensionserviceconfigs

NODE          NAMESPACE   TYPE                     ID           VERSION
mynode        runtime     ExtensionServiceConfig   openconnect  1
```

## Configuration

The extension can be configured through environment variables:

| Variable | Description | Required |
|----------|-------------|----------|
| `OPENCONNECT_SERVER` | VPN server URL (e.g., `vpn.example.com` or `https://vpn.example.com/path`) | Yes |
| `OPENCONNECT_PROTOCOL` | VPN protocol: `anyconnect`, `nc`, `gp`, `pulse`, `f5`, `fortinet`, `array` | No (defaults to anyconnect) |
| `OPENCONNECT_USER` | Username for authentication | No |
| `OPENCONNECT_PASSWORD` | Password for authentication | No |
| `OPENCONNECT_CERTIFICATE` | Path to client certificate file | No |
| `OPENCONNECT_PRIVATE_KEY` | Path to private key file | No |
| `OPENCONNECT_SERVERCERT` | Server certificate fingerprint (e.g., `pin-sha256:ABC123...`) | Recommended |
| `OPENCONNECT_EXTRA_ARGS` | Additional CLI arguments | No |

### Protocol Reference

| Protocol | Value | Description |
|----------|-------|-------------|
| Cisco AnyConnect | `anyconnect` | Default protocol |
| Juniper Network Connect | `nc` | Juniper SSL VPN |
| GlobalProtect | `gp` | Palo Alto Networks GlobalProtect |
| Pulse Secure | `pulse` | Pulse Secure (formerly Juniper) |
| F5 BIG-IP | `f5` | F5 Networks BIG-IP Edge Client |
| Fortinet | `fortinet` | Fortinet SSL VPN |
| Array Networks | `array` | Array Networks SSL VPN |

### Server Certificate Fingerprint

For security, it is recommended to specify the server certificate fingerprint using `OPENCONNECT_SERVERCERT`. This prevents man-in-the-middle attacks.

To obtain the fingerprint, you can run:

```bash
openconnect --authenticate vpn.example.com
```

And note the `pin-sha256:` value shown during certificate verification.

## Troubleshooting

### Check service status

```bash
talosctl services openconnect
```

### View service logs

```bash
talosctl logs openconnect
```

### Check tunnel interface

```bash
talosctl get links
```

### Common Issues

1. **Connection refused**: Ensure the VPN server URL is correct and accessible from the node.

2. **Certificate errors**: Use `OPENCONNECT_SERVERCERT` with the correct fingerprint, or use `OPENCONNECT_EXTRA_ARGS=--no-cert-check` (not recommended for production).

3. **Authentication failures**: Verify username/password or certificate paths are correct.

4. **Protocol mismatch**: Ensure `OPENCONNECT_PROTOCOL` matches your VPN server type.
