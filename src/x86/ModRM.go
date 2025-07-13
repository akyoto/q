package x86

// ModRM is used to generate a mode-register-memory suffix.
// - mod: 2 bits. The addressing mode.
// - reg: 3 bits. Register reference or opcode extension.
// - rm:  3 bits. Register operand.
func ModRM(mod AddressMode, reg byte, rm byte) byte {
	return (mod << 6) | (reg << 3) | rm
}