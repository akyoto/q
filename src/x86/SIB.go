package x86

// Scale encodes the scale factor.
type Scale = byte

const (
	Scale1 = Scale(0b00)
	Scale2 = Scale(0b01)
	Scale4 = Scale(0b10)
	Scale8 = Scale(0b11)
)

// SIB is used to generate an SIB byte.
// - scale: 2 bits. Multiplies the value of the index.
// - index: 3 bits. Specifies the index register.
// - base:  3 bits. Specifies the base register.
func SIB(scale Scale, index byte, base byte) byte {
	return (scale << 6) | (index << 3) | base
}