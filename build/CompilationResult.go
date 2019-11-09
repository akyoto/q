package build

import "github.com/akyoto/asm"

// CompilationResult represents the result of a compilation.
// It includes a reference to the compiled function and the machine code.
type CompilationResult struct {
	Function  *Function
	Assembler *asm.Assembler
}
