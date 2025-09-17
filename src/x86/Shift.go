package x86

import (
	"git.urbach.dev/cli/q/src/cpu"
)

// ShiftLeft shifts the register value by the amount of bits defined in the lower byte of R1 to the left.
func ShiftLeft(code []byte, register cpu.Register) []byte {
	return encode(code, AddressDirect, 0b100, register, 8, 0xD3)
}

// ShiftRight shifts the register value by the amount of bits defined in the lower byte of R1 to the right.
func ShiftRight(code []byte, register cpu.Register) []byte {
	return encode(code, AddressDirect, 0b101, register, 8, 0xD3)
}

// ShiftRightSigned shifts the signed register value by the amount of bits defined in the lower byte of R1 to the right.
func ShiftRightSigned(code []byte, register cpu.Register) []byte {
	return encode(code, AddressDirect, 0b111, register, 8, 0xD3)
}

// ShiftLeftNumber shifts the register value by `bitCount` bits to the left.
func ShiftLeftNumber(code []byte, register cpu.Register, bitCount byte) []byte {
	code = encode(code, AddressDirect, 0b100, register, 8, 0xC1)
	return append(code, bitCount)
}

// ShiftRightNumber shifts the register value by `bitCount` bits to the right.
func ShiftRightNumber(code []byte, register cpu.Register, bitCount byte) []byte {
	code = encode(code, AddressDirect, 0b101, register, 8, 0xC1)
	return append(code, bitCount)
}

// ShiftRightSignedNumber shifts the signed register value by `bitCount` bits to the right.
func ShiftRightSignedNumber(code []byte, register cpu.Register, bitCount byte) []byte {
	code = encode(code, AddressDirect, 0b111, register, 8, 0xC1)
	return append(code, bitCount)
}