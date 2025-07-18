//go:build linux

package fs

import (
	"strings"
	"syscall"
	"unsafe"
)

// Walk calls your callback function for every file name inside the directory.
// It doesn't distinguish between files and directories.
func Walk(directory string, callBack func(string)) error {
	fd, err := syscall.Open(directory, 0, 0)

	if err != nil {
		return err
	}

	defer syscall.Close(fd)
	buffer := make([]byte, 512)

	for {
		n, err := syscall.ReadDirent(fd, buffer)

		if err != nil {
			return err
		}

		if n <= 0 {
			break
		}

		readBuffer := buffer[:n]

		for len(readBuffer) > 0 {
			dirent := (*syscall.Dirent)(unsafe.Pointer(&readBuffer[0]))
			readBuffer = readBuffer[dirent.Reclen:]

			if dirent.Ino == 0 {
				continue
			}

			if dirent.Name[0] == '.' {
				continue
			}

			bytePointer := (*byte)(unsafe.Pointer(&dirent.Name[0]))
			name := unsafe.String(bytePointer, 256)
			null := strings.IndexByte(name, 0)
			callBack(strings.Clone(name[:null]))
		}
	}

	return nil
}