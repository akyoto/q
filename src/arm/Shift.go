package arm

import (
	"git.urbach.dev/cli/q/src/cpu"
)

// ShiftLeftNumber shifts the register value a specified amount of bits to the left.
func ShiftLeftNumber(destination cpu.Register, source cpu.Register, bits int) uint32 {
	return 0b110100110<<23 | reg2BitmaskImm(destination, source, 1, 64-bits, (^bits)&mask6)
}

// ShiftRightSignedNumber shifts the signed register value a specified amount of bits to the right.
func ShiftRightSignedNumber(destination cpu.Register, source cpu.Register, bits int) uint32 {
	return 0b100100110<<23 | reg2BitmaskImm(destination, source, 1, bits&mask6, 0b111111)
}