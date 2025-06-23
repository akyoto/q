package x86

import (
	"git.urbach.dev/cli/q/src/cpu"
)

// SubRegisterNumber subtracts a number from the given register.
func SubRegisterNumber(code []byte, register cpu.Register, number int) []byte {
	return encodeNum(code, AddressDirect, 0b101, register, number, 0x83, 0x81)
}

// SubRegisterRegister subtracts a register value from another register.
func SubRegisterRegister(code []byte, register cpu.Register, operand cpu.Register) []byte {
	return encode(code, AddressDirect, operand, register, 8, 0x29)
}