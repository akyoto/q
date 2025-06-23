package arm

import "git.urbach.dev/cli/q/src/cpu"

// NegateRegister negates the value in the source register and writes it to the destination register.
func NegateRegister(destination cpu.Register, source cpu.Register) uint32 {
	return 0b11001011<<24 | reg3Imm(destination, ZR, source, 0)
}