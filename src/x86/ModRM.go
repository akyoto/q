package x86

// ModRM is used to generate a mode-register-memory suffix.
// - mod: 2 bits. The addressing mode.
// - reg: 3 bits. Second register operand or opcode extension denoted by /digit.
// - rm:  3 bits. First register operand.
func ModRM(mod AddressMode, reg byte, rm byte) byte {
	return (mod << 6) | (reg << 3) | rm
}