package x86

import "git.urbach.dev/cli/q/src/cpu"

// Load loads from memory with a register offset into a register.
func Load(code []byte, destination cpu.Register, base cpu.Register, offset cpu.Register, scale Scale, length byte) []byte {
	return memAccessDynamic(code, 0x8A, 0x8B, destination, base, offset, scale, length)
}

// LoadZeroExtend loads from memory with a register offset into a register and zero-extends it.
func LoadZeroExtend(code []byte, destination cpu.Register, base cpu.Register, offset cpu.Register, scale Scale, length byte) []byte {
	var (
		opCode = byte(0xB7)
		mod    = AddressMemory
	)

	if length == 1 {
		opCode = 0xB6
	}

	if offset == SP {
		offset, base = base, offset
	}

	w := byte(1)
	r, destination := split(destination)
	x, offset := split(offset)
	b, base := split(base)

	if base == R5 || base == R13 {
		mod = AddressMemoryOffset8
	}

	code = append(code, REX(w, r, x, b))
	code = append(code, 0x0F)
	code = append(code, opCode)
	code = append(code, ModRM(mod, byte(destination), 0b100))
	code = append(code, SIB(scale, byte(offset), byte(base)))

	if mod == AddressMemoryOffset8 {
		code = append(code, 0x00)
	}

	return code
}