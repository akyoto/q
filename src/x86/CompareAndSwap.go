package x86

import "git.urbach.dev/cli/q/src/cpu"

// CompareAndSwap compares R0 with the memory contents.
// If they are equal, source register is written to memory.
// If they are not equal, R0 is loaded with the value in memory.
func CompareAndSwap(code []byte, base cpu.Register, offset cpu.Register, scale Scale, length byte, source cpu.Register) []byte {
	return memAccess(code, source, base, offset, scale, length, length, 0x0FB0, 0x0FB1)
}