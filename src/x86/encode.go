package x86

import "git.urbach.dev/cli/q/src/cpu"

// encode is the core function that encodes an instruction.
func encode(code []byte, mod AddressMode, reg cpu.Register, rm cpu.Register, numBytes byte, opCodes ...byte) []byte {
	w := byte(0) // Indicates a 64-bit register.
	r := byte(0) // Extension to the "reg" field in ModRM.
	x := byte(0) // Extension to the SIB index field.
	b := byte(0) // Extension to the "rm" field in ModRM or the SIB base (r8 up to r15 use this).

	if numBytes == 8 {
		w = 1
	}

	if reg > 0b111 {
		r = 1
		reg &= 0b111
	}

	if rm > 0b111 {
		b = 1
		rm &= 0b111
	}

	if w != 0 || r != 0 || x != 0 || b != 0 || (numBytes == 1 && (reg == SP || reg == R5 || reg == R6 || reg == R7)) {
		code = append(code, REX(w, r, x, b))
	}

	code = append(code, opCodes...)
	code = append(code, ModRM(mod, byte(reg), byte(rm)))
	return code
}