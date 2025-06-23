package arm

// Call branches to a PC-relative offset, setting the register X30 to PC+4.
// The offset starts from the address of this instruction and is encoded as "imm26" times 4.
// This instruction is also known as BL (branch with link).
func Call(offset int) uint32 {
	return uint32(0b100101<<26) | uint32(offset&mask26)
}