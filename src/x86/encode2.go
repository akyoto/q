package x86

import "git.urbach.dev/cli/q/src/cpu"

// encode2 is the core function that encodes an instruction with an index.
func encode2(code []byte, mod AddressMode, reg cpu.Register, base cpu.Register, index cpu.Register, scale Scale, length byte, opCode uint32) []byte {
	var (
		w byte
		r byte
		x byte
		b byte
	)

	if length == 8 {
		w = 1
	} else {
		w = 0
	}

	r, reg = split(reg)
	x, index = split(index)
	b, base = split(base)

	if length == 2 {
		code = append(code, 0x66)
	}

	if opCode&0xFF000000 != 0 {
		code = append(code, byte(opCode>>24))
	}

	if w != 0 || r != 0 || x != 0 || b != 0 || (length == 1 && reg >= SP && reg <= R7) {
		code = append(code, REX(w, r, x, b))
	}

	if opCode&0xFF0000 != 0 {
		code = append(code, byte(opCode>>16))
	}

	if opCode&0xFF00 != 0 {
		code = append(code, byte(opCode>>8))
	}

	code = append(code, byte(opCode))
	code = append(code, ModRM(mod, byte(reg), 0b100))
	code = append(code, SIB(scale, byte(index), byte(base)))

	if mod == AddressMemoryOffset8 {
		code = append(code, 0x00)
	}

	return code
}