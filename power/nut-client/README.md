# nut-client extension

## Usage

Enable the extension in the machine configuration before installing Talos:

```yaml
machine:
  install:
    extensions:
      - image: ghcr.io/siderolabs/nut-client:<VERSION>
```

Configure the extension via .machine.files 
You must replace upsmonHost and upsmonPasswd to match configuration on your nut server.  
See [upsd.users](https://networkupstools.org/docs/man/upsd.users.html) man page for details.


On Talos SHUTDOWNCMD must be `/sbin/poweroff`

```yaml
machine:
  files:
    - path: /var/etc/nut/upsmon.conf
      permissions: 0o600
      op: create
      content: |-
        MONITOR ${upsmonHost} 1 remote ${upsmonPasswd} slave
        SHUTDOWNCMD "/sbin/poweroff"

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
