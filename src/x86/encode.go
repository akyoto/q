package x86

import "git.urbach.dev/cli/q/src/cpu"

// encode is the core function that encodes an instruction.
func encode(code []byte, mod AddressMode, reg cpu.Register, rm cpu.Register, length byte, opCode uint32) []byte {
	var (
		w byte
		r byte
		b byte
	)

	if length == 8 {
		w = 1
	} else {
		w = 0
	}

	r, reg = split(reg)
	b, rm = split(rm)

	if length == 2 {
		code = append(code, 0x66)
	}

	if opCode&0xFF000000 != 0 {
		code = append(code, byte(opCode>>24))
	}

	if w != 0 || r != 0 || b != 0 || (length == 1 && reg >= SP && reg <= R7) {
		code = append(code, REX(w, r, 0, b))
	}

	if opCode&0xFF0000 != 0 {
		code = append(code, byte(opCode>>16))
	}

	if opCode&0xFF00 != 0 {
		code = append(code, byte(opCode>>8))
	}

	code = append(code, byte(opCode))
	code = append(code, ModRM(mod, byte(reg), byte(rm)))

	if mod != AddressDirect && rm == 0b100 {
		code = append(code, SIB(Scale1, 0b100, 0b100))
	}

	return code
}