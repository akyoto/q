package x86

import "git.urbach.dev/cli/q/src/cpu"

// SetIfEqual sets the register to 1 if the result was equal, otherwise to 0.
func SetIfEqual(code []byte, destination cpu.Register) []byte {
	return append(code, 0x0F, 0x94, ModRM(AddressDirect, 0, byte(destination)))
}

// SetIfNotEqual sets the register to 1 if the result was not equal, otherwise to 0.
func SetIfNotEqual(code []byte, destination cpu.Register) []byte {
	return append(code, 0x0F, 0x95, ModRM(AddressDirect, 0, byte(destination)))
}

// SetIfGreater sets the register to 1 if the result was greater, otherwise to 0.
func SetIfGreater(code []byte, destination cpu.Register) []byte {
	return append(code, 0x0F, 0x9F, ModRM(AddressDirect, 0, byte(destination)))
}

// SetIfGreaterEqual sets the register to 1 if the result was greater or equal, otherwise to 0.
func SetIfGreaterEqual(code []byte, destination cpu.Register) []byte {
	return append(code, 0x0F, 0x9D, ModRM(AddressDirect, 0, byte(destination)))
}

// SetIfLess sets the register to 1 if the result was less, otherwise to 0.
func SetIfLess(code []byte, destination cpu.Register) []byte {
	return append(code, 0x0F, 0x9C, ModRM(AddressDirect, 0, byte(destination)))
}

// SetIfLessEqual sets the register to 1 if the result was less or equal, otherwise to 0.
func SetIfLessEqual(code []byte, destination cpu.Register) []byte {
	return append(code, 0x0F, 0x9E, ModRM(AddressDirect, 0, byte(destination)))
}

// SetIfUnsignedGreater sets the register to 1 if the result was greater using unsigned comparison, otherwise to 0.
func SetIfUnsignedGreater(code []byte, destination cpu.Register) []byte {
	return append(code, 0x0F, 0x97, ModRM(AddressDirect, 0, byte(destination)))
}

// SetIfUnsignedGreaterEqual sets the register to 1 if the result was greater or equal using unsigned comparison, otherwise to 0.
func SetIfUnsignedGreaterEqual(code []byte, destination cpu.Register) []byte {
	return append(code, 0x0F, 0x93, ModRM(AddressDirect, 0, byte(destination)))
}

// SetIfUnsignedLess sets the register to 1 if the result was less using unsigned comparison, otherwise to 0.
func SetIfUnsignedLess(code []byte, destination cpu.Register) []byte {
	return append(code, 0x0F, 0x92, ModRM(AddressDirect, 0, byte(destination)))
}

// SetIfUnsignedLessEqual sets the register to 1 if the result was less or equal using unsigned comparison, otherwise to 0.
func SetIfUnsignedLessEqual(code []byte, destination cpu.Register) []byte {
	return append(code, 0x0F, 0x96, ModRM(AddressDirect, 0, byte(destination)))
}