package build

import (
	"github.com/akyoto/q/spec"
)

// Function represents a function.
type Function struct {
	Name        string
	Parameters  []spec.Variable
	ReturnTypes []spec.Type
	File        *File
	TokenStart  int
	TokenEnd    int
	compiler    *Compiler
}

// Compile compiles the function code.
func (function *Function) Compile() error {
	function.compiler.assembler.AddLabel(function.Name)
	err := function.compiler.Run()

	if err != nil {
		return err
	}

	function.compiler.assembler.Return()
	return nil
}
