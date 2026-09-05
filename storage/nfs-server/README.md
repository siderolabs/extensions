# nfs-server

> **NOTE: Security Consideration**
>
> This extension exports filesystems over the network. `rpc.mountd` (port 20048 by default)
> and the kernel NFS server (port 2049) listen on all interfaces. Only install it on nodes
> that are meant to export filesystems, and use host-level firewalling or upstream network
> controls to restrict access.

This extension provides the `nfs-utils` server-side daemons, turning a Talos node into an
NFS server. The client-side tools live in the [`nfs-utils`](../nfs-utils) extension.

## What's Included

- **exportfs**: Syncs the kernel export table (`/var/lib/nfs/etab`) from `/etc/exports`
- **rpc.mountd**: Services NFSv3 `MOUNT` requests and the kernel's export upcalls
- **nfsdcld**: NFSv4 client tracking daemon, required for state recovery across restarts
- **rpc.nfsd**: Starts the in-kernel NFS server threads

## Requirements

- Talos **v1.14.0 or newer**, which provides host-mode extension services.
- The **`nfsd`** extension, for the `nfsd` kernel module. The supervisor waits for the
  module before mounting the host's `nfsd` and `rpc_pipefs` kernel filesystems.
- The **`nfs-utils`** extension, for `rpcbind` (RPC portmapper) and `rpc.statd` (NFSv3 lock
  state monitoring). The NFS server service depends on `ext-rpcbind`.

## Configuration

Supply `/etc/exports` through an `EtcFileConfig`:

```yaml
apiVersion: v1alpha1
kind: EtcFileConfig
name: exports
mode: 0o644
contents: |
  /var/mnt/nfs 10.0.0.0/8(rw,sync,no_subtree_check,fsid=0)
```

This preserves the safer default `root_squash` behavior. Add `no_root_squash` only
when the exported network and every root-capable client are explicitly trusted.

The host-side `libtirpc` also needs `/etc/netconfig`. For an IPv4 server:

```yaml
apiVersion: v1alpha1
kind: EtcFileConfig
name: netconfig
mode: 0o644
contents: |
  udp        tpi_clts      v     inet     udp     -       -
  tcp        tpi_cots_ord  v     inet     tcp     -       -
```

Use a Talos user volume under `/var/mnt` as the export backing store.

`rpc.mountd` listens on TCP and UDP port 20048. NFSv3 clients that do not use rpcbind
discovery should set `mountPort: 20048`; the NFS data port is 2049.

## How It Works

A single host-mode supervisor shares the Talos host's mount namespace and starts the server
components in this order:

1. Mount `nfsd` at `/proc/fs/nfsd` and `rpc_pipefs` at `/var/lib/nfs/rpc_pipefs`.
2. Start `rpc.mountd`, then wait until that exact child process owns port 20048.
3. Start `nfsdcld`, then wait until its database and `rpc_pipefs` inotify watch are open.
4. Populate the kernel export table with `exportfs -r`.
5. Start the kernel NFS server threads with `rpc.nfsd`.
6. Verify that `nfsdcld` attached to the kernel's newly created `nfsd/cld` pipe.

That ordering mirrors upstream's systemd units, where `nfs-server.service` is ordered
`After=nfs-mountd.service nfsdcld.service` and runs `exportfs -r` as its `ExecStartPre`.
The kernel creates the `nfsd/cld` pipe when the server starts, so `nfsdcld` must already be
initialized and watching `rpc_pipefs`; the supervisor then fails startup unless the daemon
attaches to that pipe. Startup probes and utility calls are time-bounded; a failed or hung
step tears down the partial stack and lets Talos restart it. Cleanup is bounded to finish
before Talos's host-process termination grace expires, and child daemons receive a kernel
parent-death signal as a final safeguard.

Talos v1.14 reports a host extension service as running as soon as its supervisor process
starts; the extension service schema has no health probe. That state therefore represents
supervisor liveness, not NFS protocol readiness. Operational readiness is established when
`/proc/fs/nfsd/threads` is nonzero and ports 2049 and 20048 are listening. Failed initialization
is bounded and causes the supervisor to exit rather than remain indefinitely half-started.

The supervisor reloads the kernel export table when `/etc/exports` changes. If either
userspace daemon exits, it first stops the kernel NFS threads and unloads the exports,
then exits so Talos restarts the complete stack in the required order. A normal service
stop or node shutdown performs the same cleanup and forwards termination to both child
daemons.

Nothing else on Talos mounts the `nfsd` or `rpc_pipefs` filesystems. The static supervisor
mounts both idempotently before starting the daemons. Running the NFS server in the host
mount namespace is required because OCI runtimes reject nested mounts under `/proc`.

State persists across reboots in `/var/lib/nfs`.

## Protocol versions

`nfs-utils` is built with `--enable-nfsv4 --enable-nfsv41`, and `rpc.nfsd` is started
without any `--no-nfs-version` flag, so NFSv3, v4.0, v4.1 and v4.2 are all served.

For NFSv4 the export table needs a pseudo-root — give the top-level export `fsid=0`, as in
the example above.

**Client tracking** is handled by `nfsdcld`. Talos kernels are built with
`CONFIG_NFSD_LEGACY_CLIENT_TRACKING` unset, so neither the usermode-helper nor the
`/var/lib/nfs/v4recovery` fallback is compiled in — `nfsdcld` is the only client-tracking
method the kernel can use. Without it `nfsd` still starts (its return value is discarded in
`nfs4_state_start_net()`) and logs `NFSD: Unable to initialize client recovery tracking!`,
but `nfsd4_client_record_check()` then returns `-EOPNOTSUPP` and NFSv4 clients cannot
reclaim opens or locks after a server restart.

**ID mapping** needs nothing: `rpc.idmapd` is not shipped, and the kernel server defaults
to `nfsd.nfs4_disable_idmapping=1`, so AUTH_SYS clients exchange numeric UIDs/GIDs.
Kerberos (`--disable-gss`) and NFS junctions (`--disable-junction`) are not built.

## References

- [exportfs man page](https://linux.die.net/man/8/exportfs)
- [rpc.mountd man page](https://linux.die.net/man/8/rpc.mountd)
- [rpc.nfsd man page](https://linux.die.net/man/8/rpc.nfsd)
- [nfsdcld man page](https://linux.die.net/man/8/nfsdcld)
- [exports(5)](https://linux.die.net/man/5/exports)
