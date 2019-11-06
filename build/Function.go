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
	Compiler    *Compiler
	TokenStart  int
	TokenEnd    int
}

// Compile compiles the function code.
func (function *Function) Compile() error {
	compiler := NewCompiler(function.File.tokens[function.TokenStart:function.TokenEnd], function.File.build)
	function.Compiler = compiler
	compiler.assembler.AddLabel(function.Name)
	err := compiler.Run()

	if err != nil {
		return err
	}

	compiler.assembler.Return()
	return nil
}
