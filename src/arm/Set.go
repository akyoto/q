package arm

import "git.urbach.dev/cli/q/src/cpu"

// SetIfEqual sets the register to 1 if the result was equal, otherwise to 0.
func SetIfEqual(destination cpu.Register) uint32 {
	return cset(EQ, destination)
}

// SetIfNotEqual sets the register to 1 if the result was not equal, otherwise to 0.
func SetIfNotEqual(destination cpu.Register) uint32 {
	return cset(NE, destination)
}

// SetIfGreater sets the register to 1 if the result was greater, otherwise to 0.
func SetIfGreater(destination cpu.Register) uint32 {
	return cset(GT, destination)
}

// SetIfGreaterEqual sets the register to 1 if the result was greater or equal, otherwise to 0.
func SetIfGreaterEqual(destination cpu.Register) uint32 {
	return cset(GE, destination)
}

// SetIfLess sets the register to 1 if the result was less, otherwise to 0.
func SetIfLess(destination cpu.Register) uint32 {
	return cset(LT, destination)
}

// SetIfLessEqual sets the register to 1 if the result was less or equal, otherwise to 0.
func SetIfLessEqual(destination cpu.Register) uint32 {
	return cset(LE, destination)
}

// SetIfUnsignedGreater sets the register to 1 if the result was greater using unsigned comparison, otherwise to 0.
func SetIfUnsignedGreater(destination cpu.Register) uint32 {
	return cset(HI, destination)
}

// SetIfUnsignedGreaterEqual sets the register to 1 if the result was greater or equal using unsigned comparison, otherwise to 0.
func SetIfUnsignedGreaterEqual(destination cpu.Register) uint32 {
	return cset(HS, destination)
}

// SetIfUnsignedLess sets the register to 1 if the result was less using unsigned comparison, otherwise to 0.
func SetIfUnsignedLess(destination cpu.Register) uint32 {
	return cset(LO, destination)
}

// SetIfUnsignedLessEqual sets the register to 1 if the result was less or equal using unsigned comparison, otherwise to 0.
func SetIfUnsignedLessEqual(destination cpu.Register) uint32 {
	return cset(LS, destination)
}

// cset encodes a conditional set instruction.
func cset(cond condition, destination cpu.Register) uint32 {
	return 0b10011010100<<21 | uint32(cond^1)<<12 | 1<<10 | reg3(destination, 0b11111, 0b11111)
}