# YubiHSM Connector

This extension provides the [YubiHSM Connector](https://developers.yubico.com/yubihsm-connector/) service, which acts as an HTTP-to-USB bridge for YubiHSM 2 hardware security modules.

The connector allows applications to communicate with YubiHSM 2 devices via a local HTTP API instead of requiring direct USB access.

## Prerequisites

- YubiHSM 2 device connected via USB to the Talos node

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

Once installed, the yubihsm-connector service starts automatically and listens on `127.0.0.1:12345` by default.

### Configuring Network Access

By default, the connector only listens on localhost for security. To allow network access from pods or other machines, configure the extension via `ExtensionServiceConfig`:

```yaml
---
apiVersion: v1alpha1
kind: ExtensionServiceConfig
name: yubihsm-connector
configFiles:
  - content: |
      listen: "0.0.0.0:12345"
      syslog: "false"
    mountPath: /usr/local/etc/yubihsm-connector/yubihsm-connector.yaml
```

Apply the configuration:

```bash
talosctl patch mc -p @yubihsm-connector-config.yaml
```

Verify the configuration is applied:

```bash
talosctl get extensionserviceconfigs
```

### Configuration Options

The connector configuration file supports the following options:

| Option | Description | Default |
|--------|-------------|---------|
| `listen` | Address and port to listen on | `127.0.0.1:12345` |
| `syslog` | Enable syslog output | `false` |
| `serial` | YubiHSM device serial (if multiple devices) | (auto-detect) |
| `cert` | Path to TLS certificate for HTTPS | (none) |
| `key` | Path to TLS private key for HTTPS | (none) |

### Verifying the Service

Check that the service is running:

```bash
talosctl services ext-yubihsm-connector
```

View service logs:

```bash
talosctl logs ext-yubihsm-connector
```

### Using from Pods

To access the YubiHSM from Kubernetes pods, you can connect to the connector's HTTP API at `http://<node-ip>:12345` (after configuring network access).

For applications using the YubiHSM SDK, set the connector URL:

```bash
export YUBIHSM_CONNECTOR_URL="http://<node-ip>:12345"
```

## Security Considerations

- **Default localhost binding**: The connector binds to localhost only by default. Only expose it to the network if necessary.
- **No authentication**: The connector does not provide authentication. Restrict network access using firewall rules or only bind to trusted interfaces.
- **USB device access**: The extension requires read-write access to `/dev` for USB device communication.

## Troubleshooting

### Service fails to start

Check the logs for errors:

```bash
talosctl logs ext-yubihsm-connector
```

Common issues:
- **"no such file or directory"**: Usually indicates a library loading issue. Ensure the extension image is properly built with musl libc.
- **USB device not found**: Verify the YubiHSM 2 is connected and recognized by the kernel (`talosctl dmesg | grep -i yubi`).

### Device not detected

Ensure udevd is running and the USB device is connected:

```bash
talosctl services udevd
talosctl dmesg | grep -i usb
```

## References

- [YubiHSM 2 Documentation](https://developers.yubico.com/YubiHSM2/)
- [YubiHSM Connector Documentation](https://developers.yubico.com/yubihsm-connector/)
- [YubiHSM Shell](https://developers.yubico.com/yubihsm-shell/)
