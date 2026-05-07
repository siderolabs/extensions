# uhid

This system extension provides the `uhid` kernel module for Talos Linux.

The uhid module exposes `/dev/uhid`, which allows user-space applications to create virtual HID devices. This is required by game streaming software such as [Sunshine](https://github.com/LizardByte/Sunshine) to emulate gamepad and controller input received from Moonlight clients.

## Usage

Enable the extension in your Talos machine config:

```yaml
machine:
  kernel:
    modules:
      - name: uhid
```

## Compatibility

Requires `CONFIG_UHID=m` in the Talos kernel configuration.
