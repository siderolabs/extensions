// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Command awg-up brings an AmneziaWG interface up from a wg-style .conf file.
//
// It is the minimal subset of upstream amnezia-vpn/amneziawg-tools' awg-quick
// shell script that's actually needed inside a Talos system-extension service
// container:
//
//  1. ip link add <iface> type amneziawg
//  2. ip address add <Address> dev <iface>
//  3. ip link set mtu <MTU> dev <iface>
//  4. ip link set up dev <iface>
//  5. awg setconf <iface> <stripped-config-on-stdin>
//  6. for each peer's AllowedIPs: ip route add <prefix> dev <iface>
//
// All netlink ops use vishvananda/netlink (same library Kubernetes' net plugins
// use). The wireguard-protocol-with-amneziawg-extensions config is applied via
// the bundled `awg` C binary (~50 KB) over stdin — reimplementing the genl
// protocol in Go would be a maintenance burden that doesn't earn its keep.
//
// No shell, no /proc dependency, no /dev/fd workaround. Exits cleanly after
// the interface is up; the kernel state persists. Designed for use as the
// entrypoint of an extension service container.
package main

import (
	"fmt"
	"log"
	"os"
)

const defaultInterface = "awg0"

func main() {
	log.SetFlags(0)
	log.SetPrefix("awg-up: ")

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: awg-up <path-to-config>\n")
		os.Exit(2)
	}

	cfgPath := os.Args[1]
	iface := os.Getenv("AWG_INTERFACE")
	if iface == "" {
		iface = defaultInterface
	}

	cfg, err := ParseFile(cfgPath)
	if err != nil {
		log.Fatalf("parse %s: %v", cfgPath, err)
	}

	if err := BringUp(iface, cfg); err != nil {
		log.Fatalf("bring up %s: %v", iface, err)
	}

	if err := SetConf(iface, cfg); err != nil {
		log.Fatalf("setconf %s: %v", iface, err)
	}

	if err := InstallRoutes(iface, cfg); err != nil {
		log.Fatalf("install routes: %v", err)
	}

	log.Printf("interface %s up; %d peer(s), %d route(s)", iface, len(cfg.Peers), cfg.RouteCount())
}
