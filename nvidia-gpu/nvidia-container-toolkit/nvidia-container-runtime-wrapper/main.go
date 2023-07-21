package main

import (
	"log"
	"os"
	"path/filepath"

	"golang.org/x/sys/unix"
)

func main() {
	cmdName := filepath.Base(os.Args[0])
	switch cmdName {
	case
		"nvidia-container-runtime",
		"nvidia-container-runtime-hook",
		"nvidia-container-runtime.cdi",
		"nvidia-container-runtime.legacy",
		"nvidia-container-toolkit",
		"nvidia-ctk":
		execCommand(cmdName, os.Args[1:])
	default:
		log.Fatalf("nvidia-container-runtime-wrapper: unknown command %s\n", cmdName)
	}
}

func execCommand(cmdName string, args []string) {
	environ := os.Environ()
	environ = append(environ, "XDG_CONFIG_HOME=/usr/local/etc")

	realCmdName := cmdName + ".real"

	cmdArgs := []string{realCmdName}

	if cmdName == "nvidia-container-runtime-hook" {
		cmdArgs = append(
			cmdArgs,
			"-config",
			"/usr/local/etc/nvidia-container-runtime/config.toml",
		)
	}

	cmdArgs = append(cmdArgs, args...)

	cmdFullPath := filepath.Join("/usr/local/bin", realCmdName)

	if err := unix.Exec(cmdFullPath, cmdArgs, environ); err != nil {
		log.Fatalf("nvidia-container-runtime-wrapper: error execing %s %v\n", cmdFullPath, err)
	}
}
