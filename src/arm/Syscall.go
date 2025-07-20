package arm

// Syscall is the primary way to communicate with the OS kernel.
func Syscall(imm16 int) uint32 {
	return 0b11010100000<<21 | uint32(imm16&mask16)<<5 | 0b00001
}