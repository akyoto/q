package codegen

// counter is the data type for counters.
type counter uint16

// count stores how often a certain statement appeared so we can generate a unique label from it.
type count struct {
	Data      counter
	Branch    counter
	SubBranch counter
}