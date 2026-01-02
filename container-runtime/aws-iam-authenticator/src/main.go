// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	log.Printf("starting the aws-iam-authenticator init process")
	defer log.Printf("finishing the aws-iam-authenticator init process")

	// Create the working directory
	workDir := "/usr/local/lib/aws-iam-authenticator"
	if err := os.MkdirAll(workDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating directory %s: %v\n", workDir, err)
		os.Exit(1)
	}

	// Change to the working directory
	if err := os.Chdir(workDir); err != nil {
		fmt.Fprintf(os.Stderr, "Error changing to directory %s: %v\n", workDir, err)
		os.Exit(1)
	}

	args := []string{"init"}

	if clusterID := os.Getenv("CLUSTER_ID"); clusterID != "" {
		args = append(args, fmt.Sprintf("-i=%s", clusterID))
	}

	if hostname := os.Getenv("HOSTNAME"); hostname != "" {
		args = append(args, fmt.Sprintf("--hostname=%s", hostname))
	}

	if address := os.Getenv("ADDRESS"); address != "" {
		args = append(args, fmt.Sprintf("--address=%s", address))
	}

	if config := os.Getenv("CONFIG_FILE"); config != "" {
		args = append(args, fmt.Sprintf("--config=%s", config))
	}

	if logFormat := os.Getenv("LOG_FORMAT"); logFormat != "" {
		args = append(args, fmt.Sprintf("--log-format=%s", logFormat))
	}

	log.Printf("executing: /usr/local/bin/aws-iam-authenticator %v", args)

	cmd := exec.Command("/usr/local/bin/aws-iam-authenticator", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
