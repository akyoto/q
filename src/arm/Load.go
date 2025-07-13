package arm

import "git.urbach.dev/cli/q/src/cpu"

// LoadRegister loads from memory into a register.
func LoadRegister(destination cpu.Register, base cpu.Register, mode AddressMode, offset int, length byte) uint32 {
	common := 1<<22 | memory(destination, base, mode, offset)

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