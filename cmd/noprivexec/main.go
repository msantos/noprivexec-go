package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	"codeberg.org/msantos/noprivexec-go/pkg/noprivexec"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage:", os.Args[0], "<cmd> <...>")
		os.Exit(2)
	}

	if err := noprivexec.Enable(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	arg0, err := exec.LookPath(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(127)
	}

	if err := syscall.Exec(arg0, os.Args[1:], os.Environ()); err != nil {
		fmt.Println(err)
	}
	os.Exit(126)
}
