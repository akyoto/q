package arm

import "git.urbach.dev/cli/q/src/cpu"

// SubRegisterNumber subtracts a number from the given register.
func SubRegisterNumber(destination cpu.Register, source cpu.Register, number int) (code uint32, encodable bool) {
	return subRegisterNumber(destination, source, number, 0)
}

// SubRegisterRegister subtracts a register from a register.
func SubRegisterRegister(destination cpu.Register, source cpu.Register, operand cpu.Register) uint32 {
	return subRegisterRegister(destination, source, operand, 0)
}

// subRegisterNumber subtracts the register and optionally updates the condition flags based on the result.
func subRegisterNumber(destination cpu.Register, source cpu.Register, number int, flags uint32) (code uint32, encodable bool) {
	shift := uint32(0)

	if number > mask12 {
		if number&mask12 != 0 {
			return 0, false
		}

		shift = 1
		number >>= 12

		if number > mask12 {
			return 0, false
		}
	}

	return flags<<29 | 0b110100010<<23 | shift<<22 | reg2Imm(destination, source, number), true
}

// subRegisterRegister subtracts the registers and optionally updates the condition flags based on the result.
func subRegisterRegister(destination cpu.Register, source cpu.Register, operand cpu.Register, flags uint32) uint32 {
	return flags<<29 | 0b11001011000<<21 | reg3(destination, source, operand)
}