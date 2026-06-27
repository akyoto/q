package x86

import (
	"git.urbach.dev/cli/q/src/cpu"
)

// StoreFixedOffsetNumber writes a number to a memory address.
func StoreFixedOffsetNumber(code []byte, base cpu.Register, offset int32, length byte, number int32) []byte {
	code = memAccessFixedOffset(code, 0b000, base, offset, length, 0xC6, 0xC7)
	return appendNumber(code, length, number)
}

// StoreFixedOffsetRegister writes the contents of the register to a memory address.
func StoreFixedOffsetRegister(code []byte, base cpu.Register, offset int32, length byte, register cpu.Register) []byte {
	return memAccessFixedOffset(code, register, base, offset, length, 0x88, 0x89)
}