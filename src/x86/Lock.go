package x86

// Lock causes the next instruction to be atomic.
func Lock(code []byte) []byte {
	return append(code, 0xF0)
}