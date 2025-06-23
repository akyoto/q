package x86

import "git.urbach.dev/cli/q/src/cpu"

// NegateRegister negates the value in the register.
func NegateRegister(code []byte, register cpu.Register) []byte {
	return encode(code, AddressDirect, 0b011, register, 8, 0xF7)
}