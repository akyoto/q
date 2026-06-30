package x86

// Jump continues program flow at the new offset.
// The offset is relative to the next instruction.
func Jump8(code []byte, offset int8) []byte {
	return append(code, 0xEB, byte(offset))
}

// JumpIfEqual jumps if the result was equal.
func Jump8IfEqual(code []byte, offset int8) []byte {
	return append(code, 0x74, byte(offset))
}

// JumpIfNotEqual jumps if the result was not equal.
func Jump8IfNotEqual(code []byte, offset int8) []byte {
	return append(code, 0x75, byte(offset))
}

// JumpIfGreater jumps if the result was greater.
func Jump8IfGreater(code []byte, offset int8) []byte {
	return append(code, 0x7F, byte(offset))
}

// JumpIfGreaterOrEqual jumps if the result was greater or equal.
func Jump8IfGreaterEqual(code []byte, offset int8) []byte {
	return append(code, 0x7D, byte(offset))
}

// JumpIfLess jumps if the result was less.
func Jump8IfLess(code []byte, offset int8) []byte {
	return append(code, 0x7C, byte(offset))
}

// JumpIfLessOrEqual jumps if the result was less or equal.
func Jump8IfLessEqual(code []byte, offset int8) []byte {
	return append(code, 0x7E, byte(offset))
}

// Jump8IfUnsignedGreater jumps if the result was greater using unsigned comparison.
func Jump8IfUnsignedGreater(code []byte, offset int8) []byte {
	return append(code, 0x77, byte(offset))
}

// Jump8IfUnsignedGreaterEqual jumps if the result was greater or equal using unsigned comparison.
func Jump8IfUnsignedGreaterEqual(code []byte, offset int8) []byte {
	return append(code, 0x73, byte(offset))
}

// Jump8IfUnsignedLess jumps if the result was less using unsigned comparison.
func Jump8IfUnsignedLess(code []byte, offset int8) []byte {
	return append(code, 0x72, byte(offset))
}

// Jump8IfUnsignedLessEqual jumps if the result was less or equal using unsigned comparison.
func Jump8IfUnsignedLessEqual(code []byte, offset int8) []byte {
	return append(code, 0x76, byte(offset))
}