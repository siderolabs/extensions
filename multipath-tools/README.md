## Multipath-tools
This extension adds multipath-tools to Talos. Mutipath is an essential feature in many enterprise environments because it enables high availability and fault tolerance of storage subsystems. Typically these storage subsystems are connected through a Fiber Channel network, but other methods like iSCSI exist as well. This extension will start the multpath daemon as a Talos service and use a default config file packaged with this extension.

### Customization
The default configuration file is located in src/rootfs/etc/multipath.conf, you may wish to update this to suit the needs of your particular storage subsystem. In particular the load balancing algorithm in there right now is round-robin. This is due to Talos not yet having additional kernel modules available for more advanced load balancing algorithms. In the future this may change and you may want to customize this part of the configuration.

### Limitations
The multipath bindings file in this extension is currently ephemeral. If the Talos node reboots the bindings file will be recreated on startup. This could be troublesome in some envrionments. Future versions of this extension may address this limitation.
