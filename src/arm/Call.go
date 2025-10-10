package arm

import "git.urbach.dev/cli/q/src/cpu"

// Call branches to a PC-relative offset, setting the register X30 to PC+4.
// The offset starts from the address of this instruction and is encoded as "imm26" times 4.
// This instruction is also known as BL (branch with link).
func Call(offset int) (code uint32, encodable bool) {
	if offset < -33554432 || offset > 33554431 {
		return 0, false
	}

	return uint32(0b100101<<26) | uint32(offset&mask26), true
}

// Calls a function whose address is stored in the given register.
func CallRegister(register cpu.Register) uint32 {
	return uint32(0b1101011000111111<<16) | uint32(register)<<5
}