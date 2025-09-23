package arm

import "git.urbach.dev/cli/q/src/cpu"

// AddRegisterNumber adds a number to a register.
func AddRegisterNumber(destination cpu.Register, source cpu.Register, number int) (code uint32, encodable bool) {
	if number < 0 {
		return subRegisterNumber(destination, source, uint(-number), 0)
	}

	return addRegisterNumber(destination, source, uint(number), 0)
}

// AddRegisterRegister adds a register to a register.
func AddRegisterRegister(destination cpu.Register, source cpu.Register, operand cpu.Register) uint32 {
	return addRegisterRegister(destination, source, operand, 0)
}

// addRegisterNumber adds the register and optionally updates the condition flags based on the result.
func addRegisterNumber(destination cpu.Register, source cpu.Register, number uint, flags uint32) (code uint32, encodable bool) {
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

	return flags<<29 | 0b100100010<<23 | shift<<22 | reg2Imm(destination, source, number), true
}

// addRegisterRegister adds the registers and optionally updates the condition flags based on the result.
func addRegisterRegister(destination cpu.Register, source cpu.Register, operand cpu.Register, flags uint32) uint32 {
	return flags<<29 | 0b10001011000<<21 | reg3(destination, source, operand)
}