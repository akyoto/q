package arm

import "git.urbach.dev/cli/q/src/cpu"

// OrRegisterNumber performs a bitwise OR using a register and a number.
func OrRegisterNumber(destination cpu.Register, source cpu.Register, number int) (uint32, bool) {
	n, immr, imms, encodable := encodeLogicalImmediate(uint(number))
	return 0b101100100<<23 | reg2BitmaskImm(destination, source, n, immr, imms), encodable
}

// OrRegisterRegister performs a bitwise OR using two registers.
func OrRegisterRegister(destination cpu.Register, source cpu.Register, operand cpu.Register) uint32 {
	return 0b10101010<<24 | reg3(destination, source, operand)
}