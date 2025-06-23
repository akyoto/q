package x86

// ExtendR0ToR2 doubles the size of R0 (RAX) by sign-extending it to R2 (RDX).
// This is also known as CQO.
func ExtendR0ToR2(code []byte) []byte {
	return append(code, 0x48, 0x99)
}