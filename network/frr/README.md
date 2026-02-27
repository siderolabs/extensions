# FRR (Free Range Routing) Extension for Talos

This extension provides FRR for BGP routing on Talos hosts, with built-in MetalLB VRF integration for advertising Kubernetes LoadBalancer IPs via BGP.

## Overview

### Purpose

1. **FRR (Free Range Routing)** for BGP routing on Talos hosts
2. **MetalLB VRF integration** with a veth pair for Kubernetes LoadBalancer IP advertisement
3. **Private IPv6 point-to-point connection** (`fd00::/8`) between FRR and MetalLB speaker
4. **Dynamic configuration** via Jinja2 template (`frr.conf.j2`) rendered using `j2cli`
5. **Interface discovery** from MAC addresses specified in `FE_MACS` environment variable

### Architecture

- **FRR runs in host network namespace** and manages both:
  - Fabric-facing BGP peering (eBGP with leaf switches via physical interfaces)
  - MetalLB-facing BGP peering (iBGP-style in VRF `metallb`)

- **MetalLB speaker** runs with `hostNetwork: true`, connects to FRR via a veth pair:
  - `veth-metallb` interface: in **host namespace** - MetalLB speaker binds here
  - `veth-frr` interface: in **VRF metallb** - FRR listens here for MetalLB connections

- **Route flow**: MetalLB advertises LoadBalancer IPs → FRR VRF BGP → imported to default VRF → advertised to fabric

### Components

| Component | Description |
|-----------|-------------|
| `frr-startup.sh` | Bash script that runs before FRR to set up VRF, veth pair, IPv6 addresses, MTU, and generate `frr.conf` |
| `frr.conf.j2` | Jinja2 template for FRR configuration with variables from environment |
| `daemons` | FRR daemons configuration (zebra, bgpd, staticd, bfdd enabled) |

### Startup Script Responsibilities

1. **Validate mandatory environment variables** (`NODE_IP`, `ASN_LOCAL`, and either `FE_MACS` or `FE_PORT_NAMES`)
2. **Resolve interface names** from MAC addresses in `FE_MACS` (CSV format)
3. **Auto-detect MTU** from first fabric interface (if `INTERFACE_MTU` not set)
4. **Create VRF `metallb`** with configurable routing table ID (default: 88)
5. **Create veth pair**:
   - `veth-frr` assigned to VRF `metallb` with `PEER_IP_FRR` (default: `fda1:b2c3:d4e5:0001::0/127`)
   - `veth-metallb` in host namespace with `PEER_IP_METALLB` (default: `fda1:b2c3:d4e5:0001::1/127`)
6. **Set MTU** on veth pair from `INTERFACE_MTU` (auto-detected or user-provided)
7. **Generate `/etc/frr/frr.conf`** from Jinja2 template using `jinja2-cli`
8. **Exec to FRR's `docker-start`** script to launch FRR daemons

## Network Topology

