package arm

import "git.urbach.dev/cli/q/src/cpu"

// Loads thread-local storage base address into the destination register.
func LoadTLS(destination cpu.Register) uint32 {
	return 1<<21 | StoreTLS(destination)
}

// Stores source register as the thread-local storage base address.
func StoreTLS(source cpu.Register) uint32 {
	return uint32(0b1101_0101_0001_1011_1101_0000_0100_0000) | uint32(source)
}