package x86

import (
	"encoding/binary"

	"git.urbach.dev/cli/q/src/cpu"
)

// memAccessFixedOffset encodes a memory access.
func memAccessFixedOffset(code []byte, register cpu.Register, base cpu.Register, offset int32, length byte, opCode8 uint32, opCode32 uint32) []byte {
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

	code = encode(code, mod, register, base, length, opCode)

	switch mod {
	case AddressMemoryOffset8:
		return append(code, byte(offset))
	case AddressMemoryOffset32:
		return binary.LittleEndian.AppendUint32(code, uint32(offset))
	default:
		return code
	}
}