```
┌──────────────────────────────────────────────────────────────────┐
│                         Talos Host                                │
│                                                                   │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │                     Default VRF (Host)                       │ │
│  │                                                              │ │
│  │  ┌──────────────────────────┐                               │ │
│  │  │   MetalLB BGP Speaker    │                               │ │
│  │  │   (hostNetwork: true)    │                               │ │
│  │  │   ASN: 4000099999        │                               │ │
│  │  │                          │                               │ │
│  │  │   Binds to veth-metallb  │                               │ │
│  │  │   fda1:...:0001::1/127   │                               │ │
│  │  └────────────┬─────────────┘                               │ │
│  │               │                                              │ │
│  │   veth-metallb (in host ns)                                  │ │
│  │   fda1:...:0001::1/127                                       │ │
│  │               │ veth                                         │ │
│  └───────────────┼──────────────────────────────────────────────┘ │
│                  │                                                │
│  ┌───────────────┼──────────────────────────────────────────────┐ │
│  │               │         VRF: metallb (table 88)              │ │
│  │               ▼                                               │ │
│  │   veth-frr (in VRF)                                          │ │
│  │   fda1:...:0001::0/127                                       │ │
│  │               │                                               │ │
│  │   ┌───────────▼─────────────────────────────────────────┐    │ │
│  │   │              FRR                                     │    │ │
│  │   │   router bgp 4000099998 vrf metallb                  │    │ │
│  │   │   - Listens passively on veth-frr:179                │    │ │
│  │   │   - Receives LB IP routes from MetalLB               │    │ │
│  │   │                                                      │    │ │
│  │   │   router bgp 65001 (default VRF)                     │    │ │
│  │   │   - Peers with fabric via FE_MACS/FE_PORT_NAMES ifaces│    │ │
│  │   │   - Imports routes from VRF metallb                  │    │ │
│  │   │   - Advertises LB IPs to fabric                      │    │ │
│  │   └──────────────────────────────────────────────────────┘    │ │
│  └───────────────────────────────────────────────────────────────┘ │
│                                                                   │
│  Physical Interfaces (resolved from FE_MACS or from FE_PORT_NAMES)│
│  ├── eth0 ──► Leaf Switch 1 (eBGP unnumbered, IPv6 link-local)   │
│  └── eth1 ──► Leaf Switch 2 (eBGP unnumbered, IPv6 link-local)   │
└───────────────────────────────────────────────────────────────────┘
```

## Host Networking Architecture

This extension supports advertising host IPs from loopback and dummy interfaces to enable:

1. **Host Traffic HA/Load Balancing via ECMP** - Multiple equal-cost paths to host IPs
2. **Anycast Kubernetes API** - Control plane nodes advertise a shared VIP via `dummy1`

### Interface IP Advertisement

FRR advertises connected routes from these interfaces to fabric switches:

| Interface | Purpose |
|-----------|---------|
| `lo` | Primary loopback IP (node identity) |
| `dummy0` | Additional service IPs for the host |
| `dummy1` | Anycast IPs (e.g., Kubernetes API VIP on control plane nodes) |

This is configured via the `loopbacks` route-map which matches and redistributes connected routes from `lo`, `dummy0`, and `dummy1`.

### ECMP Load Balancing

When multiple paths exist to a destination (e.g., dual-homed hosts with two fabric uplinks), FRR enables ECMP via:

```
bgp bestpath as-path multipath-relax
```

This allows traffic to be load-balanced across multiple equal-cost paths, providing both redundancy and increased bandwidth.

### Anycast Kubernetes API

For highly available Kubernetes API access:

1. Configure `dummy1` on all control plane nodes with the same anycast VIP (e.g., `10.0.0.100/32`)
2. FRR advertises this IP from all control plane nodes
3. Fabric switches see multiple paths to the VIP and load-balance traffic via ECMP
4. If a control plane node fails, its route is withdrawn and traffic flows to remaining nodes

```
┌─────────────────────────────────────────────────────────────────────┐
│                         Leaf/Spine Fabric                           │
│                                                                     │
│   Routes to 10.0.0.100/32:                                          │
│   ├── via Control Plane 1 (eth0)                                    │
│   ├── via Control Plane 2 (eth0)                                    │
│   └── via Control Plane 3 (eth0)                                    │
│                                                                     │
│   ECMP distributes API traffic across all healthy control planes    │
└─────────────────────────────────────────────────────────────────────┘
         │                    │                    │
         ▼                    ▼                    ▼
┌─────────────────┐  ┌─────────────────┐  ┌─────────────────┐
│ Control Plane 1 │  │ Control Plane 2 │  │ Control Plane 3 │
│                 │  │                 │  │                 │
│ dummy1:         │  │ dummy1:         │  │ dummy1:         │
│ 10.0.0.100/32   │  │ 10.0.0.100/32   │  │ 10.0.0.100/32   │
│                 │  │                 │  │                 │
│ kube-apiserver  │  │ kube-apiserver  │  │ kube-apiserver  │
└─────────────────┘  └─────────────────┘  └─────────────────┘
```

### Configuring Dummy Interfaces on Talos

Add dummy interfaces via Talos machine config:

