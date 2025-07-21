package arm

import (
	"git.urbach.dev/cli/q/src/cpu"
)

// ShiftLeft shifts the register value a specified amount of bits to the left.
func ShiftLeft(destination cpu.Register, source cpu.Register, operand cpu.Register) uint32 {
	return 0b10011010110<<21 | 0b001000<<10 | reg3(destination, source, operand)
}

// ShiftRightSigned shifts the signed register value a specified amount of bits to the right.
func ShiftRightSigned(destination cpu.Register, source cpu.Register, operand cpu.Register) uint32 {
	return 0b10011010110<<21 | 0b001010<<10 | reg3(destination, source, operand)
}

// ShiftLeftNumber shifts the register value a specified amount of bits to the left.
func ShiftLeftNumber(destination cpu.Register, source cpu.Register, bits int) uint32 {
	return 0b110100110<<23 | reg2BitmaskImm(destination, source, 1, 64-bits, (^bits)&mask6)
}

// ShiftRightSignedNumber shifts the signed register value a specified amount of bits to the right.
func ShiftRightSignedNumber(destination cpu.Register, source cpu.Register, bits int) uint32 {
	return 0b100100110<<23 | reg2BitmaskImm(destination, source, 1, bits&mask6, 0b111111)
}