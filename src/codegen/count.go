package codegen

// counter is the data type for counters.
type counter uint16

// count stores how often a certain statement appeared so we can generate a unique label from it.
type count struct {
	Assert    counter
	Branch    counter
	Data      counter
	Loop      counter
	SubBranch counter
	Switch    counter
}