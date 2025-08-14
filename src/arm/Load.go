package arm

import "git.urbach.dev/cli/q/src/cpu"

// LoadDynamicRegister loads from memory into a register with a dynamic offset.
func LoadDynamicRegister(destination cpu.Register, base cpu.Register, mode AddressMode, offset cpu.Register, length byte) uint32 {
	return size(length)<<30 | 0b111<<27 | 0b01<<22 | 1<<21 | (LSL << 13) | uint32(mode)<<10 | reg3(destination, base, offset)
}

// LoadRegister loads from memory into a register.
func LoadRegister(destination cpu.Register, base cpu.Register, mode AddressMode, offset int, length byte) uint32 {
	return size(length)<<30 | 0b111<<27 | 0b01<<22 | memory(destination, base, mode, offset)
}