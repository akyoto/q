package x86

import "git.urbach.dev/cli/q/src/cpu"

// PopRegister pops a value from the stack and saves it into the register.
func PopRegister(code []byte, register cpu.Register) []byte {
	if register > 0b111 {
		code = append(code, REX(0, 0, 0, 1))
		register &= 0b111
	}

	return append(
		code,
		0x58+byte(register),
	)
}