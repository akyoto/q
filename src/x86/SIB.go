package x86

// ScaleFactor encodes the scale factor.
type ScaleFactor = byte

const (
	Scale1 = ScaleFactor(0b00)
	Scale2 = ScaleFactor(0b01)
	Scale4 = ScaleFactor(0b10)
	Scale8 = ScaleFactor(0b11)
)

// SIB is used to generate an SIB byte.
// - scale: 2 bits. Multiplies the value of the index.
// - index: 3 bits. Specifies the index register.
// - base:  3 bits. Specifies the base register.
func SIB(scale ScaleFactor, index byte, base byte) byte {
	return (scale << 6) | (index << 3) | base
}