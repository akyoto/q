package arm

// AddressMode encodes the addressing mode.
type AddressMode uint32

const (
	UnscaledImmediate = AddressMode(0b00)
	PostIndex         = AddressMode(0b01)
	PreIndex          = AddressMode(0b11)
)