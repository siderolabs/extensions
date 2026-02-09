# frr

> [!WARNING]
> This SystemExtension does **not** support dynamic reloading. A configuration change triggers a restart of FRR daemons!

Adds [FRR (Free Range Routing)](https://frrouting.org/) suite as a Talos system extension.

Each routing protocol daemon runs as a separate Talos extension service.

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Available Services

| Service | Description | Depends on |
|---------|-------------|------------|
| `ext-frr-mgmtd` | Centralized config manager | network |
| `ext-frr-zebra` | Kernel routing manager | ext-frr-mgmtd |
| `ext-frr-staticd` | Static routes | ext-frr-zebra |
| `ext-frr-bgpd` | BGP daemon | ext-frr-zebra |
| `ext-frr-bfdd` | BFD daemon | ext-frr-zebra |

All daemons start automatically without configuration.

## Configuration

FRR 10.x uses `mgmtd` as the centralized configuration manager. In this architecture:

- **`mgmtd`** reads config files for zebra and staticd (`zebra.conf`, `staticd.conf`) and pushes configuration to those daemons via internal protocol.
- **`bgpd`** and **`bfdd`** manage their own configuration files directly.

Therefore, `ExtensionServiceConfig` documents should target:

| Config for | ExtensionServiceConfig `name` | Mount path |
|------------|-------------------------------|------------|
| zebra (interfaces, RA, addresses) | `frr-mgmtd` | `/usr/local/etc/frr/zebra.conf` |
| staticd (static routes) | `frr-mgmtd` | `/usr/local/etc/frr/staticd.conf` |
| bgpd (BGP) | `frr-bgpd` | `/usr/local/etc/frr/bgpd.conf` |
| bfdd (BFD) | `frr-bfdd` | `/usr/local/etc/frr/bfdd.conf` |

> [!NOTE]
> Zebra and staticd config files are mounted into the **mgmtd** container, not into the daemon's own container. This is because `mgmtd` is responsible for reading and applying their configuration.

When config is applied, Talos restarts the corresponding service (`ext-frr-mgmtd` for zebra/staticd config, `ext-frr-bgpd` for BGP config).

## Usage

### Basic BGP peering

```yaml
---
apiVersion: v1alpha1
kind: ExtensionServiceConfig
name: frr-bgpd
configFiles:
  - mountPath: /usr/local/etc/frr/bgpd.conf
    content: |
      frr defaults datacenter
      !
      router bgp 65001
        bgp router-id 10.0.0.1
        no bgp default ipv4-unicast
        neighbor eth0 interface remote-as external
        !
        address-family ipv6 unicast
          redistribute connected
          neighbor eth0 activate
        exit-address-family
      exit
```

Apply the patch to your node's MachineConfigs:

```bash
talosctl patch mc -p @frr-config.yaml
```

## Debugging

View per-daemon logs:

```bash
talosctl logs ext-frr-mgmtd
talosctl logs ext-frr-zebra
talosctl logs ext-frr-staticd
talosctl logs ext-frr-bgpd
talosctl logs ext-frr-bfdd
```

### Using vtysh

This extension does not include `vtysh`. To access the FRR CLI, deploy a DaemonSet with the official FRR image and mount the VTY socket directory from the host:

```yaml
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: frr-vtysh
  namespace: kube-system
spec:
  selector:
    matchLabels:
      app: frr-vtysh
  template:
    metadata:
      labels:
        app: frr-vtysh
    spec:
      hostNetwork: true
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
              - matchExpressions:
                  - key: extensions.talos.dev/frr
                    operator: Exists
      containers:
        - name: vtysh
          image: quay.io/frrouting/frr:10.5.1
          command: ["vtysh"]
          stdin: true
          tty: true
          volumeMounts:
            - name: frr-run
              mountPath: /var/run/frr
      volumes:
        - name: frr-run
          hostPath:
            path: /var/run/frr
```

```bash
kubectl attach -it -n kube-system ds/frr-vtysh
```

## Examples

### BGP unnumbered peering with Top of Rack switches

Uses IPv6 Link-Local Addresses (LLA) for BGP sessions — no manual IP addressing on point-to-point links.

```yaml
---
apiVersion: v1alpha1
kind: ExtensionServiceConfig
name: frr-mgmtd
configFiles:
  - mountPath: /usr/local/etc/frr/zebra.conf
    content: |
      frr defaults datacenter
      !
      interface eth0
        no ipv6 nd suppress-ra
      exit
      !
      interface eth1
        no ipv6 nd suppress-ra
      exit
---
apiVersion: v1alpha1
kind: ExtensionServiceConfig
name: frr-bgpd
configFiles:
  - mountPath: /usr/local/etc/frr/bgpd.conf
    content: |
      frr defaults datacenter
      !
      ipv6 prefix-list FABRIC_EXPORT seq 10 permit fd00::/8 ge 128
      ipv6 prefix-list DEFAULT_ONLY_V6 seq 10 permit ::/0
      !
      route-map FABRIC_IN permit 10
        match ipv6 address prefix-list DEFAULT_ONLY_V6
      exit
      !
      route-map FABRIC_OUT permit 10
        match ipv6 address prefix-list FABRIC_EXPORT
      exit
      !
      router bgp 65001
        bgp router-id 10.0.0.1
        bgp ebgp-requires-policy
        no bgp default ipv4-unicast
        neighbor FABRIC peer-group
        neighbor FABRIC remote-as external
        neighbor FABRIC bfd
        neighbor FABRIC capability extended-nexthop
        neighbor eth0 interface v6only peer-group FABRIC
        neighbor eth1 interface v6only peer-group FABRIC
        !
        address-family ipv6 unicast
          redistribute connected
          neighbor FABRIC activate
          neighbor FABRIC route-map FABRIC_IN in
          neighbor FABRIC route-map FABRIC_OUT out
        exit-address-family
      exit
```
