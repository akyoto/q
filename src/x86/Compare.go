package x86

import "git.urbach.dev/cli/q/src/cpu"

// Compares the register with the number and sets the status flags in the EFLAGS register.
func CompareRegisterNumber(code []byte, register cpu.Register, number int) []byte {
	return encodeNum(code, AddressDirect, 0b111, register, number, 0x83, 0x81)
}

// CompareRegisterRegister compares a register with a register and sets the status flags in the EFLAGS register.
func CompareRegisterRegister(code []byte, registerA cpu.Register, registerB cpu.Register) []byte {
	return encode(code, AddressDirect, registerB, registerA, 8, 0x39)
}