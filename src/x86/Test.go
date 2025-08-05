package x86

import "git.urbach.dev/cli/q/src/cpu"

// TestRegisterRegister computes the bitwise logical AND of the operands
// and sets the SF, ZF, and PF status flags according to the result.
func TestRegisterRegister(code []byte, registerA cpu.Register, registerB cpu.Register) []byte {
	return encode(code, AddressDirect, registerB, registerA, 8, 0x85)
}