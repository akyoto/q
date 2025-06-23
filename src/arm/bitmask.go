package arm

import "math/bits"

// encodeLogicalImmediate encodes a bitmask immediate.
// The algorithm used here was made by Dougall Johnson.
func encodeLogicalImmediate(val uint) (N int, immr int, imms int, encodable bool) {
	if val == 0 || ^val == 0 {
		return 0, 0, 0, false
	}

	rotation := bits.TrailingZeros(clearTrailingOnes(val))
	normalized := bits.RotateLeft(val, -(rotation & 63))

	zeroes := bits.LeadingZeros(normalized)
	ones := bits.TrailingZeros(^normalized)
	size := zeroes + ones

	immr = -rotation & (size - 1)
	imms = -(size << 1) | (ones - 1)
	N = (size >> 6)

	if bits.RotateLeft(val, -(size&63)) != val {
		return 0, 0, 0, false
	}

	return N, immr, (imms & 0x3F), true
}

// clearTrailingOnes clears trailing one bits.
func clearTrailingOnes(x uint) uint {
	return x & (x + 1)
}