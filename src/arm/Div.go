package arm

import "git.urbach.dev/cli/q/src/cpu"

// DivSignedRegisterRegister divides source by operand and stores the value in the destination.
func DivSignedRegisterRegister(destination cpu.Register, source cpu.Register, operand cpu.Register) uint32 {
	return 1<<10 | DivUnsignedRegisterRegister(destination, source, operand)
}

// DivUnsignedRegisterRegister divides unsigned source by unsigned operand and stores the value in the destination.
func DivUnsignedRegisterRegister(destination cpu.Register, source cpu.Register, operand cpu.Register) uint32 {
	return 0b10011010110<<21 | 0b00001<<11 | reg3(destination, source, operand)
}