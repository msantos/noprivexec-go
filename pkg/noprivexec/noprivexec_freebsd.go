package noprivexec

import (
	"syscall"
	"unsafe"
)

const (
	P_PID = 0

	PROC_NO_NEW_PRIVS_ENABLE = 1
	PROC_NO_NEW_PRIVS_CTL    = 19
)

func Enable() error {
	data := PROC_NO_NEW_PRIVS_ENABLE
	_, _, errno := syscall.Syscall6(
		syscall.SYS_PROCCTL,
		P_PID,
		0,
		PROC_NO_NEW_PRIVS_CTL,
		uintptr(unsafe.Pointer(&data)),
		0,
		0,
	)
	if errno != 0 {
		return errno
	}
	return nil
}
