package x86

import "git.urbach.dev/cli/q/src/cpu"

// MulRegisterNumber multiplies a register with a number.
func MulRegisterNumber(code []byte, register cpu.Register, number int) []byte {
	return encodeNum(code, AddressDirect, register, register, number, 0x6B, 0x69)
}

// MulRegisterRegister multiplies a register with another register.
func MulRegisterRegister(code []byte, register cpu.Register, operand cpu.Register) []byte {
	return encode(code, AddressDirect, register, operand, 8, 0x0F, 0xAF)
}