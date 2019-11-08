package build

import (
	"errors"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/build/register"
)

// Compile turns a function into machine code.
// It is executed for all function bodies.
func Compile(function *Function, environment *Environment) (*asm.Assembler, error) {
	assembler := asm.New()
	assembler.AddLabel(function.Name)

	scopes := &ScopeStack{}
	scopes.Push()

	registers := register.NewManager()
	err := declareParameters(function, scopes, registers)

	if err != nil {
		return nil, err
	}

	state := State{
		assembler:   assembler,
		scopes:      scopes,
		registers:   registers,
		environment: environment,
		function:    function,
		tokens:      function.Tokens(),
	}

	err = state.ProcessTokens()

	if err != nil {
		return nil, err
	}

	assembler.Return()
	return assembler, nil
}

// declareParameters declares the given parameters as variables inside the scope.
// It also assigns a register to each variable.
func declareParameters(function *Function, scopes *ScopeStack, registers *register.Manager) error {
	for _, parameter := range function.Parameters {
		register := registers.FindFreeRegister()

		if register == nil {
			return errors.New("Exceeded maximum number of parameters")
		}

		variable := &Variable{
			Name: parameter.Name,
		}

		variable.BindRegister(register)
		scopes.Add(variable)
	}

	return nil
}
