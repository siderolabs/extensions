// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"errors"
	"fmt"
	"net"

	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"
)

// InstallRoutes adds one route per AllowedIPs prefix, scoped to the
// amneziawg interface. Skips:
//   - prefixes covering the local subnet (the kernel installs that as a
//     "connected" route automatically when the address is added),
//   - prefixes that already exist on the interface (idempotent re-runs).
//
// We deliberately do NOT install a default route — that's an operator
// concern (e.g. when the mesh is the *only* connection path) and belongs
// in machine config, not here.
func InstallRoutes(iface string, cfg *Config) error {
	link, err := netlink.LinkByName(iface)
	if err != nil {
		return err
	}

	existing, err := netlink.RouteList(link, unix.AF_UNSPEC)
	if err != nil {
		return err
	}
	have := map[string]struct{}{}
	for _, r := range existing {
		if r.Dst == nil {
			continue
		}
		have[r.Dst.String()] = struct{}{}
	}

	for _, peer := range cfg.Peers {
		for _, p := range peer.AllowedIPs {
			if p.Bits() == 0 {
				// 0.0.0.0/0 or ::/0 — "full tunnel" config. awg-quick handles
				// this with fwmark + policy-based routing tables so the WG
				// tunnel's own UDP packets don't loop back through itself.
				// We don't implement that machinery (out of scope for the
				// Talos cluster-mesh use case); installing a bare /0 route
				// here would clobber the operator's default route.
				//
				// If you actually want full-tunnel routing on a Talos node,
				// configure it via machine.network.routes or the host's
				// route table, not via AllowedIPs.
				return fmt.Errorf("AllowedIPs = %s is a default route — full-tunnel routing isn't supported by awg-up; use specific subnets or configure routing via machine.network", p)
			}
			dst := &net.IPNet{
				IP:   p.Masked().Addr().AsSlice(),
				Mask: net.CIDRMask(p.Bits(), p.Addr().BitLen()),
			}
			if _, dup := have[dst.String()]; dup {
				continue
			}
			route := &netlink.Route{
				LinkIndex: link.Attrs().Index,
				Dst:       dst,
				Scope:     netlink.SCOPE_LINK,
			}
			if err := netlink.RouteAdd(route); err != nil {
				// EEXIST shouldn't happen given our dedup above, but be lenient
				// to survive races / leftover state.
				if errors.Is(err, unix.EEXIST) {
					continue
				}
				return fmt.Errorf("add route %s dev %s: %w", dst, iface, err)
			}
		}
	}
	return nil
}
