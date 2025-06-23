package x86

import (
	"git.urbach.dev/cli/q/src/cpu"
)

// AddRegisterNumber adds a number to the given register.
func AddRegisterNumber(code []byte, register cpu.Register, number int) []byte {
	return encodeNum(code, AddressDirect, 0b000, register, number, 0x83, 0x81)
}

// AddRegisterRegister adds a register value into another register.
func AddRegisterRegister(code []byte, register cpu.Register, operand cpu.Register) []byte {
	return encode(code, AddressDirect, operand, register, 8, 0x01)
}