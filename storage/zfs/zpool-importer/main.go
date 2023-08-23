package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/thediveo/gons/reexec"
)

func init() {
	reexec.Register("zpool-importer", func() {
		cmd := exec.Command("/usr/local/sbin/zpool", "import", "-fa")

		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintf(os.Stderr, "err: %v, %s", err, out)
		}

		fmt.Fprintf(os.Stdout, `"zpool-importer: \"zpool import -fa\" ran successfully; exiting...\n"`)
	})
}

func main() {
	reexec.CheckAction()

	var result string

	if err := reexec.RunReexecAction(
		"zpool-importer",
		reexec.Result(&result),
		reexec.Namespaces([]reexec.Namespace{
			{
				Type: "mnt",
				Path: "/proc/1/ns/mnt",
			},
		}),
	); err != nil {
		log.Fatalf("zpool-importer: reexec failed: %v", err)
	}

	fmt.Printf("zpool-importer: reexec succeeded: %s\n", result)
}
