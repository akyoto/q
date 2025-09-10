package x86

import "git.urbach.dev/cli/q/src/cpu"

// LoadRegister loads from memory into a register.
func LoadRegister(code []byte, destination cpu.Register, base cpu.Register, offset int8, scale ScaleFactor, length byte) []byte {
	return memAccess(code, 0x8A, 0x8B, destination, base, offset, scale, length)
}