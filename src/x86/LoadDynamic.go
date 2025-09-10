package x86

import "git.urbach.dev/cli/q/src/cpu"

// LoadDynamicRegister loads from memory with a register offset into a register.
func LoadDynamicRegister(code []byte, destination cpu.Register, base cpu.Register, offset cpu.Register, scale Scale, length byte) []byte {
	return memAccessDynamic(code, 0x8A, 0x8B, destination, base, offset, scale, length)
}

// LoadDynamicRegisterZeroExtend loads from memory with a register offset into a register and zero-extends it.
func LoadDynamicRegisterZeroExtend(code []byte, destination cpu.Register, base cpu.Register, offset cpu.Register, scale Scale, length byte) []byte {
	var (
		w      = byte(1)
		r      = byte(0)
		x      = byte(0)
		b      = byte(0)
		opCode = byte(0xB7)
		mod    = AddressMemory
	)

	if length == 1 {
		opCode = 0xB6
	}

	if offset == SP {
		offset, base = base, offset
	}

	if destination > 0b111 {
		r = 1
		destination &= 0b111
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