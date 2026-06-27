package x86

import "git.urbach.dev/cli/q/src/cpu"

// LoadFixedOffset loads from memory with a signed offset from -128 to 127 into a register.
func LoadFixedOffset(code []byte, destination cpu.Register, base cpu.Register, offset int32, length byte) []byte {
	return memAccessFixedOffset(code, destination, base, offset, length, 0x8A, 0x8B)
}

// LoadFixedOffsetSignExtend loads from memory with a signed offset from -128 to 127 into a register and sign-extends it.
func LoadFixedOffsetSignExtend(code []byte, destination cpu.Register, base cpu.Register, offset int32, length byte) []byte {
	return memLoadExtend(code, destination, base, offset, length, 0x0FBE, 0x0FBF, 0x63)
}

// LoadFixedOffsetZeroExtend loads from memory with a signed offset from -128 to 127 into a register and zero-extends it.
func LoadFixedOffsetZeroExtend(code []byte, destination cpu.Register, base cpu.Register, offset int32, length byte) []byte {
	return memLoadExtend(code, destination, base, offset, length, 0x0FB6, 0x0FB7, 0)
}