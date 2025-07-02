package ssa2asm

// Counter is the data type for counters.
type Counter uint16

// Count stores how often a certain statement appeared so we can generate a unique label from it.
type Count struct {
	Data Counter
}