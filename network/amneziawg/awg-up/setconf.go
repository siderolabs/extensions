// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// SetConf writes the wg-format part of the config to a tempfile under the
// writable scratch dir and runs `awg setconf <iface> <tempfile>`.
//
// We shell out to the `awg` C binary (~50 KB, bundled in the extension's
// rootfs at /usr/local/bin/awg) because it speaks the netlink-genl protocol
// the AmneziaWG kernel module exposes — including the obfuscation parameter
// fields (Jc, Jmin, Jmax, S1, S2, H1-H4) that aren't part of stock WireGuard
// and aren't supported by golang.zx2c4.com/wireguard/wgctrl.
//
// Reimplementing that protocol in Go would mean duplicating the AmneziaWG
// netlink-attribute table on every upstream release. Not worth the upkeep
// for ~50 KB of bundle size and a single exec call.
//
// We could `cmd.Stdin = strings.NewReader(...)` and pass /dev/stdin to awg,
// but the extension service container doesn't mount /dev so /dev/stdin
// doesn't resolve. Tempfile in the writable config dir is simpler and works
// without extra mounts.
func SetConf(iface string, cfg *Config) error {
	dir := filepath.Dir(awgConfigDir(cfg))
	f, err := os.CreateTemp(dir, "setconf-*.conf")
	if err != nil {
		return fmt.Errorf("create tempfile in %s: %w", dir, err)
	}
	defer os.Remove(f.Name())

	if _, err := f.WriteString(cfg.WGFormat()); err != nil {
		f.Close()
		return fmt.Errorf("write tempfile: %w", err)
	}
	if err := f.Close(); err != nil {
		return fmt.Errorf("close tempfile: %w", err)
	}

	// Absolute path: extension service containers run with an empty or minimal
	// $PATH, and the bundled awg binary is at a known location in our rootfs.
	cmd := exec.Command("/usr/local/bin/awg", "setconf", iface, f.Name())

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("awg setconf %s: %w (stderr: %s)", iface, err, strings.TrimSpace(stderr.String()))
	}
	return nil
}

// awgConfigDir returns a writable directory we can put the tempfile in.
// /var/run/amneziawg is bind-mounted into the extension service container
// at startup (see services/amneziawg.yaml).
func awgConfigDir(*Config) string {
	return "/var/run/amneziawg/"
}
