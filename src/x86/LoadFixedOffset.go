package x86

import "git.urbach.dev/cli/q/src/cpu"

// LoadFixedOffset loads from memory with a signed offset from -128 to 127 into a register.
func LoadFixedOffset(code []byte, destination cpu.Register, base cpu.Register, offset int8, scale Scale, length byte) []byte {
	return memAccess(code, destination, base, offset, scale, length, 0x8A, 0x8B)
}

// LoadFixedOffsetSignExtend loads from memory with a signed offset from -128 to 127 into a register and sign-extends it.
func LoadFixedOffsetSignExtend(code []byte, destination cpu.Register, base cpu.Register, offset int8, scale Scale, length byte) []byte {
	var (
		opCode = byte(0xBF)
		mod    = AddressMemory
	)

	if length == 1 {
		opCode = 0xBE
	}

	if length == 4 {
		opCode = 0x63
	}

	if offset != 0 || base == R5 || base == R13 {
		mod = AddressMemoryOffset8
	}

	w := byte(1)
	r, destination := split(destination)
	b, base := split(base)

	code = append(code, REX(w, r, 0, b))

	if length != 4 {
		code = append(code, 0x0F)
	}

	code = append(code, opCode)
	code = append(code, ModRM(mod, byte(destination), byte(base)))

	if base == SP || base == R12 {
		code = append(code, SIB(scale, 0b100, 0b100))
	}

	if mod == AddressMemoryOffset8 {
		code = append(code, byte(offset))
	}

	return code
}