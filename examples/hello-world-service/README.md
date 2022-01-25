# Example Talos Extension Service

This repository is an example of an extension service.

## Usage

Enable the extension in the machine configuration before installing Talos:

```yaml
machine:
  install:
    extensions:
      - image: ghcr.io/siderolabs/hello-world-service:<VERSION>
```

Once this example extension is installed, it will provide simple HTTP server which responds with a message on port 80:

```bash
$ curl http://<IP>/
Hello from Talos Linux Extension Service!
```

Extension service appears in the service list (please note the `ext-` prefix):

```bash
$ talosctl services
NODE         SERVICE           STATE     HEALTH   LAST CHANGE   LAST EVENT
172.20.0.5   apid              Running   OK       1m37s ago     Health check successful
172.20.0.5   containerd        Running   OK       1m38s ago     Health check successful
172.20.0.5   cri               Running   OK       1m37s ago     Health check successful
172.20.0.5   ext-hello-world   Running   ?        1m38s ago     Started task ext-hello-world (PID 1100) for container ext-hello-world
172.20.0.5   kubelet           Running   OK       1m30s ago     Health check successful
172.20.0.5   machined          Running   ?        1m40s ago     Service started as goroutine
172.20.0.5   udevd             Running   OK       1m38s ago     Health check successful
```

Run `talosctl service ext-hello-world` to see the detailed service state:

```bash
$ talosctl service ext-hello-world
NODE     172.20.0.5
ID       ext-hello-world
STATE    Running
HEALTH   ?
EVENTS   [Running]: Started task ext-hello-world (PID 1100) for container ext-hello-world (2m47s ago)
         [Preparing]: Creating service runner (2m47s ago)
         [Preparing]: Running pre state (2m47s ago)
         [Waiting]: Waiting for service "containerd" to be "up" (2m48s ago)
         [Waiting]: Waiting for service "containerd" to be "up", network (2m49s ago)
```

The service can be started and stopped via `talosctl`:

```bash
$ talosctl  service ext-hello-world stop
NODE         RESPONSE
172.20.0.5   Service "ext-hello-world" stopped
$ talosctl service ext-hello-world start
NODE         RESPONSE
172.20.0.5   Service "ext-hello-world" started
```

Use `talosctl logs` to access service logs:

```bash
$ talosctl logs ext-hello-world
172.20.0.5: 2022/02/16 18:45:21 starting the hello world service
172.20.0.5: 2022/02/16 18:52:33 stopping the hello world service
172.20.0.5: 2022/02/16 18:52:35 starting the hello world service
```
