# nfs-utils

This extension provides `rpcbind` and `rpc.statd` daemons for NFSv3 file locking support.

## What's Included

- **rpcbind**: Converts RPC program numbers into universal addresses, required for RPC-based services
- **rpc.statd**: NSM (Network Status Monitor) service daemon that notifies NFS peers of host restarts
- **sm-notify**: Companion utility for rpc.statd

## Use Case

These services are essential for:
- NFSv3 mounts with file locking support
- Kubernetes environments using Trident or other NFS-based storage provisioners
- Any scenario requiring POSIX file locking on NFSv3 shares

## How It Works

The extension runs two containerized services:

1. **rpcbind**: Starts first and provides the RPC portmapper service
2. **rpc.statd**: Depends on rpcbind and provides NFS lock state monitoring

Both services start automatically and persist state across reboots in `/var/lib/nfs/statd` and `/var/lib/rpcbind`.

## References

- [rpcbind man page](https://linux.die.net/man/8/rpcbind)
- [rpc.statd man page](https://linux.die.net/man/8/rpc.statd)
- [Related Trident issue](https://github.com/NetApp/trident/issues/806#issuecomment-2399332314)
