package arm

import "git.urbach.dev/cli/q/src/cpu"

// MulRegisterRegister multiplies `multiplicand` with `multiplier` and saves the result in `destination`
func MulRegisterRegister(destination cpu.Register, multiplicand cpu.Register, multiplier cpu.Register) uint32 {
	return 0b10011011000<<21 | reg4(destination, multiplicand, multiplier, ZR)
}

// MultiplySubtract multiplies `multiplicand` with `multiplier`, subtracts the product from `minuend` and saves the result in `destination`.
func MultiplySubtract(destination cpu.Register, multiplicand cpu.Register, multiplier cpu.Register, minuend cpu.Register) uint32 {
	return 0b10011011000<<21 | 1<<15 | reg4(destination, multiplicand, multiplier, minuend)
}