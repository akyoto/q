package x86

import "git.urbach.dev/cli/q/src/cpu"

// split splits a 4 bit register index into the upper bit and the lower 3 bits.
func split(reg cpu.Register) (upper byte, lower cpu.Register) {
	if reg > 0b111 {
		return 1, reg & 0b111
	}

	return 0, reg
}