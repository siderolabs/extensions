# wpa-supplicant

Runs [`wpa_supplicant`](https://w1.fi/wpa_supplicant/) as a system extension so a
Talos node can join a WPA2/WPA3-Personal (PSK/SAE) WiFi network. Credentials are
supplied at runtime via an `ExtensionServiceConfig`, so changing SSID/PSK does
**not** require rebuilding the image.

## Prerequisites

WiFi on Talos needs kernel support that stock images may not ship. This
extension only provides the supplicant on top of it:

1. **`cfg80211` and your adapter's in-tree driver.** `wpa_supplicant` talks to
   the kernel's wireless stack over `nl80211`, which requires `cfg80211` plus the
   in-tree driver module for your WiFi adapter. Talos curates its kernel module
   list and may omit wireless modules, so you likely need a **custom Talos image
   that includes `cfg80211` and your driver** (add the relevant kernel modules,
   e.g. `kernel/net/wireless/cfg80211.ko` and your driver `.ko`, and rebuild).
   Without them, no `wlan*` interface appears.
2. **The `net.ifnames=0` kernel argument** so the interface is deterministically
   `wlan0` (the service is hardcoded to `-i wlan0`).

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).
Build your Talos image with this `wpa-supplicant` extension and
`--extra-kernel-arg net.ifnames=0`, on a kernel that includes `cfg80211` and your
WiFi driver.

## Usage

Configure the extension via an `ExtensionServiceConfig` document that delivers
`wpa_supplicant.conf`, plus a `KernelModuleConfig` to load the driver. The
module `name` must be the exact name `modprobe` expects (e.g. `iwlwifi`):

```yaml
---
apiVersion: v1alpha1
kind: KernelModuleConfig
name: iwlwifi
---
apiVersion: v1alpha1
kind: ExtensionServiceConfig
name: wpa-supplicant
configFiles:
  - content: |
      country=US
      ctrl_interface=/run/wpa_supplicant
      update_config=1
      network={
        ssid="Your Network"
        psk="your-password"
        key_mgmt=WPA-PSK SAE
      }
    mountPath: /etc/wpa_supplicant/wpa_supplicant.conf
```

Apply it to your node's MachineConfig:

```bash
talosctl patch mc -p @wifi.yaml
```

Also enable DHCP on the WiFi link:

```yaml
machine:
  network:
    interfaces:
      - interface: wlan0
        dhcp: true
```

Verify it is in place:

```bash
talosctl get extensionserviceconfigs

NODE     NAMESPACE   TYPE                     ID               VERSION
mynode   runtime     ExtensionServiceConfig   wpa-supplicant   1
```

## Configuration

The `configFiles` content is a standard
[`wpa_supplicant.conf`](https://man.archlinux.org/man/wpa_supplicant.conf.5).

## Verify

```bash
talosctl --nodes <node> service ext-wpa-supplicant status   # STATE: Running
talosctl --nodes <node> logs ext-wpa-supplicant             # association log
talosctl --nodes <node> get links wlan0                     # operationalState: up
talosctl --nodes <node> get addresses | grep wlan0          # DHCP lease
```

A successful association looks like:

```
wlan0: Associated with <bssid>
wlan0: WPA: Key negotiation completed [PTK=CCMP GTK=CCMP]
wlan0: CTRL-EVENT-CONNECTED - Connection to <bssid> completed
```
