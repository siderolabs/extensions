# bird2
>
> [!WARNING]
> This SystemExtensio does **not** support dynamic reloading. A configuration change triggers a restart of bird!

Adds [bird2](https://bird.network.cz/) routing daemon as Talos system extensions.

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

1. Configure the extension via `ExtensionServiceConfig` document.

```yaml
---
apiVersion: v1alpha1
kind: ExtensionServiceConfig
name: bird2
configFiles:
  - mountPath: /usr/local/etc/bird.conf
    content: |
      # This is just a minimal NOOP config!
      log stderr all;
      router id 6.6.6.6;
      protocol device {}
```

1. Apply the patch to your node's MachineConfigs

```bash
talosctl patch mc -p @bird2-config.yaml
```

## Debugging

A Bird client can be spawned by a node debug container, e.g.: `chroot /proc/$(pgrep bird)/cwd /usr/local/sbin/birdcl`

## Examples

### Dual Top of Rack uplink with CNI integration
>
> [!WARNING]
> This is a lab configuration, review and adjust to your needs before production. For this to work, ToR switches must advertise the default route.

This bird2 configuration peers with a top of rack switches, listens on localhost for BGP sessions established by a CNI (here [Cilium](https://cilium.io)) and redistributes them within the network fabric:

```yaml
---
apiVersion: v1alpha1
kind: ExtensionServiceConfig
name: bird2
configFiles:
  - mountPath: /usr/local/etc/bird.conf
    content: |
      define FABRIC_AS = 65000;
      define TALOS_NODE_AS = 65001;
      define CILIUM_AS = 65002;

      log stderr all;
      debug protocols all;

      router id from "dummy0";

      protocol device {
        scan time 10;
      }

      protocol direct loopback {
        interface "dummy0";
        ipv4 {
          import all;
          export none;
        };
      }

      protocol kernel {
        merge paths on;
        learn off;
        ipv4 {
          export filter {
            if proto = "cilium" then reject;
            if proto = "loopback" then reject;
            accept;
          };
          import none;
        };
      }

      protocol bgp fabric_enp33s0f1np1 {
        local as TALOS_NODE_AS;
        source address 10.66.66.6;
        neighbor 10.66.66.1 as FABRIC_AS;
        ipv4 {
          import all;
          export all;
          next hop self;
        };
        interface "enp33s0f1np1";
      };

      protocol bgp cilium {
        passive on;
        multihop 2;  # bypass bird statup check for localhost IP address

        local as TALOS_NODE_AS;
        neighbor 127.0.0.1 as CILIUM_AS;

        ipv4 {
          import all;
          export none;
        };
      };
```

Cilium configuration (just for demonstration purposes!):

```yaml
---
apiVersion: cilium.io/v2
kind: CiliumBGPPeerConfig
metadata:
  name: cilium-bird2-nodelocal
spec:
  families:
    - afi: ipv4
      safi: unicast
      advertisements:
        matchLabels:
          bgp: cilium-peer
---
apiVersion: cilium.io/v2
kind: CiliumBGPAdvertisement
metadata:
  name: pod-cidr-advertisement
  labels:
    bgp: cilium-bird2-nodelocal
spec:
  advertisements:
    - advertisementType: "PodCIDR"
---
apiVersion: cilium.io/v2
kind: CiliumBGPClusterConfig
metadata:
  name: bird2-nodelocal-bgp
spec:
  nodeSelector: {}
  bgpInstances:
    - name: "bird2-nodelocal-bgp-65002"
      localASN: 65002
      peers:
        - name: "bird2-nodelocal"
          peerASN: 65001
          peerAddress: 127.0.0.1
          peerConfigRef:
            name: "cilium-bird2-nodelocal"
```
