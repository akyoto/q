package x86

import "git.urbach.dev/cli/q/src/cpu"

// Compares the register with the number and sets the status flags in the EFLAGS register.
func CompareRegisterNumber(code []byte, register cpu.Register, number int) []byte {
	return encodeNum(code, AddressDirect, 0b111, register, number, 0x83, 0x81)
}

// CompareRegisterRegister compares a register with a register and sets the status flags in the EFLAGS register.
func CompareRegisterRegister(code []byte, registerA cpu.Register, registerB cpu.Register, length byte) []byte {
	opCode := uint32(0x39)

	if length == 1 {
		opCode = 0x38
	}

	return encode(code, AddressDirect, registerB, registerA, length, opCode)
}