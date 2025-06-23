package x86

import (
	"git.urbach.dev/cli/q/src/cpu"
)

// AndRegisterNumber performs a bitwise AND using a register and a number.
func AndRegisterNumber(code []byte, register cpu.Register, number int) []byte {
	return encodeNum(code, AddressDirect, 0b100, register, number, 0x83, 0x81)
}

// AndRegisterRegister performs a bitwise AND using two registers.
func AndRegisterRegister(code []byte, register cpu.Register, operand cpu.Register) []byte {
	return encode(code, AddressDirect, operand, register, 8, 0x21)
}