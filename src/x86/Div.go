package x86

import "git.urbach.dev/cli/q/src/cpu"

// DivRegister divides RDX:RAX by the value in the register.
func DivRegister(code []byte, divisor cpu.Register) []byte {
	rex := byte(0x48)

	if divisor > 0b111 {
		rex++
		divisor &= 0b111
	}

	return append(
		code,
		rex,
		0xF7,
		0xF8+byte(divisor),
	)
}