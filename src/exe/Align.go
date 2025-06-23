package exe

// Align calculates the next aligned address.
func Align[T int | uint | int64 | uint64 | int32 | uint32](n T, alignment T) T {
	return (n + (alignment - 1)) & ^(alignment - 1)
}

// AlignPad calculates the next aligned address and the padding needed.
func AlignPad[T int | uint | int64 | uint64 | int32 | uint32](n T, alignment T) (T, T) {
	aligned := Align(n, alignment)
	return aligned, aligned - n
}