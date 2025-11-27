# gpio-pinctrl extension

This extension provides Intel GPIO/Pinctrl drivers for Apollo Lake platforms.

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Included Modules

- `pinctrl-intel` - Intel pinctrl core driver
- `pinctrl-broxton` - Intel Apollo Lake (Broxton) pinctrl driver

## Usage

Enable the modules in your machine configuration:

```yaml
machine:
  kernel:
    modules:
      - name: pinctrl-intel
      - name: pinctrl-broxton
```

## Verification

Check that modules are loaded:

```bash
talosctl -n <node-ip> read /proc/modules | grep pinctrl
```

Expected output:
```
pinctrl_broxton 24576 0 - Live 0x...
pinctrl_intel 40960 1 pinctrl_broxton, Live 0x...
```

Verify GPIO devices are available:

```bash
talosctl -n <node-ip> ls -la /dev/gpiochip*
```

Check dmesg for driver initialization:

```bash
talosctl -n <node-ip> dmesg | grep -i 'pinctrl.*INT3452'
```

## Testing with libgpiod

From a privileged pod with `/dev` mounted:

```bash
gpiodetect
gpioinfo
```

## Troubleshooting

### Module not loading

Ensure the extension is properly installed and the modules are listed in your machine config.

### No /dev/gpiochip* devices

Verify Talos is running a kernel built with `CONFIG_GPIO_CDEV=y`. Check dmesg for pinctrl driver initialization messages.

### ABI/signature errors

Ensure the extension was built against the same kernel version as your Talos installation.
