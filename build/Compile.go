package build

import (
	"github.com/akyoto/q/build/assembler"
	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/instruction"
	"github.com/akyoto/q/build/register"
)

// Compile turns a function into machine code.
// It is executed for all function bodies.
func Compile(function *Function, environment *Environment, optimize bool, verbose bool) error {
	defer close(function.Finished)

	scopes := &ScopeStack{}
	scopes.Push()

	registers := register.NewManager()
	err := declareParameters(function, scopes, registers)

	if err != nil {
		return err
	}

	tokens := function.Tokens()
	instructions, instrErr := instruction.FromTokens(tokens)

	if instrErr != nil {
		return function.Error(instrErr.Position, instrErr)
	}

	assembler := assembler.New(verbose)
	assembler.AddLabel(function.Name)
	function.assembler = assembler

	state := State{
		assembler:    assembler,
		scopes:       scopes,
		registers:    registers,
		environment:  environment,
		function:     function,
		tokens:       tokens,
		instructions: instructions,
	}

	if optimize {
		state.useIncDec = true
		state.ignoreContracts = true
	}

	// Compile the function
	err = state.CompileInstructions()

	if err != nil {
		return function.Error(state.tokenCursor, err)
	}

	// Check for unused variables
	for _, variable := range scopes.Unused() {
		return function.Error(variable.Position, &errors.UnusedVariable{VariableName: variable.Name})
	}

	// End with a return statement
	assembler.Return()
	return nil
}

// declareParameters declares the given parameters as variables inside the scope.
// It also assigns a register to each variable.
func declareParameters(function *Function, scopes *ScopeStack, registers *register.Manager) error {
	for i, parameter := range function.Parameters {
		if i >= len(registers.Call) {
			return errors.ExceededMaxParameters
		}

		register := registers.Call[i]

		variable := &Variable{
			Name:     parameter.Name,
			Position: 0,
		}

		_ = variable.SetRegister(register)
		scopes.Add(variable)
	}

	return nil
}
