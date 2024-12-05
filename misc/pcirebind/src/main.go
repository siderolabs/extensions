package main

import (
	"fmt"
	"os"
	"strings"
)

type rebind struct {
	id        string
	oldDriver string
	newDriver string
}

// writeToSysFile writes to the specified sysfs file
func writeToSysFile(path, content string) error {
	file, err := os.OpenFile(path, os.O_WRONLY, 0200) // 0200 is write-only permission
	if err != nil {
		e := fmt.Errorf("failed to open %s: %v", path, err)
		fmt.Printf("[error] %v", e)

		return e
	}
	defer file.Close()

	if _, err = file.WriteString(content); err != nil {
		e := fmt.Errorf("failed to write to %s: %v", path, err)
		fmt.Printf("[error] %v", e)

		return e
	}

	return nil
}

// parseKernelCmdline parses specially formatted strings out of /proc/cmdline
//
// The format is as follows:
// pcirebind.rebind=<pci_bus_id>_<current_driver>+<new_driver>
//
// example:
//
//	pcirebind.rebind=0000:04:00.0_ixgbe+vfio-pci
func parseKernelCmdline(readFunc func(string) ([]byte, error)) ([]rebind, error) {
	data, err := readFunc("/proc/cmdline")
	if err != nil {
		return nil, fmt.Errorf("failed to read /proc/cmdline: %v", err)
	}

	cmdline := string(data)
	parts := strings.Fields(cmdline)

	var rebindsList []rebind

	for _, part := range parts {
		if strings.HasPrefix(part, "pcirebind.rebind=") {
			// Extract the rebind specification
			spec := strings.TrimPrefix(part, "pcirebind.rebind=")

			// Split the specification into id, oldDriver, and newDriver
			parts := strings.Split(spec, "_")
			if len(parts) != 2 {
				fmt.Printf("Invalid rebind format: %s\n", spec)
				continue
			}
			id, drivers := parts[0], parts[1]

			driverParts := strings.Split(drivers, "+")
			if len(driverParts) != 2 {
				fmt.Printf("Invalid drivers format: %s\n", drivers)
				continue
			}
			oldDriver, newDriver := driverParts[0], driverParts[1]

			// Append the parsed rebind information to the list
			rebindsList = append(rebindsList, rebind{
				id:        id,
				oldDriver: oldDriver,
				newDriver: newDriver,
			})
		}
	}

	if len(rebindsList) > 0 {
		return rebindsList, nil
	}

	return nil, fmt.Errorf("no rebinds found in kernel command line")
}

// overrideDriver sets the `driver_override` for a given PCI Bus ID
// this is analogous to:
//
//	echo "vfio-pci" > /sys/bus/pci/devices/0000:04:00.0/driver_override
func (r *rebind) overrideDriver() error {
	bindPath := fmt.Sprintf("/sys/bus/pci/devices/%s/driver_override", r.id)
	return writeToSysFile(bindPath, r.newDriver)
}

// unbindDriver writes to `unbind` a given PCI Bus ID
// this is analogous to:
//
// echo "0000:04:00.0" > /sys/bus/pci/drivers/ixgbe/unbind
func (r *rebind) unbindDriver() error {
	bindPath := fmt.Sprintf("/sys/bus/pci/drivers/%s/unbind", r.oldDriver)
	return writeToSysFile(bindPath, r.id)
}

// bindDriver writes to `bind` a given PCI Bus ID
// this is analogous to:
//
// echo "0000:04:00.0" > /sys/bus/pci/drivers/vfio-pci/bind
func (r *rebind) bindDriver() error {
	bindPath := fmt.Sprintf("/sys/bus/pci/drivers/%s/bind", r.newDriver)
	return writeToSysFile(bindPath, r.id)
}

// v8 was last working version
func main() {
	rebinds, err := parseKernelCmdline(os.ReadFile)
	if err != nil {
		fmt.Printf("Error parsing kernel command line: %v\n", err)

		os.Exit(1)
	}

	anyError := false

	for _, rebind := range rebinds {
		if err := rebind.overrideDriver(); err != nil {
			fmt.Printf("Error writing to `driver-override`: %v\n", err)

			continue
		}

		if err := rebind.unbindDriver(); err != nil {
			fmt.Printf("Error writing to `unbind`: %v\n", err)

			continue
		}

		if err := rebind.bindDriver(); err != nil {
			fmt.Printf("Error writing to `bind`: %v\n", err)

			anyError = true

			continue
		}
	}

	if anyError == true {
		os.Exit(1)
	}
}
