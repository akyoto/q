package arm

import (
	"git.urbach.dev/cli/q/src/cpu"
)

// StorePair calculates an address from a base register value and an immediate offset multiplied by 8,
// and stores the values of two registers to the calculated address.
// This is the pre-index version of the instruction so the offset is applied to the base register before the memory access.
func StorePair(reg1 cpu.Register, reg2 cpu.Register, base cpu.Register, offset int) uint32 {
	return 0b1010100110<<22 | pair(reg1, reg2, base, offset/8)
}