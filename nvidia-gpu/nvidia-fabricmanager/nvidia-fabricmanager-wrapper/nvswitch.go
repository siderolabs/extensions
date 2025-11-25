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
	"errors"
	"log"
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
			log.Printf("nvidia-fabricmanager-wrapper: infiniband: found NVSwitch LPF device: %s\n", device.Name())
			devices = append(devices, device.Name())
		}
	}
	return
}

/*
Replication of the logic found in script "nvidia-fabricmanager-start.sh" from the FabricManager package:

	for each LPF device:
		for each LPF device port:
			skip port if the IsSMDisabled bit is set
			read the port GUID
			configure FM and SM to use that port
			stop looking for a management port

Unlike the upstream algorithm, we'll continue scanning all ports after a first management port is found to generate
useful log messages for bug reports.
*/
func findNvswitchMgmtPort() (err error, port *NVSwitchMgmtPort) {
	lpfDevs := findLpfDevices()
	if len(lpfDevs) == 0 {
		return errors.New("no NVSwitch LPF device found"), nil
	}

	if C.umad_init() < 0 {
		return errors.New("failed to initialize libibumad"), nil
	}
	defer C.umad_done()

	for _, device := range lpfDevs {
		// device name as a C string to use with libibumad
		cDeviceName := C.CString(device)
		defer C.free(unsafe.Pointer(cDeviceName))

		/*
		 * get IB device attributes
		 */

		caPtr := (*C.struct_umad_ca)(C.malloc(C.sizeof_struct_umad_ca))
		defer C.free(unsafe.Pointer(caPtr))

		if C.umad_get_ca(cDeviceName, caPtr) < 0 {
			log.Printf("nvidia-fabricmanager-wrapper: infiniband: failed to get interface attributes device=%s\n", device)
			continue
		}

		numPorts := int(caPtr.numports)
		log.Printf("nvidia-fabricmanager-wrapper: infiniband: successful read of interface attributes device=%s"+
			" numPorts=%d\n", device, numPorts)

		/*
		 * iterate over device ports
		 */

		portPtr := (*C.struct_umad_port)(C.malloc(C.sizeof_struct_umad_port))
		defer C.free(unsafe.Pointer(portPtr))

		// index 0 is not a valid port per IB specifications
		for portIdx := 1; portIdx <= numPorts; portIdx++ {
			if C.umad_get_port(cDeviceName, C.int(portIdx), portPtr) < 0 {
				log.Printf("nvidia-fabricmanager-wrapper: infiniband: failed to get port attributes device=%s port=%d\n",
					device, portIdx)
				continue
			}

			// read port GUID, we have to convert kernel __be64 to uint64
			var portGUID uint64
			buf := bytes.NewReader((*[8]byte)(unsafe.Pointer(&portPtr.port_guid))[:])
			if err := binary.Read(buf, binary.BigEndian, &portGUID); err != nil {
				log.Printf("nvidia-fabricmanager-wrapper: infiniband: failed to convert port GUID endianness device=%s port=%d\n",
					device, portIdx)
				continue
			}

			// read port capabilities
			const IsSMDisabledMask = 0x00000400
			var capMask uint32
			buf = bytes.NewReader((*[4]byte)(unsafe.Pointer(&portPtr.capmask))[:])
			if err := binary.Read(buf, binary.BigEndian, &capMask); err != nil {
				log.Printf("nvidia-fabricmanager-wrapper: infiniband: failed to convert port capmask endianness device=%s port=%d\n",
					device, portIdx)
				continue
			}
			isSMDisabled := (capMask & IsSMDisabledMask) != 0

			// read port state
			portState := uint32(portPtr.state)
			portStateStr := "Unknown"
			switch portState {
			case 1:
				portStateStr = "Down"
			case 2:
				portStateStr = "Init"
			case 4:
				portStateStr = "Active"
			}

			log.Printf("nvidia-fabricmanager-wrapper: infiniband: successful read of port attributes device=%s port=%d"+
				" guid=0x%x capabilities=0x%x isSMDisabled=%t state=%s\n", device, portIdx, portGUID, capMask, isSMDisabled,
				portStateStr)

			// still looking for a management port
			if port == nil {
				// evaluate candidate port
				if portGUID != 0 && isSMDisabled == false {
					port = &NVSwitchMgmtPort{device, portGUID}
					log.Printf("nvidia-fabricmanager-wrapper: infiniband: selected NVSwitch management port device=%s guid=0x%x\n", device, portGUID)

				}
			}
		}
	}

	if port == nil {
		return errors.New("failed to find a NVSwitch management port"), nil
	}

	return
}
