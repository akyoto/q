package x86

// Lock causes the next instruction to be atomic.
func Lock(code []byte) []byte {
	return append(code, 0xF0)
}

// SegmentBaseFS causes the next instruction to use FS as the segment base address.
func SegmentBaseFS(code []byte) []byte {
	return append(code, 0x64)
}

// SegmentBaseGS causes the next instruction to use GS as the segment base address.
func SegmentBaseGS(code []byte) []byte {
	return append(code, 0x65)
}