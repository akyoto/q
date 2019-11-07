package build

import (
	"github.com/akyoto/q/spec"
)

// Function represents a function.
type Function struct {
	Name             string
	Parameters       []Variable
	ReturnTypes      []spec.Type
	TokenStart       int
	TokenEnd         int
	NoParameterCheck bool
	parameterStart   int
	compiler         *Compiler
}

// Compile compiles the function code.
func (function *Function) Compile() error {
	function.compiler.assembler.AddLabel(function.Name)

	for _, variable := range function.Parameters {
		function.compiler.scopes.Add(&Variable{
			Name:     variable.Name,
			Register: variableRegisters[function.compiler.registerCounter],
		})

		function.compiler.registerCounter++
	}

	err := function.compiler.Run()

	if err != nil {
		return err
	}

	function.compiler.assembler.Return()
	return nil
}
