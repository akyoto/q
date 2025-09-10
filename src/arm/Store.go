package arm

import (
	"git.urbach.dev/cli/q/src/cpu"
)

// StoreDynamicRegister writes the contents of the register to a memory address with a dynamic offset.
func StoreDynamicRegister(source cpu.Register, base cpu.Register, offset cpu.Register, scale ScaleFactor, length byte) uint32 {
	return size(length)<<30 | 0b111<<27 | 1<<21 | (LSL << 13) | uint32(scale)<<10 | reg3(source, base, offset)
}

// StoreRegister writes the contents of the register to a memory address.
func StoreRegister(source cpu.Register, base cpu.Register, mode AddressMode, offset int, length byte) uint32 {
	return size(length)<<30 | 0b111<<27 | memory(source, base, mode, offset)
}