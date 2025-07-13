package x86

// AddressMode encodes the addressing mode.
type AddressMode = byte

const (
	AddressMemory         = AddressMode(0b00)
	AddressMemoryOffset8  = AddressMode(0b01)
	AddressMemoryOffset32 = AddressMode(0b10)
	AddressDirect         = AddressMode(0b11)
)