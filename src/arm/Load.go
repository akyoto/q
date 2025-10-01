package arm

import "git.urbach.dev/cli/q/src/cpu"

// LoadDynamicRegister loads from memory into a register with a register offset.
func LoadDynamicRegister(destination cpu.Register, base cpu.Register, offset cpu.Register, scale Scale, length byte) uint32 {
	return 1<<22 | StoreDynamicRegister(destination, base, offset, scale, length)
}

// LoadRegister loads from memory with a signed offset from -256 to 255 into a register.
func LoadRegister(destination cpu.Register, base cpu.Register, mode AddressMode, offset int, length byte) uint32 {
	return 1<<22 | StoreRegister(destination, base, mode, offset, length)
}