package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"go.iscode.ca/noprivexec/pkg/noprivexec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <cmd> <...>", os.Args[0])
		os.Exit(2)
	}

	if err := noprivexec.Enable(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	arg0, err := exec.LookPath(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(127)
	}

	if err := syscall.Exec(arg0, os.Args[1:], os.Environ()); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	os.Exit(126)
}
