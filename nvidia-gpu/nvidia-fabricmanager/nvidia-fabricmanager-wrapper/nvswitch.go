// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

// #cgo CFLAGS: -I./rdma-core/include/
// #cgo LDFLAGS: -L./rdma-core/lib/statics/ -l:libibumad.a
// #include <linux/types.h> /* __be64 */
// #include <infiniband/umad.h>
import "C"
import (
	"bytes"
	"encoding/binary"
	"os"
	"path"
	"unsafe"
)

type NVSwitchMgmtPort struct {
	IBDevice string
	PortGUID uint64
}

/*
Find InfiniBand devices with the capability to configure NVSwitches.
---
From: https://docs.nvidia.com/datacenter/tesla/pdf/fabric-manager-user-guide.pdf

The CX7 bridge device is integrated into the GPU baseboard, which includes two physical ports. Each port exposes one
physical function (FC PF) and one Limited physical function (LPF) to the host system, which totals four PFs. The PFs
are categorized into the following PFs:
  - Limited PFs (LPF) are designated for specific tasks in the system.
    They are used by the FM and the NVLSM to configure and set up NVSwitches, GPU, and NVLink routing information.
    LPFs are also used by telemetry agents, such as NVIBDM and DCGM, to monitor and collect data. Resetting this
    PF with FLR also resets the corresponding NVSwitch device.

To differentiate between LPFs and FC PFs, the LPF VPD information includes a vendor-specific field called SMDL, with
a non-zero value defined as SW_MNG. For bare-metal, full pass-through, and shared NVSwitch deployments, the prelaunch
script in the FM service unit file will run and query the available CX7 devices for this VPD information. The file
populates the required FM and NVLSM configuration values so that these communication entities can access the relevant
devices.
*/
func findLpfDevices() (devices []string) {
	const ibPath = "/sys/class/infiniband"

	devDir, err := os.ReadDir(ibPath)
	if err != nil {
		return
	}

	for _, device := range devDir {
		vpd, err := os.ReadFile(path.Join(ibPath, device.Name(), "device/vpd"))
		if err != nil {
			continue
		}

		if bytes.Contains(vpd, []byte("SMDL=SW_MNG")) {
			devices = append(devices, device.Name())
		}
	}
	return
}

func findNvswitchMgmtPorts() (ports []NVSwitchMgmtPort) {
	lpfDevs := findLpfDevices()
	if len(lpfDevs) == 0 {
		return
	}

	if C.umad_init() < 0 {
		return
	}

	for _, lpf := range lpfDevs {
		const maxPorts = 16
		var portGUIDs [maxPorts]C.__be64

		/*
			$ man 3 umad_get_ca_portguids

			On success, umad_get_ca_portguids() returns a non-negative value equal to the number of port GUIDs actually
			filled. Not all filled entries may be valid. Invalid entries will be 0. For example, on a CA node with only
			one port, this function returns a value of 2. In this case, the value at index 0 will be invalid as it is
			reserved for switches. On failure, a negative value is returned.
		*/
		numPort := C.umad_get_ca_portguids(C.CString(lpf), &portGUIDs[0], maxPorts)

		for i := range int(numPort) {
			var guid uint64

			// convert kernel __be64 to uint64
			buf := bytes.NewReader((*[8]byte)(unsafe.Pointer(&portGUIDs[i]))[:])
			if err := binary.Read(buf, binary.BigEndian, &guid); err != nil {
				continue
			}

			if guid != 0 {
				ports = append(ports, NVSwitchMgmtPort{lpf, guid})
			}
		}
	}

	C.umad_done()
	return
}
