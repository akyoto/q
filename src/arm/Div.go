package arm

import "git.urbach.dev/cli/q/src/cpu"

// DivSigned divides source by operand and stores the value in the destination.
func DivSigned(destination cpu.Register, source cpu.Register, operand cpu.Register) uint32 {
	return 0b10011010110<<21 | 0b000011<<10 | reg3(destination, source, operand)
}