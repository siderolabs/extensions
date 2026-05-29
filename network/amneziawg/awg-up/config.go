// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"bufio"
	"fmt"
	"io"
	"net/netip"
	"os"
	"strconv"
	"strings"
)

// Config models a parsed AmneziaWG .conf file.
//
// The fields under Interface that wg-format would NOT recognise (Address, MTU)
// are extracted here for the link-setup phase. Everything that DOES go into
// `awg setconf` is preserved in InterfaceWG and PeerWG so we can serialise
// them back out and pipe to the awg binary unchanged.
type Config struct {
	Address []netip.Prefix // [Interface] Address — for ip addr add
	MTU     int            // [Interface] MTU (optional; defaults to 1420)

	InterfaceWG map[string]string // [Interface] keys that awg setconf consumes
	Peers       []Peer
}

type Peer struct {
	PeerWG     map[string]string // keys awg setconf consumes
	AllowedIPs []netip.Prefix    // also routed via `ip route add ... dev awg0`
}

const defaultMTU = 1420

// awgSetconfInterfaceKeys are the [Interface] fields awg setconf understands.
// All other keys (Address, MTU, DNS, PreUp, PostUp, ...) are wg-quick
// concerns and must NOT be forwarded to `awg setconf`.
var awgSetconfInterfaceKeys = map[string]struct{}{
	"PrivateKey": {},
	"ListenPort": {},
	"FwMark":     {},
	// AmneziaWG obfuscation extension fields
	"Jc": {}, "Jmin": {}, "Jmax": {},
	"S1": {}, "S2": {},
	"H1": {}, "H2": {}, "H3": {}, "H4": {},
}

var awgSetconfPeerKeys = map[string]struct{}{
	"PublicKey":           {},
	"PresharedKey":        {},
	"AllowedIPs":          {},
	"Endpoint":            {},
	"PersistentKeepalive": {},
}

func ParseFile(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return Parse(f)
}

func Parse(r io.Reader) (*Config, error) {
	cfg := &Config{
		InterfaceWG: map[string]string{},
		MTU:         defaultMTU,
	}

	var section string
	var curPeer *Peer

	sc := bufio.NewScanner(r)
	lineno := 0
	for sc.Scan() {
		lineno++
		line := strings.TrimSpace(sc.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// [Section] header
		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			section = strings.TrimSpace(line[1 : len(line)-1])
			if section == "Peer" {
				cfg.Peers = append(cfg.Peers, Peer{PeerWG: map[string]string{}})
				curPeer = &cfg.Peers[len(cfg.Peers)-1]
			}
			continue
		}

		// key = value
		i := strings.IndexByte(line, '=')
		if i < 0 {
			return nil, fmt.Errorf("line %d: not a key=value", lineno)
		}
		key := strings.TrimSpace(line[:i])
		val := strings.TrimSpace(line[i+1:])

		switch section {
		case "Interface":
			if err := cfg.applyInterfaceKey(key, val); err != nil {
				return nil, fmt.Errorf("line %d (Interface.%s): %w", lineno, key, err)
			}
		case "Peer":
			if curPeer == nil {
				return nil, fmt.Errorf("line %d: Peer key before [Peer] header", lineno)
			}
			if err := curPeer.applyKey(key, val); err != nil {
				return nil, fmt.Errorf("line %d (Peer.%s): %w", lineno, key, err)
			}
		default:
			return nil, fmt.Errorf("line %d: key %s before any [Section] header", lineno, key)
		}
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}

	if len(cfg.Address) == 0 {
		return nil, fmt.Errorf("config missing [Interface] Address")
	}
	if _, ok := cfg.InterfaceWG["PrivateKey"]; !ok {
		return nil, fmt.Errorf("config missing [Interface] PrivateKey")
	}
	return cfg, nil
}

func (c *Config) applyInterfaceKey(key, val string) error {
	switch key {
	case "Address":
		for _, p := range strings.Split(val, ",") {
			prefix, err := netip.ParsePrefix(strings.TrimSpace(p))
			if err != nil {
				return err
			}
			c.Address = append(c.Address, prefix)
		}
		return nil
	case "MTU":
		n, err := strconv.Atoi(val)
		if err != nil {
			return err
		}
		c.MTU = n
		return nil
	case "DNS", "Table", "PreUp", "PostUp", "PreDown", "PostDown", "SaveConfig":
		// Recognised wg-quick fields we intentionally ignore on Talos —
		// Talos already manages DNS, has no concept of running scripts in
		// response to interface lifecycle, and routes are handled in
		// InstallRoutes below.
		return nil
	default:
		// wg-format key — feed straight to awg setconf
		if _, ok := awgSetconfInterfaceKeys[key]; ok {
			c.InterfaceWG[key] = val
			return nil
		}
		return fmt.Errorf("unknown [Interface] key %q", key)
	}
}

func (p *Peer) applyKey(key, val string) error {
	if key == "AllowedIPs" {
		for _, s := range strings.Split(val, ",") {
			prefix, err := netip.ParsePrefix(strings.TrimSpace(s))
			if err != nil {
				return err
			}
			p.AllowedIPs = append(p.AllowedIPs, prefix)
		}
		// AllowedIPs ALSO goes to awg setconf (peer cryptokey-routing table),
		// so preserve the original string.
		p.PeerWG[key] = val
		return nil
	}
	if _, ok := awgSetconfPeerKeys[key]; ok {
		p.PeerWG[key] = val
		return nil
	}
	return fmt.Errorf("unknown [Peer] key %q", key)
}

// WGFormat serialises the parts of the config that `awg setconf` expects.
// Address/MTU/DNS/Table/PreUp/etc. are wg-quick concerns and are omitted.
func (c *Config) WGFormat() string {
	var sb strings.Builder
	sb.WriteString("[Interface]\n")
	for k, v := range c.InterfaceWG {
		fmt.Fprintf(&sb, "%s = %s\n", k, v)
	}
	for _, peer := range c.Peers {
		sb.WriteString("\n[Peer]\n")
		for k, v := range peer.PeerWG {
			fmt.Fprintf(&sb, "%s = %s\n", k, v)
		}
	}
	return sb.String()
}

func (c *Config) RouteCount() int {
	n := 0
	for _, p := range c.Peers {
		n += len(p.AllowedIPs)
	}
	return n
}
