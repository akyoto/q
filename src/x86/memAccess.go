package x86

import (
	"encoding/binary"

	"git.urbach.dev/cli/q/src/cpu"
)

// memAccess encodes a memory access.
func memAccess(code []byte, register cpu.Register, base cpu.Register, offset int32, scale Scale, length byte, opCode8 byte, opCode32 byte) []byte {
	opCode := opCode32

	if length == 1 {
		opCode = opCode8
	}

	mod := AddressMemory

	if offset != 0 || base == R5 || base == R13 {
		mod = AddressMemoryOffset8

		if cpu.SizeInt(int64(offset)) > 1 {
			mod = AddressMemoryOffset32
		}
	}

	if length == 2 {
		code = append(code, 0x66)
	}

	code = encode(code, mod, register, base, length, opCode)

	if base == SP || base == R12 {
		code = append(code, SIB(scale, 0b100, 0b100))
	}

	switch mod {
	case AddressMemoryOffset8:
		return append(code, byte(offset))
	case AddressMemoryOffset32:
		return binary.LittleEndian.AppendUint32(code, uint32(offset))
	default:
		return code
	}
}