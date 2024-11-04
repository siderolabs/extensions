package main

import (
    "fmt"
    "os"
    "os/exec"
    "syscall"
)

func main() {
    fmt.Println("Starting multipath daemon...")

    //setup a chroot environment for the extension
    newRoot := "/usr/local"
    if err := syscall.Chroot(newRoot); err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
    
    //start the multipath daemon in the chroot environment
    cmd := exec.Command("/usr/sbin/multipathd", "-d")
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    if err := cmd.Run(); err != nil {
        fmt.Println("Command execution failed:", err)
    }
}
