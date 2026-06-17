package x86

import (
	"encoding/binary"

	"git.urbach.dev/cli/q/src/cpu"
)

// memLoadExtend encodes a memory load with sign or zero extension.
func memLoadExtend(code []byte, destination cpu.Register, base cpu.Register, offset int32, scale Scale, length byte, opCode8 byte, opCode16 byte, opCode32 byte) []byte {
	var (
		opCode byte
		mod    = AddressMemory
	)

	switch length {
	case 1:
		opCode = opCode8
	case 2:
		opCode = opCode16
	case 4:
		opCode = opCode32
	}

	if offset != 0 || base == R5 || base == R13 {
		mod = AddressMemoryOffset8

		if cpu.SizeInt(int64(offset)) > 1 {
			mod = AddressMemoryOffset32
		}
	}

	w := byte(1)
	r, destination := split(destination)
	b, base := split(base)

	code = append(code, REX(w, r, 0, b))

	if length != 4 {
		code = append(code, opCodePrefix0F...)
	}

	code = append(code, opCode)
	code = append(code, ModRM(mod, byte(destination), byte(base)))

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