package x86

import "git.urbach.dev/cli/q/src/cpu"

// LoadFixedOffset loads from memory with a signed offset from -128 to 127 into a register.
func LoadFixedOffset(code []byte, destination cpu.Register, base cpu.Register, offset int32, scale Scale, length byte) []byte {
	return memAccess(code, destination, base, offset, scale, length, 0x8A, 0x8B)
}

// LoadFixedOffsetSignExtend loads from memory with a signed offset from -128 to 127 into a register and sign-extends it.
func LoadFixedOffsetSignExtend(code []byte, destination cpu.Register, base cpu.Register, offset int32, scale Scale, length byte) []byte {
	return memLoadExtend(code, destination, base, offset, scale, length, 0xBE, 0xBF, 0x63)
}

// LoadFixedOffsetZeroExtend loads from memory with a signed offset from -128 to 127 into a register and zero-extends it.
func LoadFixedOffsetZeroExtend(code []byte, destination cpu.Register, base cpu.Register, offset int32, scale Scale, length byte) []byte {
	return memLoadExtend(code, destination, base, offset, scale, length, 0xB6, 0xB7, 0)
}