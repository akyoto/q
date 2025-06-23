package arm

// Jump continues program flow at the new offset.
func Jump(offset int) uint32 {
	return 0b000101<<26 | uint32(offset&mask26)
}

// JumpIfEqual jumps if the result was equal.
func JumpIfEqual(offset int) uint32 {
	return branchCond(EQ, offset)
}

// JumpIfNotEqual jumps if the result was not equal.
func JumpIfNotEqual(offset int) uint32 {
	return branchCond(NE, offset)
}

// JumpIfGreater jumps if the result was greater.
func JumpIfGreater(offset int) uint32 {
	return branchCond(GT, offset)
}

// JumpIfGreaterOrEqual jumps if the result was greater or equal.
func JumpIfGreaterOrEqual(offset int) uint32 {
	return branchCond(GE, offset)
}

// JumpIfLess jumps if the result was less.
func JumpIfLess(offset int) uint32 {
	return branchCond(LS, offset)
}

// JumpIfLessOrEqual jumps if the result was less or equal.
func JumpIfLessOrEqual(offset int) uint32 {
	return branchCond(LE, offset)
}

// branchCond performs a conditional branch to a PC-relative offset.
func branchCond(cond condition, imm19 int) uint32 {
	return 0b01010100<<24 | uint32(imm19&mask19)<<5 | uint32(cond)
}