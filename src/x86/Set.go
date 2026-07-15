package x86

import "git.urbach.dev/cli/q/src/cpu"

// SetIfEqual sets the register to 1 if the result was equal, otherwise to 0.
func SetIfEqual(code []byte, destination cpu.Register) []byte {
	return encode(code, AddressDirect, 0, destination, 1, 0x0F94)
}

// SetIfNotEqual sets the register to 1 if the result was not equal, otherwise to 0.
func SetIfNotEqual(code []byte, destination cpu.Register) []byte {
	return encode(code, AddressDirect, 0, destination, 1, 0x0F95)
}

// SetIfGreater sets the register to 1 if the result was greater, otherwise to 0.
func SetIfGreater(code []byte, destination cpu.Register) []byte {
	return encode(code, AddressDirect, 0, destination, 1, 0x0F9F)
}

// SetIfGreaterEqual sets the register to 1 if the result was greater or equal, otherwise to 0.
func SetIfGreaterEqual(code []byte, destination cpu.Register) []byte {
	return encode(code, AddressDirect, 0, destination, 1, 0x0F9D)
}

// SetIfLess sets the register to 1 if the result was less, otherwise to 0.
func SetIfLess(code []byte, destination cpu.Register) []byte {
	return encode(code, AddressDirect, 0, destination, 1, 0x0F9C)
}

// SetIfLessEqual sets the register to 1 if the result was less or equal, otherwise to 0.
func SetIfLessEqual(code []byte, destination cpu.Register) []byte {
	return encode(code, AddressDirect, 0, destination, 1, 0x0F9E)
}

// SetIfUnsignedGreater sets the register to 1 if the result was greater using unsigned comparison, otherwise to 0.
func SetIfUnsignedGreater(code []byte, destination cpu.Register) []byte {
	return encode(code, AddressDirect, 0, destination, 1, 0x0F97)
}

// SetIfUnsignedGreaterEqual sets the register to 1 if the result was greater or equal using unsigned comparison, otherwise to 0.
func SetIfUnsignedGreaterEqual(code []byte, destination cpu.Register) []byte {
	return encode(code, AddressDirect, 0, destination, 1, 0x0F93)
}

// SetIfUnsignedLess sets the register to 1 if the result was less using unsigned comparison, otherwise to 0.
func SetIfUnsignedLess(code []byte, destination cpu.Register) []byte {
	return encode(code, AddressDirect, 0, destination, 1, 0x0F92)
}

// SetIfUnsignedLessEqual sets the register to 1 if the result was less or equal using unsigned comparison, otherwise to 0.
func SetIfUnsignedLessEqual(code []byte, destination cpu.Register) []byte {
	return encode(code, AddressDirect, 0, destination, 1, 0x0F96)
}