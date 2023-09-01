# Tailscale

Adds xe-guest-utilities as system extensions.

## Installation

Simplest install
```
machine:
  install:
    extensions:
      - image: ghcr.io/siderolabs/xe-guest-utilities:7.20.2
```

```
> talosctl apply -n node myconfig.yaml
> talosctl upgrade -n node
```

## Configuration

This extension requires no configuration.
