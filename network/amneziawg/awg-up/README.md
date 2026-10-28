# awg-up

Small Go helper that brings up an AmneziaWG interface from a `wg-quick`-style
`.conf` file. Used as the entrypoint of the `ext-amneziawg` system service
container.

## What it does

Sequence on each boot:

1. `ip link add awg0 type amneziawg` (idempotent — if the interface already
   exists from a previous run, just reuse it)
2. `ip address add` each `[Interface] Address` line on the interface
3. `ip link set mtu` (defaults to 1420)
4. `ip link set up`
5. `awg setconf awg0 -` with the wg-format part of the config piped on stdin
   — this is where the AmneziaWG obfuscation parameters (`Jc`, `Jmin`,
   `Jmax`, `S1`, `S2`, `H1..H4`) get applied alongside the standard
   WireGuard `PrivateKey` / `ListenPort` / per-peer fields
6. `ip route add` for each peer's `AllowedIPs` prefix, scoped to the interface

That's it — exit 0. The kernel state persists; we don't need a keepalive
loop in the service container.

## Why not 100% pure Go?

We shell out to the bundled `awg` C binary for step 5 only. `awg` (~50 KB
in the upstream tarball) is the canonical implementation of the AmneziaWG
netlink-genl protocol and tracks every upstream change to the obfuscation
parameter encoding. `golang.zx2c4.com/wireguard/wgctrl` speaks stock
WireGuard netlink but not AmneziaWG's extension fields, so a Go-native
path would mean shipping our own parallel netlink-attribute table and
keeping it in lockstep with the C tool on every release.

50 KB and one `exec.Command` is a better trade than that maintenance burden.

If `wgctrl` grows AmneziaWG support upstream — or if `amnezia-vpn/amneziawg-tools`
ships a Go library mode — we'd drop the exec and link the package directly.

## Why not shell out for ip link / addr / route too?

Those go via `vishvananda/netlink`, which is what the rest of the Talos
ecosystem (machined, kubelet, CNI plugins) already uses. Adding `iproute2`
to the bundle would balloon the rootfs and pull in `glibc` /  `musl` / etc.
The netlink calls are syscall-thin and don't need a userspace tool.

## Building

```
CGO_ENABLED=0 go build -ldflags="-s -w" -o awg-up .
```

Strips to ~3-4 MB static binary. Fits in the extension's `scratch` rootfs
alongside the `awg` C binary (~50 KB) — same shape as `tailscale`'s
container (which bundles three Go binaries plus a userspace iptables for
its CNI integration).

## Testing

```
go test ./...
```

Unit tests cover the `.conf` parser (config_test.go) — wg-format vs.
wg-quick fields, AmneziaWG extension keys, AllowedIPs splitting,
malformed-input handling. Integration tests for the netlink ops require
the kernel module + privileged container and live in the parent
extension's CI (`siderolabs/extensions/.github/workflows/...`).

## Files

| File | Responsibility |
|---|---|
| `main.go` | argv parsing, top-level orchestration |
| `config.go` | `.conf` parser, wg-format serialiser |
| `link.go` | `ip link add` / addr / mtu / up via netlink |
| `setconf.go` | exec `awg setconf` with config on stdin |
| `route.go` | `ip route add` for each AllowedIPs prefix |
