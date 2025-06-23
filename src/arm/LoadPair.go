package arm

import "git.urbach.dev/cli/q/src/cpu"

// LoadPair calculates an address from a base register value and an immediate offset,
// loads two 64-bit doublewords from memory, and writes them to two registers.
// This is the post-index version of the instruction so the offset is applied to the base register after the memory access.
func LoadPair(reg1 cpu.Register, reg2 cpu.Register, base cpu.Register, offset int) uint32 {
	return 0b1010100011<<22 | pair(reg1, reg2, base, offset/8)
}