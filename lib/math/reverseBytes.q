reverseBytes(num uint16) -> uint16 {
	return (num << 8) | (num >> 8)
}