package arm

// ScaleFactor encodes the scale factor.
type ScaleFactor = byte

const (
	Scale1 = ScaleFactor(0b010)
	Scale2 = ScaleFactor(0b101)
	Scale4 = ScaleFactor(0b110)
	Scale8 = ScaleFactor(0b111)
)