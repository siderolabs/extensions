# joydev

This system extension provides the `joydev` kernel module for Talos Linux.

The joydev module provides the Linux joystick interface at `/dev/input/js*`. This interface is used by the browser Gamepad API and legacy joystick applications. Game streaming software such as [Sunshine](https://github.com/LizardByte/Sunshine) relies on joydev to translate evdev gamepad events into the joystick API format expected by streaming clients like Moonlight.

## Usage

Enable the extension in your Talos machine config:

```yaml
machine:
  kernel:
    modules:
      - name: joydev
```

## Compatibility

Requires `CONFIG_INPUT_JOYDEV=m` in the Talos kernel configuration.
