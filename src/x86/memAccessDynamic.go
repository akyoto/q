package x86

import "git.urbach.dev/cli/q/src/cpu"

// memAccessDynamic encodes a memory access using the value of a register as an offset.
func memAccessDynamic(code []byte, opCode8 byte, opCode32 byte, register cpu.Register, base cpu.Register, offset cpu.Register, scale ScaleFactor, length byte) []byte {
	var (
		w      = byte(0)
		r      = byte(0)
		x      = byte(0)
		b      = byte(0)
		opCode = opCode32
		mod    = AddressMemory
	)

	if length == 1 {
		opCode = opCode8
	}

	if offset == SP {
		offset, base = base, offset
	}

	if length == 8 {
		w = 1
	}

	if register > 0b111 {
		r = 1
		register &= 0b111
	}

	if offset > 0b111 {
		x = 1
		offset &= 0b111
	}

	if base > 0b111 {
		b = 1
		base &= 0b111
	}

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