package main

import (
	"fmt"
	"strings"
	"testing"
)

func mockReadFile(data string) func(string) ([]byte, error) {
	return func(filename string) ([]byte, error) {
		if filename == "/proc/cmdline" {
			return []byte(data), nil
		}
		return nil, fmt.Errorf("unexpected file read: %s", filename)
	}
}

func TestParseKernelCmdline_ValidInput(t *testing.T) {
	readFile := mockReadFile("pcirebind.rebind=0000:04:00.0_ixgbe+vfio-pci pcirebind.rebind=0000:04:01.0_ixgbe+vfio-pci")

	expected := []rebind{
		{id: "0000:04:00.0", oldDriver: "ixgbe", newDriver: "vfio-pci"},
		{id: "0000:04:01.0", oldDriver: "ixgbe", newDriver: "vfio-pci"},
	}

	result, err := parseKernelCmdline(readFile)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if len(result) != len(expected) {
		t.Fatalf("Expected %d rebinds, got %d", len(expected), len(result))
	}

	for i, r := range result {
		if r != expected[i] {
			t.Errorf("Expected rebind %+v, got %+v", expected[i], r)
		}
	}
}

func TestParseKernelCmdline_InvalidInput_NoDriverSeparator(t *testing.T) {
	readFile := mockReadFile("pcirebind.rebind=0000:04:00.0_ixgbevfio-pci")

	_, err := parseKernelCmdline(readFile)
	if err == nil || !strings.Contains(err.Error(), "no rebinds found in kernel command line") {
		t.Fatalf("Expected error for invalid input, got: %v", err)
	}
}

func TestParseKernelCmdline_InvalidInput_NoUnderscoreSeparator(t *testing.T) {
	readFile := mockReadFile("pcirebind.rebind=0000:04:00.0+vfio-pci")

	_, err := parseKernelCmdline(readFile)
	if err == nil || !strings.Contains(err.Error(), "no rebinds found in kernel command line") {
		t.Fatalf("Expected error for invalid input, got: %v", err)
	}
}

func TestParseKernelCmdline_InvalidInput_EmptyInput(t *testing.T) {
	readFile := mockReadFile("")

	_, err := parseKernelCmdline(readFile)
	if err == nil || !strings.Contains(err.Error(), "no rebinds found in kernel command line") {
		t.Fatalf("Expected error for empty input, got: %v", err)
	}
}

func TestParseKernelCmdline_ValidAndInvalidMixed(t *testing.T) {
	readFile := mockReadFile("pcirebind.rebind=0000:04:00.0_ixgbe+vfio-pci pcirebind.rebind=0000:04:01.0_invalidinput")

	expected := []rebind{
		{id: "0000:04:00.0", oldDriver: "ixgbe", newDriver: "vfio-pci"},
	}

	result, err := parseKernelCmdline(readFile)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if len(result) != len(expected) {
		t.Fatalf("Expected %d valid rebind, got %d", len(expected), len(result))
	}

	for i, r := range result {
		if r != expected[i] {
			t.Errorf("Expected rebind %+v, got %+v", expected[i], r)
		}
	}
}
