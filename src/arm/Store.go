package arm

import (
	"git.urbach.dev/cli/q/src/cpu"
)

// StoreDynamicRegister writes the contents of the register to a memory address with a dynamic offset.
func StoreDynamicRegister(source cpu.Register, base cpu.Register, mode AddressMode, offset cpu.Register, length byte) uint32 {
	common := 1<<21 | (LSL << 13) | uint32(mode)<<10 | reg3(source, base, offset)

	switch length {
	case 1:
		return 0b00111<<27 | common
	case 2:
		return 0b01111<<27 | common
	case 4:
		return 0b10111<<27 | common
	default:
		return 0b11111<<27 | common
	}
}

// StoreRegister writes the contents of the register to a memory address.
func StoreRegister(source cpu.Register, base cpu.Register, mode AddressMode, offset int, length byte) uint32 {
	common := memory(source, base, mode, offset)

	switch length {
	case 1:
		return 0b00111<<27 | common
	case 2:
		return 0b01111<<27 | common
	case 4:
		return 0b10111<<27 | common
	default:
		return 0b11111<<27 | common
	}
}