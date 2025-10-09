package arm

import "git.urbach.dev/cli/q/src/cpu"

// LoadAddress calculates the address with the PC-relative offset in the range -1048576 to 1048575
// and writes the result to the destination register.
func LoadAddress(destination cpu.Register, offset int) (code uint32, encodable bool) {
	if offset < -1048576 || offset > 1048575 {
		return 0, false
	}

	hi := uint32(offset>>2) & mask19
	lo := uint32(offset) & mask2
	return lo<<29 | 0b10000<<24 | hi<<5 | uint32(destination), true
}