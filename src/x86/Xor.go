package x86

import (
	"git.urbach.dev/cli/q/src/cpu"
)

// XorRegisterNumber performs a bitwise XOR using a register and a number.
func XorRegisterNumber(code []byte, register cpu.Register, number int) []byte {
	return encodeNum(code, AddressDirect, 0b110, register, number, 0x83, 0x81)
}

// XorRegisterRegister performs a bitwise XOR using two registers.
func XorRegisterRegister(code []byte, register cpu.Register, operand cpu.Register) []byte {
	return encode(code, AddressDirect, operand, register, 8, 0x31)
}