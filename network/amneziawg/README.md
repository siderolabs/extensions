# AmneziaWG

Provides the [AmneziaWG](https://github.com/amnezia-vpn/amneziawg-linux-kernel-module)
kernel module and userspace as a Talos system extension, bringing up an `awg0`
interface at boot — before `etcd`/`kubelet` — so it can serve as a
DPI-resistant inter-node mesh underlay.

AmneziaWG is a fork of WireGuard that adds traffic-shape obfuscation (junk
packets, header transformation) so the on-wire format isn't trivially
fingerprintable by middleboxes that selectively impair standard WireGuard.

This extension is the **kernel-module** approach (native datapath). A
userspace approach (`amneziawg-go`) is tracked separately in #1004; see that
PR's discussion for the trade-off (datapath performance vs. per-kernel
rebuild).

## How it works

The service container runs `awg-up`, a small shell-free Go helper (no bash,
no `iproute2` — it does its own netlink for link/addr/route and execs the
bundled `awg` C binary only for `awg setconf`). This mirrors the
Go-binary-entrypoint pattern of the other `network/` extensions
(`tailscale`, `nebula`, …) rather than shipping a shell + `wg-quick` script.

`awg-up` parses a `wg-quick`-style config, creates `awg0` (`type amneziawg`),
applies the AmneziaWG obfuscation parameters via `awg setconf`, and installs
a route per `AllowedIPs` prefix. It is idempotent across restarts.

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

Requires the `amneziawg-pkg` from `siderolabs/pkgs` (kernel module + `awg`).
Load the module via machine config:

```yaml
machine:
  kernel:
    modules:
      - name: amneziawg
```

## Usage

Deliver the per-node config via an `ExtensionServiceConfig` document:

```yaml
---
apiVersion: v1alpha1
kind: ExtensionServiceConfig
name: amneziawg
configFiles:
  - content: |
      [Interface]
      PrivateKey = <node private key>
      Address    = 192.0.2.1/24
      ListenPort = 51820
      # AmneziaWG obfuscation parameters — must match across the mesh
      Jc = 4
      Jmin = 40
      Jmax = 70
      S1 = 50
      S2 = 100
      H1 = 1
      H2 = 2
      H3 = 3
      H4 = 4

      [Peer]
      PublicKey  = <peer public key>
      Endpoint   = <peer public ip>:51820
      AllowedIPs = 192.0.2.2/32
      PersistentKeepalive = 25
    mountPath: /var/run/amneziawg/awg0.conf
```

> **Note:** pick the interface subnet outside the cluster's pod and service
> CIDRs (the default `serviceSubnets` `10.96.0.0/12` spans `10.96`–`10.111`).
> Addresses inside those ranges are filtered out of node-address discovery.
> See siderolabs/talos#13469.

## awg-up

The Go helper sources live in [`awg-up/`](./awg-up). Build:
`CGO_ENABLED=0 go build -ldflags="-s -w" -trimpath`. Unit tests:
`go test ./...`.
