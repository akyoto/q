package arm

import "git.urbach.dev/cli/q/src/cpu"

// SignExtend sign-extends the value in the source register to 64 bits and writes it to the destination register.
func SignExtend(destination cpu.Register, source cpu.Register, length byte) uint32 {
	return 0b10010011010<<21 | uint32(length*8-1)<<10 | uint32(source)<<5 | uint32(destination)
}