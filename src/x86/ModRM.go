package x86

// AddressMode encodes the addressing mode.
type AddressMode = byte

const (
	AddressMemory         = AddressMode(0b00)
	AddressMemoryOffset8  = AddressMode(0b01)
	AddressMemoryOffset32 = AddressMode(0b10)
	AddressDirect         = AddressMode(0b11)
)

// ModRM is used to generate a mode-register-memory suffix.
// - mod: 2 bits. The addressing mode.
// - reg: 3 bits. Register reference or opcode extension.
// - rm:  3 bits. Register operand.
func ModRM(mod AddressMode, reg byte, rm byte) byte {
	return (mod << 6) | (reg << 3) | rm
}