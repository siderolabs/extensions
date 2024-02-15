# Tailscale

Adds https://tailscale.com network interfaces as system extensions.
This means you can access your talos nodes from machines you have configured
with tailscale

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

Configure the extension via `ExtensionServiceConfig` document.

```yaml
---
apiVersion: v1alpha1
kind: ExtensionServiceConfig
name: tailscale
environment:
  - TS_AUTHKEY=<your auth key>
```

```bash
> talosctl apply -n node myconfig.yaml
> talosctl upgrade -n node
```

## Configuration

This extension runs containerboot https://pkg.go.dev/tailscale.com@v1.40.1/cmd/containerboot

Extra tailscale specific environment vars can be configured as needed in `/var/etc/tailscale/auth.env`

Current known env vars are:

- TS_AUTHKEY: the authkey to use for login.
- TS_HOSTNAME: the hostname to request for the node.
- TS_ROUTES: subnet routes to advertise.
- TS_DEST_IP: proxy all incoming Tailscale traffic to the given destination.
- TS_TAILSCALED_EXTRA_ARGS: extra arguments to 'tailscaled'.
- TS_EXTRA_ARGS: extra arguments to 'tailscale up'.
- TS_USERSPACE: run with userspace networking (the default) instead of kernel networking.
- TS_STATE_DIR: the directory in which to store tailscaled state. The data should persist across container restarts.
- TS_ACCEPT_DNS: whether to use the tailnet's DNS configuration.
- TS_KUBE_SECRET: the name of the Kubernetes secret in which to store tailscaled state.
- TS_SOCKS5_SERVER: the address on which to listen for SOCKS5 proxying into the tailnet.
- TS_OUTBOUND_HTTP_PROXY_LISTEN: the address on which to listen for HTTP proxying into the tailnet.
- TS_SOCKET: the path where the tailscaled LocalAPI socket should be created.
- TS_AUTH_ONCE: if true, only attempt to log in if not already logged in. If false (the default, for backwards compatibility), forcibly log in every time the container starts.

### Subnet routing

A pratical example is enabling subnetrouting
```
---
apiVersion: v1alpha1
kind: ExtensionServiceConfig
name: tailscale
environment:
  - TS_AUTHKEY=<your auth key>
  - TS_ROUTES=10.96.0.0/12
```

10.96.0.0/12 is the service subnet talos uses by default (if you use a custom one, you will need to change it).
This allows the k8s services to be available over tailscale (without an ingress controller!).

With this enabled, you can configure tailscales DNS to actually forward certain search domains
to coredns, making it very easy to access k8s services from an external device.
