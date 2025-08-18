htons(num uint16) -> uint16 {
	return ((num & 0xFF) << 8) | (num >> 8)
}