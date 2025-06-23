package x86

// Return transfers program control to a return address located on the top of the stack.
// The address is usually placed on the stack by a Call instruction.
func Return(code []byte) []byte {
	return append(code, 0xC3)
}