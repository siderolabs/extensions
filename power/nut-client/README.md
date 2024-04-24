# nut-client extension

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

Configure the extension via `ExtensionServiceConfig` document.
You must replace upsmonHost and upsmonPasswd to match configuration on your nut server.
See [upsd.users](https://networkupstools.org/docs/man/upsd.users.html) man page for details.

On Talos SHUTDOWNCMD must be `/sbin/poweroff`

```yaml
---
apiVersion: v1alpha1
kind: ExtensionServiceConfig
name: nut-client
configFiles:
  - content: |-
        MONITOR ${upsmonHost} 1 remote ${upsmonPasswd} slave
        SHUTDOWNCMD "/sbin/poweroff"
    mountPath: /usr/local/etc/nut/upsmon.conf
```

Then apply the patch to your node's MachineConfigs
```bash
talosctl patch mc -p @nut-config.yaml
```

You will then be able to verify that it is in place with the following command
```bash
talosctl get extensionserviceconfigs

NODE          NAMESPACE   TYPE                     ID           VERSION
mynode   runtime     ExtensionServiceConfig   nut-client   1
```

## Testing

Confirm extension service is running

```bash
$ talosctl service ext-nut-client
NODE     192.168.1.1
ID       ext-nut-client
STATE    Running
HEALTH   ?
EVENTS   [Running]: Started task ext-nut-client (PID 2263) for container ext-nut-client (59m59s ago)
         [Preparing]: Creating service runner (59m59s ago)
         [Preparing]: Running pre state (59m59s ago)
         [Waiting]: Waiting for service "cri" to be "up" (59m59s ago)
         [Waiting]: Waiting for service "cri" to be "up", network (1h0m0s ago)
         [Waiting]: Waiting for service "cri" to be registered, network (1h0m1s ago)
         [Waiting]: Waiting for service "containerd" to be "up", service "cri" to be registered, network (1h0m2s ago)
         [Waiting]: Waiting for service "containerd" to be "up", service "cri" to be "up", network (1h0m3s ago)
```

**CAUTION** this will power off all connected systems.

Trigger a 'Full System Shutdown' on the nut-server

```bash
# upsmon -c fsd
```

all connected upsmon clients should perform a full shutdown and power off.
