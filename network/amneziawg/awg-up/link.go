// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"errors"
	"fmt"
	"net"
	"net/netip"

	"github.com/vishvananda/netlink"
	"golang.org/x/sys/unix"
)

// BringUp creates (if absent) the amneziawg interface, attaches the addresses,
// sets MTU + admin up.
//
// vishvananda/netlink doesn't have a dedicated AmneziaWG link type — same as
// it didn't have one for wireguard until a few years ago — so we use
// GenericLink with the kernel-side type name "amneziawg". The kernel module
// registers under that name (see amneziawg/src/main.c in upstream).
//
// If the interface already exists from a previous run (the kernel's amneziawg
// netdev survives container exit), we accept it idempotently — same address
// + MTU contract, no recreate.
func BringUp(name string, cfg *Config) error {
	link, err := netlink.LinkByName(name)
	if err != nil {
		var nfe netlink.LinkNotFoundError
		if !errors.As(err, &nfe) {
			return fmt.Errorf("lookup link %q: %w", name, err)
		}
		// Doesn't exist — create it.
		link = &netlink.GenericLink{
			LinkAttrs: netlink.LinkAttrs{
				Name: name,
				MTU:  cfg.MTU,
			},
			LinkType: "amneziawg",
		}
		if err := netlink.LinkAdd(link); err != nil {
			return fmt.Errorf("add link %q (type amneziawg — is the kernel module loaded?): %w", name, err)
		}
		// Re-fetch to pick up the kernel-assigned attrs (index, etc.)
		link, err = netlink.LinkByName(name)
		if err != nil {
			return fmt.Errorf("re-fetch link %q: %w", name, err)
		}
	}

	if link.Attrs().MTU != cfg.MTU {
		if err := netlink.LinkSetMTU(link, cfg.MTU); err != nil {
			return fmt.Errorf("set mtu %d on %q: %w", cfg.MTU, name, err)
		}
	}

	if err := syncAddresses(link, cfg.Address); err != nil {
		return fmt.Errorf("sync addresses on %q: %w", name, err)
	}

	if err := netlink.LinkSetUp(link); err != nil {
		return fmt.Errorf("set %q up: %w", name, err)
	}
	return nil
}

func syncAddresses(link netlink.Link, want []netip.Prefix) error {
	have, err := netlink.AddrList(link, unix.AF_UNSPEC)
	if err != nil {
		return err
	}

	wantSet := map[string]netip.Prefix{}
	for _, p := range want {
		wantSet[p.String()] = p
	}

	// Remove addresses we don't want anymore.
	for _, a := range have {
		key := netipPrefixFromNetlink(a).String()
		if _, keep := wantSet[key]; !keep {
			if err := netlink.AddrDel(link, &a); err != nil {
				return fmt.Errorf("del addr %s: %w", key, err)
			}
			continue
		}
		delete(wantSet, key)
	}

	// Add new ones.
	for s, p := range wantSet {
		addr := &netlink.Addr{IPNet: &net.IPNet{
			IP:   p.Addr().AsSlice(),
			Mask: net.CIDRMask(p.Bits(), p.Addr().BitLen()),
		}}
		if err := netlink.AddrAdd(link, addr); err != nil {
			return fmt.Errorf("add addr %s: %w", s, err)
		}
	}
	return nil
}

func netipPrefixFromNetlink(a netlink.Addr) netip.Prefix {
	ip, _ := netip.AddrFromSlice(a.IP)
	bits, _ := a.Mask.Size()
	return netip.PrefixFrom(ip.Unmap(), bits)
}
