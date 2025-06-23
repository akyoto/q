package arm

import "git.urbach.dev/cli/q/src/cpu"

// CompareRegisterNumber is an alias for a subtraction that updates the conditional flags and discards the result.
func CompareRegisterNumber(register cpu.Register, number int) (code uint32, encodable bool) {
	if number < 0 {
		return addRegisterNumber(ZR, register, -number, 1)
	}

	return subRegisterNumber(ZR, register, number, 1)
}

// CompareRegisterRegister is an alias for a subtraction that updates the conditional flags and discards the result.
func CompareRegisterRegister(reg1 cpu.Register, reg2 cpu.Register) uint32 {
	return subRegisterRegister(ZR, reg1, reg2, 1)
}