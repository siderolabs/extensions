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

Create the configuration under `/var` with a Talos machine configuration patch, then set `URUNC_CONFIG_FILE` so both `urunc` and its containerd shim use it:

```yaml
machine:
  files:
    - path: /var/etc/urunc/config.toml
      op: create
      content: |
        [log]
        level = "info"
        syslog = false

        [extra_binaries.virtiofsd]
        path = "/usr/local/bin/virtiofsd"
        options = "--cache always --sandbox none"
---
apiVersion: v1alpha1
kind: EnvironmentConfig
variables:
  URUNC_CONFIG_FILE: /var/etc/urunc/config.toml
```

Talos restricts user-created machine files to `/var`, while urunc defaults to `/etc/urunc/config.toml`.
The patch writes the file to the allowed location and overrides the default path; otherwise urunc logs a warning and uses defaults.

The extension installs monitor binaries under `/usr/local/bin`; set their paths in the TOML when overriding monitor configuration.
A complete example is available in the upstream [`config.toml`](https://github.com/urunc-dev/urunc/blob/v0.8.0/deployment/urunc-deploy/config.toml).

## Downstream patch

The extension applies one patch to urunc v0.8.0. It reads the nameserver from the container's OCI-mounted `/etc/resolv.conf` and passes it to Unikraft.
This allows Kubernetes service DNS to work inside the guest instead of always using the hard-coded `8.8.8.8` fallback.
