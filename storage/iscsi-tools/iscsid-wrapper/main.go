package main

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"

	"golang.org/x/sys/unix"
)

func main() {
	log.Println("iscsid-wrapper: starting...")
	cmd := exec.Command("/usr/local/sbin/iscsi-iname")

	var cmdOut bytes.Buffer

	cmd.Stdout = &cmdOut

	if _, err := os.Stat("/etc/iscsi/initiatorname.iscsi"); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			log.Println("iscsid-wrapper: /etc/iscsi/initiatorname.iscsi does not exist, creating")
			if err := cmd.Run(); err != nil {
				log.Printf("iscsi-iname: error generating iscsi initiatorname %v\n", err)
			}

			initiatorName := fmt.Sprintf("InitiatorName=%s", cmdOut.String())
			log.Printf("iscsid-wrapper: writing %s to /etc/iscsi/initiatorname.iscsi", initiatorName)

			if err := os.WriteFile("/etc/iscsi/initiatorname.iscsi", []byte(initiatorName), 0o644); err != nil {
				log.Printf("iscsi-iname: error saving iscsi initiatorname %v\n", err)
			}
		}
	}

	log.Println("iscsid-wrapper: completed..., execing into iscsid")
	if err := unix.Exec("/usr/local/sbin/iscsid", []string{"iscsid", "-f"}, os.Environ()); err != nil {
		log.Fatalf("iscsid: error execing /usr/local/sbin/iscsid %v\n", err)
	}
}
