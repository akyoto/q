package x86

import (
	"git.urbach.dev/cli/q/src/cpu"
)

// StoreNumber writes a number to a memory address with a register offset.
func StoreNumber(code []byte, base cpu.Register, offset cpu.Register, scale Scale, length byte, number int) []byte {
	code = memAccessDynamic(code, 0xC6, 0xC7, 0b000, base, offset, scale, length)
	return appendNumber(code, length, number)
}

// StoreRegister writes the contents of a register to a memory address with a register offset.
func StoreRegister(code []byte, base cpu.Register, offset cpu.Register, scale Scale, length byte, source cpu.Register) []byte {
	return memAccessDynamic(code, 0x88, 0x89, source, base, offset, scale, length)
}