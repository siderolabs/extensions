// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// nvidia-vfio-cdi writes the CDI spec for vfio-bound NVIDIA GPUs to /run/cdi.
//
// /run/cdi is tmpfs, so this runs at every boot (restart: untilSuccess —
// a non-zero exit retries until the GPU is bound to vfio-pci, which happens
// early in boot via machine.kernel.modules vfio_pci ids=...).
//
// The spec is cdev-only on purpose: each GPU maps exclusively to its IOMMUFD
// character device /dev/vfio/devices/vfioN. Including the legacy group node
// (/dev/vfio/<group>) makes Kata pick the legacy VFIO backend and fail
// confidential guests with "ConfidentialGuest needs IOMMUFD". The cdev name
// is also the only device name registered: it is unique per device, whereas
// aliasing IOMMU-group ids or running indexes into the same CDI namespace
// can resolve a request to the wrong GPU on a multi-GPU node.
//
// It also guards the extension-dependency chain: extensions cannot declare
// dependencies on each other, so this is the loud, runtime end of the
// kata-containers -> kata-containers-snp -> this add-on pairing contract.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	drvDir  = "/sys/bus/pci/drivers/vfio-pci"
	outDir  = "/run/cdi"
	outFile = "/run/cdi/nvidia-vfio.yaml"
)

// Files other extensions in the chain provide and this add-on does not:
// the shim and virtiofsd come from kata-containers, the SNP OVMF from
// kata-containers-snp. The GPU config's virtio_fs_daemon and firmware
// paths resolve to two of them, and no kata handler starts without the
// shim.
var baseFiles = []struct {
	path, extension string
}{
	{"/usr/local/bin/containerd-shim-kata-v2", "kata-containers"},
	{"/usr/local/libexec/virtiofsd", "kata-containers"},
	{"/usr/local/share/ovmf/AMDSEV.fd", "kata-containers-snp"},
}

func run() error {
	for _, f := range baseFiles {
		if _, err := os.Stat(f.path); err != nil {
			return fmt.Errorf("extension %s is not installed on this node (%s missing): install it from the same release as this add-on — kata-qemu-nvidia-gpu-snp pods cannot start without it", f.extension, f.path)
		}
	}

	if _, err := os.Stat(drvDir); err != nil {
		return fmt.Errorf("vfio-pci driver not present yet")
	}

	devs, err := filepath.Glob(filepath.Join(drvDir, "*:*:*.*"))
	if err != nil {
		return err
	}

	var spec strings.Builder
	spec.WriteString("cdiVersion: \"0.5.0\"\nkind: nvidia.com/pgpu\ndevices:\n")

	count := 0

	for _, dev := range devs {
		vendor, err := os.ReadFile(filepath.Join(dev, "vendor"))
		if err != nil || strings.TrimSpace(string(vendor)) != "0x10de" {
			continue
		}

		bdf := filepath.Base(dev)

		// IOMMUFD cdev name (requires VFIO_DEVICE_CDEV in the host kernel).
		entries, err := os.ReadDir(filepath.Join(dev, "vfio-dev"))
		if err != nil || len(entries) == 0 {
			fmt.Fprintf(os.Stderr, "%s: no vfio-dev cdev (kernel lacks IOMMUFD/CDEV?)\n", bdf)

			continue
		}

		cdev := entries[0].Name()

		fmt.Fprintf(&spec, "  - name: %q\n    containerEdits:\n      deviceNodes:\n        - path: /dev/vfio/devices/%s\n", cdev, cdev)
		fmt.Printf("%s -> cdev=%s\n", bdf, cdev)

		count++
	}

	if count == 0 {
		return fmt.Errorf("no vfio-bound NVIDIA GPU found yet - retrying")
	}

	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return err
	}

	// Write-then-rename so containerd never reads a partial spec; a failed
	// rename must fail the service (untilSuccess retries) rather than
	// exit 0 with no spec on the node.
	tmp := outFile + ".tmp"
	if err := os.WriteFile(tmp, []byte(spec.String()), 0o644); err != nil {
		return err
	}

	if err := os.Rename(tmp, outFile); err != nil {
		return err
	}

	fmt.Printf("wrote %s (%d GPU(s))\n", outFile, count)

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		os.Exit(1)
	}
}
