package x86

import (
	"git.urbach.dev/cli/q/src/cpu"
)

// StoreFixedOffsetNumber writes a number to a memory address.
func StoreFixedOffsetNumber(code []byte, base cpu.Register, offset int8, scale Scale, length byte, number int) []byte {
	code = memAccess(code, 0xC6, 0xC7, 0b000, base, offset, scale, length)
	return appendNumber(code, length, number)
}

// StoreFixedOffsetRegister writes the contents of the register to a memory address.
func StoreFixedOffsetRegister(code []byte, base cpu.Register, offset int8, scale Scale, length byte, register cpu.Register) []byte {
	return memAccess(code, 0x88, 0x89, register, base, offset, scale, length)
}