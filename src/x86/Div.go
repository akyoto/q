package x86

import "git.urbach.dev/cli/q/src/cpu"

// DivSignedRegister performs signed division in RDX:RAX by the value in the register.
func DivSignedRegister(code []byte, divisor cpu.Register) []byte {
	return div(code, divisor, 0xF8)
}

// DivUnsignedRegister performs unsigned division in RDX:RAX by the value in the register.
func DivUnsignedRegister(code []byte, divisor cpu.Register) []byte {
	return div(code, divisor, 0xF0)
}

// div implements the encoding for the division operation.
func div(code []byte, divisor cpu.Register, opCode byte) []byte {
	rex := byte(0x48)

	if divisor > 0b111 {
		rex++
		divisor &= 0b111
	}

	return append(
		code,
		rex,
		0xF7,
		opCode+byte(divisor),
	)
}