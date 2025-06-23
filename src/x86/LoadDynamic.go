package x86

import "git.urbach.dev/cli/q/src/cpu"

// LoadDynamicRegister loads from memory with a register offset into a register.
func LoadDynamicRegister(code []byte, destination cpu.Register, base cpu.Register, offset cpu.Register, length byte) []byte {
	return memAccessDynamic(code, 0x8A, 0x8B, destination, base, offset, length)
}