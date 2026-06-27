package x86

import "git.urbach.dev/cli/q/src/cpu"

// Load loads from memory with a register offset into a register.
func Load(code []byte, destination cpu.Register, base cpu.Register, offset cpu.Register, scale Scale, length byte) []byte {
	return memAccess(code, destination, base, offset, scale, length, length, 0x8A, 0x8B)
}

// LoadSignExtend loads from memory with a register offset into a register and sign-extends it.
func LoadSignExtend(code []byte, destination cpu.Register, base cpu.Register, offset cpu.Register, scale Scale, length byte) []byte {
	if length == 4 {
		return memAccess(code, destination, base, offset, scale, length, 8, 0x63, 0x63)
	}

	return memAccess(code, destination, base, offset, scale, length, 8, 0x0FBE, 0x0FBF)
}

// LoadZeroExtend loads from memory with a register offset into a register and zero-extends it.
func LoadZeroExtend(code []byte, destination cpu.Register, base cpu.Register, offset cpu.Register, scale Scale, length byte) []byte {
	return memAccess(code, destination, base, offset, scale, length, 8, 0x0FB6, 0x0FB7)
}