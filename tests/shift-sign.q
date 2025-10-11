main() {
	assert rightShiftSigned(-1, 32) == -1
	assert rightShiftSigned32(-1) == -1
	assert rightShiftUnsigned(0xFFFFFFFFFFFFFFFF, 32) == 0xFFFFFFFF
	assert rightShiftUnsigned32(0xFFFFFFFFFFFFFFFF) == 0xFFFFFFFF
}

rightShiftSigned(x int64, n int) -> int64 {
	return x >> n
}

rightShiftSigned32(x int64) -> int64 {
	return x >> 32
}

rightShiftUnsigned(x uint64, n int) -> uint64 {
	return x >> n
}

rightShiftUnsigned32(x uint64) -> uint64 {
	return x >> 32
}