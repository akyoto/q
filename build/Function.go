package build

import (
	"errors"

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

	for _, parameter := range function.Parameters {
		register := function.compiler.registerManager.FindFreeRegister()

		if register == nil {
			return errors.New("Exceeded maximum number of parameters")
		}

		variable := &Variable{
			Name:     parameter.Name,
			Register: register,
		}

		register.UsedBy = variable
		function.compiler.scopes.Add(variable)
	}

	err := function.compiler.Run()

	if err != nil {
		return err
	}

	function.compiler.assembler.Return()
	return nil
}
