package x86

import (
	"git.urbach.dev/cli/q/src/cpu"
)

// ShiftLeftNumber shifts the register value by `bitCount` bits to the left.
func ShiftLeftNumber(code []byte, register cpu.Register, bitCount byte) []byte {
	code = encode(code, AddressDirect, 0b100, register, 8, 0xC1)
	return append(code, bitCount)
}

// ShiftRightSignedNumber shifts the signed register value by `bitCount` bits to the right.
func ShiftRightSignedNumber(code []byte, register cpu.Register, bitCount byte) []byte {
	code = encode(code, AddressDirect, 0b111, register, 8, 0xC1)
	return append(code, bitCount)
}