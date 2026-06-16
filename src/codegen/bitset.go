package codegen

import "git.urbach.dev/cli/q/src/cpu"

// bitSet is used to mark used registers.
type bitSet uint64

// Set marks the given register as used.
func (b *bitSet) Set(reg cpu.Register) {
	*b |= 1 << (reg & 63)
}

// Has returns true if the given register is used.
func (b *bitSet) Has(reg cpu.Register) bool {
	return *b&(1<<(reg&63)) != 0
}