# sysbox extension

[Sysbox](https://github.com/nestybox/sysbox) is a container runtime that runs
_system containers_: containers that can run Docker, Kubernetes or systemd
inside, without being privileged. It does so by always placing the container in
a user-namespace and by partially virtualizing `procfs` and `sysfs`.

Sysbox v0.7.0 is the first release that supports containerd (v2.0.5 or later)
as the container manager, which is what Talos uses. Earlier releases required
CRI-O.

Sysbox has three components. The extension provides all of them, but only
`sysbox-runc` is used directly by Talos:

| component     | ships as                     | runs as                                        |
| ------------- | ---------------------------- | ---------------------------------------------- |
| `sysbox-runc` | `/usr/local/bin/sysbox-runc` | OCI runtime, invoked on the host by containerd |
| `sysbox-mgr`  | `/usr/local/bin/sysbox-mgr`  | daemon, run by the `sysbox` static pod         |
| `sysbox-fs`   | `/usr/local/bin/sysbox-fs`   | daemon, run by the `sysbox` static pod         |

The two daemons cannot run as Talos extension services: those always get a
private PID namespace (Talos only shares the host's *network* namespace with
them, and the spec has no option for the others), while `sysbox-mgr` and
`sysbox-fs` act on the PIDs of the containers `sysbox-runc` creates on the host.
Seccomp notifications and `pidfd_open(2)` only resolve those in the host PID
namespace. `setns(2)` cannot help here either: a process may only join a
*descendant* PID namespace, never an ancestor. So the daemons run in a
`hostPID: true` static pod instead, executing the binaries this extension
installs on the host.

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Machine configuration

Sysbox requires unprivileged user-namespaces to be allowed on the node, which
Talos disables by default, and pods have to opt into user namespaces:

```yaml
machine:
  sysctls:
    user.max_user_namespaces: "11255"
  kubelet:
    extraArgs:
      feature-gates: UserNamespacesSupport=true
cluster:
  apiServer:
    extraArgs:
      feature-gates: UserNamespacesSupport=true
```

The feature gates are only needed on Kubernetes older than v1.33, where user
namespaces are enabled by default.

## Running the daemons

[`sysbox-static-pod.yaml`](sysbox-static-pod.yaml) is a machine configuration
patch that runs `sysbox-mgr` and `sysbox-fs` as a static pod. Apply it to the
nodes that have the extension installed:

```bash
talosctl -n <node> patch mc --patch @sysbox-static-pod.yaml
```

A static pod keeps Sysbox a per-node concern: it only runs where the extension
is installed, and comes up with the kubelet rather than depending on a
controller. The pod is privileged, uses the host PID namespace, and bind mounts
`/var` and `/run` with bidirectional propagation, so that the Sysbox data store
and the `sysbox-fs` FUSE mount are visible to `sysbox-runc` on the host. It
mounts the binaries from the host rather than shipping its own. The image is
only used as a rootfs, so any small image works. `/etc` inside it has to be
writeable because `sysbox-mgr` writes the `sysbox` subuid/subgid ranges there.
Talos' own `/etc` is read only.

Check that both containers are up:

```bash
talosctl -n <node> containers -k | grep sysbox
```

## Usage

Create the runtime class for the handler registered by the extension:

```yaml
apiVersion: node.k8s.io/v1
kind: RuntimeClass
metadata:
  name: sysbox-runc
handler: sysbox-runc
```

Pods then select it, and must request a user namespace with `hostUsers: false`:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: syscont
spec:
  runtimeClassName: sysbox-runc
  hostUsers: false
  containers:
    - name: syscont
      image: nestybox/ubuntu-noble-systemd-docker
```

## Testing

The pod should be up and running:

```bash
$ kubectl get pods
NAME      READY   STATUS    RESTARTS   AGE
syscont   1/1     Running   0          40s
```

Inside it, the container's root user owns the user namespace, and can run a
nested container engine:

```bash
$ kubectl exec syscont -- cat /proc/self/uid_map
         0     165536      65536
$ kubectl exec syscont -- docker run --rm hello-world
```

## Notes on the Talos integration

* the three binaries are built from source with upstream's static targets
  (`make static`, `make sysbox-fs-static`, `make sysbox-mgr-static`), so they
  depend on nothing on the host or in the static pod. Upstream reads the
  architecture and the build tags off the build host's `uname` rather than off
  the node the binaries end up on, so [`pkg.yaml`](pkg.yaml) sets them itself;
* `sysbox-ipc` ships no generated gRPC stubs. They are built with `protoc` and
  the pre-v1.4 `protoc-gen-go`, the last one to implement the `plugins=grpc`
  generator sysbox's protobuf makefiles call;
* `sysbox-mgr` refuses to start unless `rsync`, `modprobe` and `iptables` are
  installed, and uses `rsync` at runtime to sync the data stores in and out of
  the containers. `modprobe` and `iptables` are part of Talos. This extension
  builds `rsync` and installs it to `/usr/local/bin`;
* `sysbox-fs` mounts its per-container FUSE filesystem by executing
  `fusermount3`. The extension builds it (statically linked, so that the pod
  needs no library from the host) and installs it as
  `/usr/local/lib/sysbox/fusermount3`. It does not go in `/usr/local/bin`,
  because the `fuse3` extension installs its own copy there and two extensions
  cannot ship the same path;
* Talos has no `shiftfs`; Sysbox uses ID-mapped mounts instead, which the Talos
  kernel supports.

## Updating

Renovate tracks `SYSBOX_VERSION` in [`../vars.yaml`](../vars.yaml), but bumping
it does not update the sources. Sysbox keeps its code in git submodules and
GitHub source archives do not carry submodule contents, so every `SYSBOX_*_REV`
is pinned to the commit the release tag points at, and has to be re-resolved
with it.

```bash
VERSION=0.7.0

# SYSBOX_RUNC_REV, SYSBOX_FS_REV, SYSBOX_MGR_REV, SYSBOX_IPC_REV, SYSBOX_LIBS_REV
curl -s "https://api.github.com/repos/nestybox/sysbox/git/trees/v${VERSION}" |
    jq -r '.tree[] | select(.mode == "160000") | "\(.path) \(.sha)"'

# SYSBOX_BAZIL_REV: sysbox-fs nests one more submodule, nestybox's fork of
# bazil.org/fuse, which it builds against as ./bazil
curl -s "https://api.github.com/repos/nestybox/sysbox-fs/git/trees/${SYSBOX_FS_REV}" |
    jq -r '.tree[] | select(.mode == "160000") | "\(.path) \(.sha)"'
```

`sysbox-pkgr` and `sysbox-dockerfiles` hold the deb and rpm packaging and the
upstream test images, neither of which this extension builds. A submodule that
shows up in those listings without a matching `_REV` variable is a source
[`pkg.yaml`](pkg.yaml) is missing.

The checksums follow from the revisions, so

```bash
make update-checksums
```

recomputes them from the `_REV` and `_VERSION` variables changed in the working
tree. The same command covers a `protoc`, `libfuse` or `rsync` bump.

`SYSBOX_PROTOC_GEN_GO_VERSION` has to stay on the v1.3 series. v1.4 dropped the
`--go_out=plugins=grpc` generator that sysbox's protobuf makefiles call.
