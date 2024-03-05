package noprivexec

import (
	"golang.org/x/sys/unix"
)

const (
	PLEDGENAMES = "stdio rpath wpath cpath dpath tmppath inet mcast fattr chown flock unix dns getpw sendfd recvfd tape tty proc exec prot_exec settime ps vminfo id pf route wroute audio video bpf unveil error disklabel drm vmm"
)

func Enable() error {
	return unix.Pledge(PLEDGENAMES, PLEDGENAMES)
}
