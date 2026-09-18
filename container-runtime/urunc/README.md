# urunc extension

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

This extension installs `urunc`, `containerd-shim-urunc-v2`, the monitor binaries, and a containerd runtime handler named `urunc`.

Apply the following manifest to add the runtime class:

```yaml
apiVersion: node.k8s.io/v1
kind: RuntimeClass
metadata:
  name: urunc
handler: urunc
```

## Configuration

On Talos Linux 1.14 and later, provide `/etc/urunc/config.toml` with an [`EtcFileConfig`](https://docs.siderolabs.com/talos/v1.14/reference/configuration/runtime/etcfileconfig):

```yaml
apiVersion: v1alpha1
kind: EtcFileConfig
name: urunc/config.toml
mode: 0o644
contents: |
  [log]
  level = "info"
  syslog = false

  [extra_binaries.virtiofsd]
  path = "/usr/local/bin/virtiofsd"
  options = "--cache always --sandbox none"
```

The `name` is relative to `/etc`, so the document writes the file to urunc's default configuration path without an environment override.

The extension installs monitor binaries under `/usr/local/bin`; set their paths in the TOML when overriding monitor configuration.
A complete example is available in the upstream [`config.toml`](https://github.com/urunc-dev/urunc/blob/v0.8.0/deployment/urunc-deploy/config.toml).

## Downstream patch

The extension applies one patch to urunc v0.8.0. It reads the nameserver from the container's OCI-mounted `/etc/resolv.conf` and passes it to Unikraft.
This allows Kubernetes service DNS to work inside the guest instead of always using the hard-coded `8.8.8.8` fallback.
