package arm

// Jump continues program flow at the new 26 bit offset multiplied by 4.
func Jump(offset int) (code uint32, encodable bool) {
	if offset < -33554432 || offset > 33554431 {
		return 0, false
	}

	return 0b000101<<26 | uint32(offset&mask26), true
}

// JumpIfEqual jumps if the result was equal.
func JumpIfEqual(offset int) (code uint32, encodable bool) {
	return branchCond(EQ, offset)
}

// JumpIfNotEqual jumps if the result was not equal.
func JumpIfNotEqual(offset int) (code uint32, encodable bool) {
	return branchCond(NE, offset)
}

// JumpIfGreater jumps if the result was greater.
func JumpIfGreater(offset int) (code uint32, encodable bool) {
	return branchCond(GT, offset)
}

// JumpIfGreaterEqual jumps if the result was greater or equal.
func JumpIfGreaterEqual(offset int) (code uint32, encodable bool) {
	return branchCond(GE, offset)
}

// JumpIfLess jumps if the result was less.
func JumpIfLess(offset int) (code uint32, encodable bool) {
	return branchCond(LT, offset)
}

// JumpIfLessEqual jumps if the result was less or equal.
func JumpIfLessEqual(offset int) (code uint32, encodable bool) {
	return branchCond(LE, offset)
}

// JumpIfUnsignedGreater jumps if the result was greater using unsigned comparison.
func JumpIfUnsignedGreater(offset int) (code uint32, encodable bool) {
	return branchCond(HI, offset)
}

// JumpIfUnsignedGreaterEqual jumps if the result was greater or equal using unsigned comparison.
func JumpIfUnsignedGreaterEqual(offset int) (code uint32, encodable bool) {
	return branchCond(HS, offset)
}

// JumpIfUnsignedLess jumps if the result was less using unsigned comparison.
func JumpIfUnsignedLess(offset int) (code uint32, encodable bool) {
	return branchCond(LO, offset)
}

// JumpIfUnsignedLessEqual jumps if the result was less or equal using unsigned comparison.
func JumpIfUnsignedLessEqual(offset int) (code uint32, encodable bool) {
	return branchCond(LS, offset)
}

// branchCond performs a conditional branch to the new 19 bit offset multiplied by 4.
func branchCond(cond condition, offset int) (code uint32, encodable bool) {
	if offset < -262144 || offset > 262143 {
		return 0, false
	}

	return 0b01010100<<24 | uint32(offset&mask19)<<5 | uint32(cond), true
}