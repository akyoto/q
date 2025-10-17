package x86

import "git.urbach.dev/cli/q/src/cpu"

// memAccessDynamic encodes a memory access using the value of a register as an offset.
func memAccessDynamic(code []byte, opCode8 byte, opCode32 byte, register cpu.Register, base cpu.Register, offset cpu.Register, scale Scale, length byte) []byte {
	var (
		opCode = opCode32
		mod    = AddressMemory
	)

	if length == 1 {
		opCode = opCode8
	}

	if offset == SP {
		offset, base = base, offset
	}

	w := byte(0)

	if length == 8 {
		w = 1
	}

	r, register := split(register)
	x, offset := split(offset)
	b, base := split(base)

	if base == R5 || base == R13 {
		mod = AddressMemoryOffset8
	}

	if length == 2 {
		code = append(code, 0x66)
	}

	code = append(code, REX(w, r, x, b))
	code = append(code, opCode)
	code = append(code, ModRM(mod, byte(register), 0b100))
	code = append(code, SIB(scale, byte(offset), byte(base)))

	if mod == AddressMemoryOffset8 {
		code = append(code, 0x00)
	}

	return code
}