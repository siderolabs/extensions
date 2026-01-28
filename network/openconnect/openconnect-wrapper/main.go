// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sys/unix"
)

const (
	openconnectBinPath = "/usr/local/bin/openconnect"
	vpncScriptPath     = "/etc/vpnc/vpnc-script"
	stateDir           = "/var/lib/openconnect"
	runDir             = "/var/run/openconnect"
	pidFile            = "/var/run/openconnect/openconnect.pid"
	passwordFile       = "/var/lib/openconnect/.password"
)

// Environment variable names
const (
	envServer      = "OPENCONNECT_SERVER"
	envProtocol    = "OPENCONNECT_PROTOCOL"
	envUser        = "OPENCONNECT_USER"
	envPassword    = "OPENCONNECT_PASSWORD"
	envCertificate = "OPENCONNECT_CERTIFICATE"
	envPrivateKey  = "OPENCONNECT_PRIVATE_KEY"
	envServerCert  = "OPENCONNECT_SERVERCERT"
	envExtraArgs   = "OPENCONNECT_EXTRA_ARGS"
)

// Supported protocols
var validProtocols = map[string]bool{
	"anyconnect": true,
	"nc":         true, // Juniper Network Connect
	"gp":         true, // GlobalProtect
	"pulse":      true, // Pulse Secure
	"f5":         true,
	"fortinet":   true,
	"array":      true,
}

func main() {
	log.Printf("openconnect-wrapper: initializing...")

	// Ensure required directories exist
	if err := os.MkdirAll(stateDir, 0o755); err != nil {
		log.Fatalf("failed to create state directory: %v", err)
	}
	if err := os.MkdirAll(runDir, 0o755); err != nil {
		log.Fatalf("failed to create run directory: %v", err)
	}

	// Get required server address
	server := os.Getenv(envServer)
	if server == "" {
		log.Fatalf("%s environment variable is required", envServer)
	}
	log.Printf("server: %s", server)

	// Build command arguments
	args := []string{openconnectBinPath}

	// Add vpnc-script for route/DNS management
	args = append(args, "--script", vpncScriptPath)

	// Add PID file
	args = append(args, "--pid-file", pidFile)

	// Protocol selection
	if protocol := os.Getenv(envProtocol); protocol != "" {
		if !validProtocols[protocol] {
			log.Fatalf("invalid protocol %q; valid protocols: anyconnect, nc, gp, pulse, f5, fortinet, array", protocol)
		}
		args = append(args, "--protocol", protocol)
		log.Printf("protocol: %s", protocol)
	}

	// Username
	if user := os.Getenv(envUser); user != "" {
		args = append(args, "--user", user)
		log.Printf("user: %s", user)
	}

	// Password handling - write to file and use --passwd-on-stdin
	password := os.Getenv(envPassword)
	if password != "" {
		if err := os.WriteFile(passwordFile, []byte(password), 0o600); err != nil {
			log.Fatalf("failed to write password file: %v", err)
		}
		args = append(args, "--passwd-on-stdin")
		log.Printf("password: configured via stdin")
	}

	// Client certificate authentication
	if cert := os.Getenv(envCertificate); cert != "" {
		// Check if it's a path or inline content
		certPath := cert
		if !filepath.IsAbs(cert) && !strings.HasPrefix(cert, "/") {
			// Write inline certificate to file
			certPath = filepath.Join(stateDir, "client.crt")
			if err := os.WriteFile(certPath, []byte(cert), 0o600); err != nil {
				log.Fatalf("failed to write certificate file: %v", err)
			}
		}
		args = append(args, "--certificate", certPath)
		log.Printf("certificate: configured")
	}

	// Private key for certificate auth
	if key := os.Getenv(envPrivateKey); key != "" {
		keyPath := key
		if !filepath.IsAbs(key) && !strings.HasPrefix(key, "/") {
			// Write inline key to file
			keyPath = filepath.Join(stateDir, "client.key")
			if err := os.WriteFile(keyPath, []byte(key), 0o600); err != nil {
				log.Fatalf("failed to write private key file: %v", err)
			}
		}
		args = append(args, "--sslkey", keyPath)
		log.Printf("private key: configured")
	}

	// Server certificate fingerprint (recommended for security)
	if serverCert := os.Getenv(envServerCert); serverCert != "" {
		args = append(args, "--servercert", serverCert)
		log.Printf("server cert fingerprint: configured")
	}

	// Additional arguments
	if extraArgs := os.Getenv(envExtraArgs); extraArgs != "" {
		// Split extra args by space
		for _, arg := range strings.Fields(extraArgs) {
			args = append(args, arg)
		}
		log.Printf("extra args: %s", extraArgs)
	}

	// Add server URL as the last argument
	args = append(args, server)

	log.Printf("starting openconnect with args: %v", args)

	// If password is set, we need to handle stdin
	if password != "" {
		// We can't use unix.Exec with stdin redirection, so we need a different approach
		// Open the password file and redirect stdin
		f, err := os.Open(passwordFile)
		if err != nil {
			log.Fatalf("failed to open password file: %v", err)
		}

		// Duplicate file descriptor to stdin
		if err := unix.Dup2(int(f.Fd()), int(os.Stdin.Fd())); err != nil {
			f.Close()
			log.Fatalf("failed to redirect stdin: %v", err)
		}
		f.Close()
	}

	// Execute openconnect binary
	if err := unix.Exec(openconnectBinPath, args, filterEnv(os.Environ())); err != nil {
		log.Fatalf("error executing openconnect: %v", err)
	}
}

// filterEnv removes sensitive environment variables from being passed to the subprocess
func filterEnv(env []string) []string {
	filtered := make([]string, 0, len(env))
	sensitiveVars := map[string]bool{
		envPassword:   true,
		envPrivateKey: true,
	}

	for _, e := range env {
		parts := strings.SplitN(e, "=", 2)
		if len(parts) == 2 && sensitiveVars[parts[0]] {
			continue
		}
		filtered = append(filtered, e)
	}

	return filtered
}
