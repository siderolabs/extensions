// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"strings"
	"testing"
)

func TestParse_BasicMeshNode(t *testing.T) {
	in := `
[Interface]
PrivateKey = abc123==
Address    = 192.0.2.1/24
ListenPort = 51820
Jc = 4
Jmin = 40
Jmax = 70
S1 = 50
S2 = 100
H1 = 1027374982
H2 = 2090294422
H3 = 2059343813
H4 = 2072281631

[Peer]
PublicKey  = peer-pub==
Endpoint   = 192.0.2.4:51820
AllowedIPs = 192.0.2.2/32, 10.244.1.0/24
PersistentKeepalive = 25
`
	cfg, err := Parse(strings.NewReader(in))
	if err != nil {
		t.Fatal(err)
	}

	if len(cfg.Address) != 1 || cfg.Address[0].String() != "192.0.2.1/24" {
		t.Errorf("Address: got %v want [192.0.2.1/24]", cfg.Address)
	}
	if cfg.MTU != 1420 {
		t.Errorf("MTU default: got %d want 1420", cfg.MTU)
	}
	if cfg.InterfaceWG["PrivateKey"] != "abc123==" {
		t.Errorf("PrivateKey: got %q", cfg.InterfaceWG["PrivateKey"])
	}
	if cfg.InterfaceWG["Jc"] != "4" {
		t.Errorf("Jc obfuscation param: got %q", cfg.InterfaceWG["Jc"])
	}
	if len(cfg.Peers) != 1 {
		t.Fatalf("Peers: got %d want 1", len(cfg.Peers))
	}
	p := cfg.Peers[0]
	if p.PeerWG["Endpoint"] != "192.0.2.4:51820" {
		t.Errorf("Peer Endpoint: got %q", p.PeerWG["Endpoint"])
	}
	if len(p.AllowedIPs) != 2 {
		t.Fatalf("AllowedIPs: got %d want 2", len(p.AllowedIPs))
	}
	if p.AllowedIPs[1].String() != "10.244.1.0/24" {
		t.Errorf("AllowedIPs[1]: got %s", p.AllowedIPs[1])
	}
}

func TestParse_RejectsWGFormatPollution(t *testing.T) {
	// wg-quick fields are silently dropped (recognised, intentionally ignored),
	// truly unknown ones are an error.
	in := `
[Interface]
PrivateKey = k==
Address    = 10.0.0.1/24
PostUp     = echo hello
GarbageKey = nope
`
	_, err := Parse(strings.NewReader(in))
	if err == nil {
		t.Fatal("expected error on unknown key")
	}
	if !strings.Contains(err.Error(), "GarbageKey") {
		t.Errorf("error should name the unknown key, got: %v", err)
	}

	// Same input without GarbageKey — PostUp is silently ignored.
	in2 := strings.Replace(in, "GarbageKey = nope\n", "", 1)
	cfg, err := Parse(strings.NewReader(in2))
	if err != nil {
		t.Fatalf("unexpected error after removing GarbageKey: %v", err)
	}
	if _, ok := cfg.InterfaceWG["PostUp"]; ok {
		t.Error("PostUp should have been silently dropped, not forwarded to awg setconf")
	}
}

func TestParse_MissingRequired(t *testing.T) {
	cases := []struct {
		name, in, wantErr string
	}{
		{
			name: "no address",
			in: `
[Interface]
PrivateKey = k==
`,
			wantErr: "Address",
		},
		{
			name: "no private key",
			in: `
[Interface]
Address = 10.0.0.1/24
`,
			wantErr: "PrivateKey",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := Parse(strings.NewReader(c.in))
			if err == nil {
				t.Fatal("expected error")
			}
			if !strings.Contains(err.Error(), c.wantErr) {
				t.Errorf("error %q should mention %q", err, c.wantErr)
			}
		})
	}
}

func TestParse_MultipleAddresses(t *testing.T) {
	in := `
[Interface]
PrivateKey = k==
Address = 10.0.0.1/24, fd00::1/64
`
	cfg, err := Parse(strings.NewReader(in))
	if err != nil {
		t.Fatal(err)
	}
	if len(cfg.Address) != 2 {
		t.Fatalf("got %d addresses, want 2", len(cfg.Address))
	}
	if !cfg.Address[1].Addr().Is6() {
		t.Errorf("second address should be IPv6, got %v", cfg.Address[1])
	}
}

func TestParse_IPv6Peer(t *testing.T) {
	in := `
[Interface]
PrivateKey = k==
Address = fd00::1/64

[Peer]
PublicKey = p==
AllowedIPs = fd00::2/128, 2001:db8::/32
`
	cfg, err := Parse(strings.NewReader(in))
	if err != nil {
		t.Fatal(err)
	}
	if len(cfg.Peers[0].AllowedIPs) != 2 {
		t.Fatalf("got %d AllowedIPs, want 2", len(cfg.Peers[0].AllowedIPs))
	}
	if !cfg.Peers[0].AllowedIPs[0].Addr().Is6() {
		t.Errorf("AllowedIPs[0] should be IPv6")
	}
	if cfg.Peers[0].AllowedIPs[1].Bits() != 32 {
		t.Errorf("AllowedIPs[1] prefix length got %d, want 32", cfg.Peers[0].AllowedIPs[1].Bits())
	}
}

func TestParse_PeerBeforeSection(t *testing.T) {
	in := `
PublicKey = nope==
`
	_, err := Parse(strings.NewReader(in))
	if err == nil {
		t.Fatal("expected error for key without preceding [Section] header")
	}
}

func TestWGFormat_OmitsWGQuickFields(t *testing.T) {
	in := `
[Interface]
PrivateKey = k==
Address    = 10.0.0.1/24
MTU        = 1280
ListenPort = 51820
Jc         = 4

[Peer]
PublicKey  = p==
Endpoint   = 1.1.1.1:51820
AllowedIPs = 10.0.0.2/32
`
	cfg, err := Parse(strings.NewReader(in))
	if err != nil {
		t.Fatal(err)
	}
	out := cfg.WGFormat()

	for _, bad := range []string{"Address", "MTU", "DNS", "PostUp", "PreUp"} {
		if strings.Contains(out, bad) {
			t.Errorf("WGFormat output should not contain wg-quick field %q, got:\n%s", bad, out)
		}
	}
	for _, want := range []string{"PrivateKey", "ListenPort", "Jc", "PublicKey", "Endpoint", "AllowedIPs"} {
		if !strings.Contains(out, want) {
			t.Errorf("WGFormat output missing %q, got:\n%s", want, out)
		}
	}
}
