reverseBytes(num uint16) -> uint16 {
	return (num << 8) | (num >> 8)
}

rotateLeft(num uint64, n int) -> uint64 {
	return (num << n) | (num >> (64 - n))
}