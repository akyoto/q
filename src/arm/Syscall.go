package arm

// Syscall is the primary way to communicate with the OS kernel.
func Syscall() uint32 {
	return 0xD4000001
}