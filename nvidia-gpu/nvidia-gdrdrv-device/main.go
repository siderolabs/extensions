// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"syscall" // Primarily for Mknod and Stat_t
	"time"

	// Import the unix package from x/sys for Mkdev, Major, Minor
	// You'll need to run 'go get golang.org/x/sys/unix'
	"golang.org/x/sys/unix"
)

const (
	deviceName      = "gdrdrv" // The driver name we're looking for
	devicePath      = "/dev/" + deviceName
	procDevicesPath = "/proc/devices" // File containing major numbers
	targetPerms     = 0o666           // Desired file permissions (rw-rw-rw-)
	minorNumber     = 0               // Standard minor number for gdrdrv
	retryDelay      = 1 * time.Second // How long to wait between checks
)

func findDeviceMajor() (int, error) {
	file, err := os.Open(procDevicesPath)
	if err != nil {
		return -1, fmt.Errorf("error opening %s: %w", procDevicesPath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		fields := strings.Fields(line)

		if len(fields) == 2 && fields[1] == deviceName {
			major, err := strconv.Atoi(fields[0])
			if err != nil {
				return -1, fmt.Errorf("error parsing major number '%s' for %s: %w", fields[0], deviceName, err)
			}
			return major, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return -1, fmt.Errorf("error reading %s: %w", procDevicesPath, err)
	}

	return -1, fmt.Errorf("device '%s' not found in %s", deviceName, procDevicesPath)
}

func main() {
	log.SetPrefix("gdrdrv-mknod: ")
	log.Printf("waiting for %s driver to register in %s...", deviceName, procDevicesPath)

	var major int
	var err error

	// Loop indefinitely until the major number is found
	for {
		major, err = findDeviceMajor()
		if err == nil {
			log.Printf("Found %s major number: %d", deviceName, major)
			break
		}
		// Log the error and wait before retrying
		log.Printf("waiting for device... (%v)", err)
		time.Sleep(retryDelay)
	}

	// Attempt to remove the device node if it already exists
	log.Printf("checking for existing device node at %s...", devicePath)
	if _, err := os.Stat(devicePath); err == nil {
		log.Printf("removing existing %s...", devicePath)
		if err := os.Remove(devicePath); err != nil {
			log.Fatalf("failed to remove existing %s: %v", devicePath, err)
		}
	} else if !os.IsNotExist(err) {
		// If stat failed for a reason other than the file not existing, it's a problem
		log.Fatalf("error: checking status of %s: %v", devicePath, err)
	}

	log.Printf("creating character device %s (Major: %d, Minor: %d)...", devicePath, major, minorNumber)

	mode := uint32(syscall.S_IFCHR | targetPerms)

	// Combine major and minor numbers into the format needed by Mknod
	// We use unix.Mkdev from the x/sys/unix package
	dev := unix.Mkdev(uint32(major), uint32(minorNumber))

	// Execute the Mknod syscall
	if err := syscall.Mknod(devicePath, mode, int(dev)); err != nil {
		log.Fatalf("Failed to create device node with mknod: %v", err)
	}

	// Mknod *should* set permissions, but explicitly call Chmod for robustness
	log.Printf("Setting permissions on %s to %#o...", devicePath, targetPerms)
	if err := os.Chmod(devicePath, targetPerms); err != nil {
		log.Printf("Warning: Failed to set permissions with chmod (mknod might have already set them): %v", err)
	}

	// Verify creation and permissions by stating the file (like ls -l does)
	fileInfo, err := os.Stat(devicePath)
	if err != nil {
		log.Fatalf("Failed to stat the created device %s: %v", devicePath, err)
	}

	// Extract major/minor from stat info for confirmation log
	devDetails := ""
	if stat_t, ok := fileInfo.Sys().(*syscall.Stat_t); ok {
		maj := unix.Major(uint64(stat_t.Rdev))
		min := unix.Minor(uint64(stat_t.Rdev))
		devDetails = fmt.Sprintf(" (Major: %d, Minor: %d)", maj, min)
	}

	log.Printf("Successfully created %s: Mode=%s%s",
		devicePath,
		fileInfo.Mode().String(),
		devDetails,
	)
}
