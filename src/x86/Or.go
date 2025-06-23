package x86

import (
	"git.urbach.dev/cli/q/src/cpu"
)

// OrRegisterNumber performs a bitwise OR using a register and a number.
func OrRegisterNumber(code []byte, register cpu.Register, number int) []byte {
	return encodeNum(code, AddressDirect, 0b001, register, number, 0x83, 0x81)
}

// OrRegisterRegister performs a bitwise OR using two registers.
func OrRegisterRegister(code []byte, register cpu.Register, operand cpu.Register) []byte {
	return encode(code, AddressDirect, operand, register, 8, 0x09)
}