package x86

import "git.urbach.dev/cli/q/src/cpu"

// memAccess encodes a memory access.
func memAccess(code []byte, opCode8 byte, opCode32 byte, register cpu.Register, base cpu.Register, offset int8, scale ScaleFactor, length byte) []byte {
	opCode := opCode32

	if length == 1 {
		opCode = opCode8
	}

	mod := AddressMemory

	if offset != 0 || base == R5 || base == R13 {
		mod = AddressMemoryOffset8
	}

	if length == 2 {
		code = append(code, 0x66)
	}

	code = encode(code, mod, register, base, length, opCode)

	if base == SP || base == R12 {
		code = append(code, SIB(scale, 0b100, 0b100))
	}

	if mod == AddressMemoryOffset8 {
		code = append(code, byte(offset))
	}

	return code
}