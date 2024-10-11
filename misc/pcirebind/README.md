## pcirebind

This is a Talos Linux Extension that can be used to rebind the driver on a given PCI Bus ID. It is used internally to unbind NICs from `ixgbe` and apply the `vfio-pci` driver to enable VPP to take control of the devices.

#### Applying to a Talos Linux server
After embedding the extension in your Talos installer you'll need to modify the kernel arguments.

Add additional kernel args under `.machine.install.extraKernelArgs` in the format of:

``` yaml
# 0000:04:00.00 is the PCI Bus ID
# ixgbe is the driver to unbind
# vfio-pci is the drive to bind
machine:
  install:
    extensions:
      extraKernelArgs:
        - pcirebind.rebind=0000:04:00.00_ixgbe+vfio-pci
        - pcirebind.rebind=0000:04:00.01_ixgbe+vfio-pci
```

Then trigger a reboot with `upgrade`:

``` sh
talosctl -e <endpoint> -n <node> --talosconfig=./talosconfig upgrade
```

Once the server reboots you can check the status of the service with the following commands:

``` sh
talosctl -e <endpoint> -n <node> --talosconfig=./talosconfig service ext-pcirebind
NODE     10.50.12.211
ID       ext-pcirebind
STATE    Finished
HEALTH   ?
EVENTS   [Finished]: Service finished successfully (2m36s ago)
         [Running]: Started task ext-pcirebind (PID 4210) for container ext-pcirebind (2m37s ago)
         [Preparing]: Creating service runner (2m37s ago)
         [Preparing]: Running pre state (2m37s ago)
         [Waiting]: Waiting for file "/sys/bus/pci/drivers/vfio-pci/bind" to exist (2m42s ago)
         [Waiting]: Waiting for service "containerd" to be "up", file "/sys/bus/pci/drivers/vfio-pci/bind" to exist (2m44s ago)
         [Starting]: Starting service (2m44s ago)
         
talosctl -e <endpoint> -n <node> --talosconfig=./talosconfig logs ext-pcirebind
```

