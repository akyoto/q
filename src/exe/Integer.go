package exe

// Signed represents signed integers.
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Unsigned represents unsigned integers.
type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Integer represents any type of integer, signed or unsigned.
type Integer interface {
	Signed | Unsigned
}