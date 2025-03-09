// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/vishvananda/netlink"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s <DEVNAME> <OFFSET>", os.Args)
	}

	devName := string(os.Args[1])
	offset, err := strconv.ParseInt(os.Args[2], 10, 64)
	if err != nil {
		log.Fatalf("Invalid offset: %v", err)
	}

	if _, err := os.Stat("/sys/firmware/devicetree/base/hat/custom_5"); os.IsNotExist(err) {
		os.Exit(1)
	}

	data, err := os.ReadFile("/sys/firmware/devicetree/base/hat/custom_5")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	baseMAC := string(data)

	if baseMAC == "" {
		os.Exit(2)
	}

	MAC := strings.Split(baseMAC, ":")
	lastNum, err := strconv.ParseInt(MAC[len(MAC)-1], 16, 16)
	if err != nil {
		log.Fatalf("Failed to parse last number: %v", err)
	}
	lastNum = lastNum + offset
	MAC[len(MAC)-1] = fmt.Sprintf("%02x", lastNum)

	link, err := netlink.LinkByName(devName)
	if err != nil {
		log.Fatalf("Failed to get interface: %v", err)
	}

	hwAddr, err := net.ParseMAC(strings.Join(MAC, ":"))
	if err != nil {
		log.Fatalf("Failed to parse MAC address: %v", err)
	}

	err = netlink.LinkSetHardwareAddr(link, hwAddr)
	if err != nil {
		log.Fatalf("Failed to set MAC address: %v", err)
	}

	fmt.Printf("MAC address of %s changed to %s\n", devName, MAC)
	os.Exit(0)
}
