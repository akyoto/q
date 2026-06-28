package x86

import "git.urbach.dev/cli/q/src/cpu"

// CompareAndSwap compares R0 with the memory contents.
// If they are equal, source register is written to memory.
// If they are not equal, R0 is loaded with the value in memory.
func CompareAndSwap(code []byte, base cpu.Register, offset cpu.Register, scale Scale, length byte, source cpu.Register) []byte {
	return memAccess(code, source, base, offset, scale, length, length, 0x0FB0, 0x0FB1)
}

// CompareAndSwapFixedOffset is a CompareAndSwap with a fixed offset.
func CompareAndSwapFixedOffset(code []byte, base cpu.Register, offset int32, length byte, source cpu.Register) []byte {
	return memAccessFixedOffset(code, source, base, offset, length, 0x0FB0, 0x0FB1)
}