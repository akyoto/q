//go:build linux

package memfile

import (
	"os"

	"golang.org/x/sys/unix"
)

// New creates a new anonymous in-memory file.
func New(name string) (*os.File, error) {
	fd, err := unix.MemfdCreate(name, unix.MFD_CLOEXEC)

	if err != nil {
		return nil, err
	}

	memFile := os.NewFile(uintptr(fd), name)
	return memFile, nil
}