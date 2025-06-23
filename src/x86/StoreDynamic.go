package x86

import (
	"encoding/binary"

	"git.urbach.dev/cli/q/src/cpu"
)

// StoreDynamicNumber writes a number to a memory address with a register offset.
func StoreDynamicNumber(code []byte, base cpu.Register, offset cpu.Register, length byte, number int) []byte {
	code = memAccessDynamic(code, 0xC6, 0xC7, 0b000, base, offset, length)

	switch length {
	case 8, 4:
		return binary.LittleEndian.AppendUint32(code, uint32(number))

	case 2:
		return binary.LittleEndian.AppendUint16(code, uint16(number))
	}

	return append(code, byte(number))
}

// StoreDynamicRegister writes the contents of a register to a memory address with a register offset.
func StoreDynamicRegister(code []byte, base cpu.Register, offset cpu.Register, length byte, source cpu.Register) []byte {
	return memAccessDynamic(code, 0x88, 0x89, source, base, offset, length)
}