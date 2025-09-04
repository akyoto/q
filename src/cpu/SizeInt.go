package cpu

import "math"

// SizeInt tells you how many bytes are needed to encode this signed number.
func SizeInt[T int | int8 | int16 | int32 | int64](number T) int {
	x := int64(number)

	switch {
	case x >= math.MinInt8 && x <= math.MaxInt8:
		return 1

	case x >= math.MinInt16 && x <= math.MaxInt16:
		return 2

	case x >= math.MinInt32 && x <= math.MaxInt32:
		return 4

	default:
		return 8
	}
}