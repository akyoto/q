package exe

import (
	"bytes"
)

// Align calculates the next aligned address (alignment must be a power of 2).
func Align[T Integer](n T, alignment T) T {
	return (n + (alignment - 1)) & ^(alignment - 1)
}

// AlignPad calculates the next aligned address and the padding needed (alignment must be a power of 2).
func AlignPad[T Integer](n T, alignment T) (T, T) {
	aligned := Align(n, alignment)
	return aligned, aligned - n
}

// Pad calculates the padding (alignment must be a power of 2).
func Pad[T Integer](n T, alignment T) T {
	return -n & (alignment - 1)
}

// PadBuffer pads the buffer to the given alignment (alignment must be a power of 2).
func PadBuffer[T Integer](b *bytes.Buffer, alignment T) {
	padding := Pad(T(b.Len()), alignment)

	if padding > 0 {
		b.Write(make([]byte, padding))
	}
}

// PadSlice pads the slice to the given alignment (alignment must be a power of 2).
func PadSlice[E any, T Integer](slice []E, alignment T) []E {
	pad := Pad(T(len(slice)), alignment)

	if pad == 0 {
		return slice
	}

	if T(cap(slice)) >= T(len(slice))+pad {
		old := T(len(slice))
		slice = slice[:old+pad]
		clear(slice[old:])
		return slice
	}

	return append(slice, make([]E, pad)...)
}