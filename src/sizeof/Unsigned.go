package sizeof

import "math"

// Unsigned tells you how many bytes are needed to encode this unsigned number.
func Unsigned[T uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64](number T) int {
	x := uint64(number)

	switch {
	case x <= math.MaxUint8:
		return 1

	case x <= math.MaxUint16:
		return 2

	case x <= math.MaxUint32:
		return 4

	default:
		return 8
	}
}