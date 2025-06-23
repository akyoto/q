package arm

import "git.urbach.dev/cli/q/src/cpu"

// XorRegisterNumber performs a bitwise XOR using a register and a number.
func XorRegisterNumber(destination cpu.Register, source cpu.Register, number int) (uint32, bool) {
	n, immr, imms, encodable := encodeLogicalImmediate(uint(number))
	return 0b110100100<<23 | reg2BitmaskImm(destination, source, n, immr, imms), encodable
}

// XorRegisterRegister performs a bitwise XOR using two registers.
func XorRegisterRegister(destination cpu.Register, source cpu.Register, operand cpu.Register) uint32 {
	return 0b11001010<<24 | reg3(destination, source, operand)
}