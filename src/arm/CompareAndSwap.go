package arm

import "git.urbach.dev/cli/q/src/cpu"

// CompareAndSwap compares `oldValue` with the memory at `address`.
// If they are equal, `newValue` register is written to memory.
// If they are not equal, `oldValue` is loaded with the value in memory.
func CompareAndSwap(oldValue cpu.Register, newValue cpu.Register, address cpu.Register, length byte) uint32 {
	return size(length)<<30 | 0b001000111<<21 | 0b111111<<10 | reg3(newValue, address, oldValue)
}