package x86

import (
	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/sizeof"
)

// PushNumber pushes a number onto the stack.
func PushNumber(code []byte, number int32) []byte {
	length := sizeof.Signed(number)

	if length >= 2 {
		return append(
			code,
			0x68,
			byte(number),
			byte(number>>8),
			byte(number>>16),
			byte(number>>24),
		)
	}

	return append(code, 0x6A, byte(number))
}

// PushRegister pushes the value inside the register onto the stack.
func PushRegister(code []byte, register cpu.Register) []byte {
	if register > 0b111 {
		code = append(code, REX(0, 0, 0, 1))
		register &= 0b111
	}

	return append(
		code,
		0x50+byte(register),
	)
}