//go:build linux

package memfile

import (
	"os"
	"strconv"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

// Exec executes an in-memory file.
func Exec(file *os.File) error {
	empty, err := syscall.BytePtrFromString("")

	if err != nil {
		return err
	}

	argv := []string{"/proc/self/fd/" + strconv.Itoa(int(file.Fd()))}
	argvp, err := syscall.SlicePtrFromStrings(argv)

	if err != nil {
		return err
	}

	envv := os.Environ()
	envvp, err := syscall.SlicePtrFromStrings(envv)

	if err != nil {
		return err
	}

	_, _, errno := syscall.Syscall6(
		unix.SYS_EXECVEAT,
		file.Fd(),
		uintptr(unsafe.Pointer(empty)),
		uintptr(unsafe.Pointer(&argvp[0])),
		uintptr(unsafe.Pointer(&envvp[0])),
		uintptr(unix.AT_EMPTY_PATH),
		0,
	)

	return errno
}