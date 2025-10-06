package arm

import (
	"git.urbach.dev/cli/q/src/cpu"
)

// StoreRegister writes the contents of the register to a memory address with a register offset.
func StoreRegister(source cpu.Register, base cpu.Register, offset cpu.Register, scale Scale, length byte) uint32 {
	return size(length)<<30 | 0b111<<27 | 1<<21 | (LSL << 13) | uint32(scale&1)<<12 | uint32(RegisterOffset)<<10 | reg3(source, base, offset)
}

// StoreFixedOffsetRegister writes the contents of the register to a memory address with a signed offset from -256 to 255.
func StoreFixedOffsetRegister(source cpu.Register, base cpu.Register, mode AddressMode, offset int, length byte) uint32 {
	return size(length)<<30 | 0b111<<27 | memory(source, base, mode, offset)
}