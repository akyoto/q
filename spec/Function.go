package spec

import "github.com/akyoto/asm"

// Function represents a function.
type Function struct {
	Name        string
	Parameters  []Variable
	ReturnTypes []Type
	Code        *asm.Assembler
}
