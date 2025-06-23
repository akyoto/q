package arm

import "git.urbach.dev/cli/q/src/cpu"

// LoadAddress calculates the address with the PC-relative offset and writes the result to the destination register.
func LoadAddress(destination cpu.Register, offset int) uint32 {
	hi := uint32(offset) >> 2
	lo := uint32(offset) & 0b11
	return lo<<29 | 0b10000<<24 | hi<<5 | uint32(destination)
}