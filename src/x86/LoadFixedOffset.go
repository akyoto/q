package x86

import "git.urbach.dev/cli/q/src/cpu"

// LoadFixedOffset loads from memory with a signed offset from -128 to 127 into a register.
func LoadFixedOffset(code []byte, destination cpu.Register, base cpu.Register, offset int8, scale Scale, length byte) []byte {
	return memAccess(code, 0x8A, 0x8B, destination, base, offset, scale, length)
}