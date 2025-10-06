package arm

import "git.urbach.dev/cli/q/src/cpu"

// Load loads from memory into a register with a register offset.
func Load(destination cpu.Register, base cpu.Register, offset cpu.Register, scale Scale, length byte) uint32 {
	return 1<<22 | StoreRegister(destination, base, offset, scale, length)
}

// LoadFixedOffset loads from memory with a signed offset from -256 to 255 into a register.
func LoadFixedOffset(destination cpu.Register, base cpu.Register, mode AddressMode, offset int, length byte) uint32 {
	return 1<<22 | StoreFixedOffsetRegister(destination, base, mode, offset, length)
}

// LoadFixedOffsetScaled loads from memory with a scaled unsigned offset from 0 to 4095 into a register.
func LoadFixedOffsetScaled(destination cpu.Register, base cpu.Register, mode AddressMode, offset uint, length byte) uint32 {
	return 1<<22 | StoreFixedOffsetRegisterScaled(destination, base, mode, offset, length)
}