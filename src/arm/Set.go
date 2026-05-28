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

// SetIfGreaterOrEqual sets the register to 1 if the result was greater or equal, otherwise to 0.
func SetIfGreaterOrEqual(destination cpu.Register) uint32 {
	return cset(GE, destination)
}

// SetIfLess sets the register to 1 if the result was less, otherwise to 0.
func SetIfLess(destination cpu.Register) uint32 {
	return cset(LT, destination)
}

// SetIfLessOrEqual sets the register to 1 if the result was less or equal, otherwise to 0.
func SetIfLessOrEqual(destination cpu.Register) uint32 {
	return cset(LE, destination)
}

// cset encodes a conditional set instruction.
func cset(cond condition, destination cpu.Register) uint32 {
	return 0b10011010100<<21 | uint32(cond^1)<<12 | 1<<10 | reg3(destination, 0b11111, 0b11111)
}