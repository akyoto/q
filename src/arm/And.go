package arm

import (
	"git.urbach.dev/cli/q/src/cpu"
)

// AndRegisterNumber performs a bitwise AND using a register and a number.
func AndRegisterNumber(destination cpu.Register, source cpu.Register, number int) (uint32, bool) {
	n, immr, imms, encodable := encodeLogicalImmediate(uint(number))
	return 0b100100100<<23 | reg2BitmaskImm(destination, source, n, immr, imms), encodable
}

// AndRegisterRegister performs a bitwise AND using two registers.
func AndRegisterRegister(destination cpu.Register, source cpu.Register, operand cpu.Register) uint32 {
	return 0b10001010<<24 | reg3Imm(destination, source, operand, 0)
}