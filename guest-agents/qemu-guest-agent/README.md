# qemu-guest-agent extension

## Usage

Enable the extension in the machine configuration before installing Talos:

```yaml
machine:
  install:
    extensions:
      - image: ghcr.io/siderolabs/qemu-guest-agent:<VERSION>
```

## Testing

Confirm extension service is running

```bash
$ talosctl service ext-qemu-guest-agent
NODE     192.168.1.1
ID       ext-qemu-guest-agent
STATE    Running
HEALTH   ?
EVENTS   [Running]: Started task ext-qemu-guest-agent (PID 1941) for container ext-qemu-guest-agent (4s ago)
         [Preparing]: Creating service runner (4s ago)
         [Preparing]: Running pre state (4s ago)
         [Waiting]: Waiting for service "cri" to be "up" (5s ago)
         [Waiting]: Waiting for service "containerd" to be "up", service "cri" to be registered, file "/dev/virtio-ports/org.qemu.guest_agent.0" to exist (6s ago)
         [Waiting]: Waiting for service "containerd" to be registered, service "cri" to be registered, file "/dev/virtio-ports/org.qemu.guest_agent.0" to exist (8s ago)
         [Waiting]: Waiting for service "containerd" to be "up", service "cri" to be "up", file "/dev/virtio-ports/org.qemu.guest_agent.0" to exist (9s ago)
```

### Proxmox

In the VM “Options” tab, ensure “QEMU Guest Agent” is set to “Enabled”. If it was not already enabled, you will need to reboot the VM for the `/dev/virtio-ports/org.qemu.guest_agent.0` device to be available.

The “IPs” field in the VM “Summary” tab should now be populated and the “Shutdown” button in the Proxmox UI should start an orderly shutdown, e.g.

```
192.168.1.1: user: warning: [2023-06-28T16:38:41.270046585Z]: [talos] shutdown via API received. actor id: b72e2882-9e06-4626-8422-bb2a7410f0ea
192.168.1.1: user: warning: [2023-06-28T16:38:41.287740585Z]: [talos] shutdown sequence: 10 phase(s)
192.168.1.1: user: warning: [2023-06-28T16:38:41.288028585Z]: [talos] phase drain (1/10): 1 tasks(s)
192.168.1.1: user: warning: [2023-06-28T16:38:41.288401585Z]: [talos] task cordonAndDrainNode (1/1): starting
...
```
