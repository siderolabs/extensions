# usb-modem-drivers extension


## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

Enable the modules in Talos Linux machine configuration, your modem might not require
all of them, so feel free to remove the ones that are not needed:

```yaml
machine:
  kernel:
    modules:
      - name: usbserial
      - name: option
      - name: cdc_mbim
      - name: qmi_wwan
      - name: usb_wwan
      - name: usbnet
      - name: ax88796b
      - name: asix
      - name: ax88179_178a
      - name: cdc_ether
      - name: cdc_ncm
      - name: ipheth
      - name: net1080
      - name: cdc_subset
      - name: zaurus
      - name: cdc_wdm
      - name: r8153_ecm
```

## Verifiying

You can verify the modules are enabled by reading the `/proc/modules` where it _should_ show the module is live.

For example:

```
❯ talosctl -n 192.168.32.5  read /proc/modules
usbnet 20480 - - Live 0xffffffffc01c9000 (O)
cdc_wdm 94208 - - Live 0xffffffffc01aa000 (O)
...
...
```

In addition, if your USB module is installed and you have a USB modem connected to the system, you should be able to verify it's presence at `/dev/cdc-wdm1` or see a new network interface at `wwan1`.

For example:

```
❯ talosctl -n 192.168.32.5  ls -l /dev/cdc-wdm1
NODE           MODE          UID   GID   SIZE(B)   LASTMOD           NAME
192.168.32.5   Dcrw-rw----   0     44    0         Sep 10 18:15:52   /dev/cdc-wdm1
```

or

```
❯ talosctl -n 192.168.32.5 dmesg
# look for lines like these:
 kern:    info: [2023-06-07T16:40:10.189868521Z]: usbcore: registered new interface driver usbhid
 kern:    info: [2023-06-07T16:40:10.190066521Z]: option 6-1.2:1.3: GSM modem (1-port) converter detected
 kern:    info: [2023-06-07T16:40:10.265654521Z]: usb 6-1.2: GSM modem (1-port) converter now attached to ttyUSB7
 kern:    info: [2023-06-07T16:40:10.280291521Z]: qmi_wwan 6-1.2:1.4: cdc-wdm1: USB WDM device
```

## Configurtion

Talos currently does not have an OS level method of dialing up cellular modems, such as `ModemManager` or `uqmi`, so you will
probably want to ignore it on the system level, pass the modems `/dev/` devices to a pod, and dial it up there using a tool such
as `ModemManager` or `uqmi`.

First, ignore the interface on the Talos level, so Talos doesn't try to bring it up:

```yaml
machine:
  network:
    interfaces:
      - interface: wwan1 # The modem interface name, if you have several modems, you can optionally use deviceSelector instead.
        ignore: true # ignore the interface.
```

Next, you can create a pod that might contain something like this:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: ModemManager
  namespace: modems # you might have to adjust your security settings for this namespace
spec:
  hostNetwork: true
  nodeName: modems-server
  containers:
  - name: ubuntu-container
    image: ubuntu:latest
    command: ["/bin/bash", "-c", "--"]
    args: ["while true; do sleep 30; done;"]
    securityContext:
      privileged: true
    volumeMounts:
    - name: dev
      mountPath: /dev/
  volumes:
  - name: dev
    hostPath:
      path: /dev/ttyUSB1 # you might want to mount all of the /dev/ttyUSB* and /dev/cdc-wdm* devices
      type: Directory
```

Now the modem should be available to you normally on the pod.
