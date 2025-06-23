package arm

// Nop does nothing. This can be used for alignment purposes.
func Nop() uint32 {
	return 0xD503201F
}