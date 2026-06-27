package x86

import (
	"git.urbach.dev/cli/q/src/cpu"
)

// memAccess encodes a memory access using the value of a register as an offset.
func memAccess(code []byte, register cpu.Register, base cpu.Register, offset cpu.Register, scale Scale, length byte, operandLength byte, opCode8 uint32, opCode32 uint32) []byte {
	opCode := opCode32

	if length == 1 {
		opCode = opCode8
	}

	if offset == SP {
		offset, base = base, offset
	}

	mod := AddressMemory

	if base == R5 || base == R13 {
		mod = AddressMemoryOffset8
	}

	return encode2(code, mod, register, base, offset, scale, operandLength, opCode)
}