```yaml
machine:
  network:
    interfaces:
      - interface: dummy0
        dummy: true
        addresses:
          - 10.0.1.1/32  # Host IP
      - interface: dummy1
        dummy: true
        vip:
          ip: 10.0.0.100  # Anycast K8s API VIP (control plane only)
```

## Environment Variables

| Variable | Required | Default | Description |
|----------|----------|---------|-------------|
| `NODE_IP` | **Yes** | - | Node's primary IPv4 address (BGP router-id) |
| `ASN_LOCAL` | **Yes** | - | Local BGP AS number for fabric peering |
| `FE_MACS` | No* | - | Comma-separated MAC addresses of fabric-facing interfaces |
| `FE_PORT_NAMES` | No* | - | Comma-separated interface names of fabric-facing interfaces |
| `ASN_VRF_METALLB` | No | `4000099998` | BGP ASN for FRR's MetalLB VRF session |
| `ASN_METALLB_SPEAKER` | No | `4000099999` | BGP ASN for MetalLB speaker |
| `VRF_METALLB` | No | `metallb` | Name of the MetalLB VRF |
| `METALLB_VRF_ROUTE_TABLE_ID` | No | `88` | Linux routing table ID for VRF |
| `PEER_IP_FRR` | No | `fda1:b2c3:d4e5:0001::0` | FRR-side IPv6 address (in VRF) |
| `PEER_IP_METALLB` | No | `fda1:b2c3:d4e5:0001::1` | MetalLB-side IPv6 address (in host ns) |
| `PEER_IP_PREFIX` | No | `127` | IPv6 prefix length for p2p link |
| `INTERFACE_MTU` | No | Auto-detect from first fabric interface | MTU for veth pair |
| `ENABLE_BFD` | No | `true` | Enable BFD for fabric BGP peers |
| `FRR_PROFILE` | No | `datacenter` | FRR defaults profile (`datacenter` or `traditional`) |

\* Either `FE_MACS` or `FE_PORT_NAMES` is required (not both).

## Usage

### Extension Configuration

Configure the FRR extension via `ExtensionServiceConfig`:

```yaml
---
apiVersion: v1alpha1
kind: ExtensionServiceConfig
name: frr
environment:
  - NODE_IP=10.0.0.1
  - ASN_LOCAL=65001
  - FE_MACS=00:11:22:33:44:55,00:11:22:33:44:56
```

Or using interface names directly:

```yaml
---
apiVersion: v1alpha1
kind: ExtensionServiceConfig
name: frr
environment:
  - NODE_IP=10.0.0.1
  - ASN_LOCAL=65001
  - FE_PORT_NAMES=eth0,eth1
```

### MetalLB Configuration

Configure MetalLB to peer with FRR via the veth interface:

```yaml
apiVersion: metallb.io/v1beta2
kind: BGPPeer
metadata:
  name: frr-peer
  namespace: metallb-system
spec:
  myASN: 4000099999
  peerASN: 4000099998
  peerAddress: fda1:b2c3:d4e5:0001::0
  sourceAddress: fda1:b2c3:d4e5:1::1
```

### Example: Advertising LoadBalancer IPs

1. MetalLB speaker (with `hostNetwork: true`) binds to `veth-metallb` interface
2. MetalLB initiates BGP connection to FRR at `fda1:b2c3:d4e5:0001::0:179`
3. FRR's VRF BGP process (`router bgp 4000099998 vrf metallb`) accepts the connection
4. MetalLB advertises LoadBalancer IPs to FRR
5. FRR imports routes from VRF metallb into the default VRF
6. FRR advertises the LoadBalancer IPs to fabric switches via eBGP

## FRR Daemons

The following FRR daemons are enabled:

- **zebra** - Routing table manager
- **bgpd** - BGP routing daemon
- **staticd** - Static route daemon
- **bfdd** - BFD (Bidirectional Forwarding Detection) daemon

## Building

This extension is built using the Talos extensions build system:

```bash
make frr
```

## License

This extension is licensed under the Mozilla Public License 2.0.
FRR is licensed under GPL-2.0-or-later.
