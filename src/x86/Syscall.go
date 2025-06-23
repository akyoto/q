package x86

// Syscall is the primary way to communicate with the OS kernel.
func Syscall(code []byte) []byte {
	return append(code, 0x0F, 0x05)
}