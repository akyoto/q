package build

import (
	"github.com/akyoto/asm"
	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/instruction"
	"github.com/akyoto/q/build/register"
)

// Compile turns a function into machine code.
// It is executed for all function bodies.
func Compile(function *Function, environment *Environment, optimize bool, verbose bool) (*asm.Assembler, error) {
	assembler := asm.New()
	assembler.AddLabel(function.Name)

	scopes := &ScopeStack{}
	scopes.Push()

	registers := register.NewManager()
	err := declareParameters(function, scopes, registers)

	if err != nil {
		return nil, err
	}

	tokens := function.Tokens()
	instructions, instrErr := instruction.FromTokens(tokens)

	if instrErr != nil {
		return nil, function.Error(instrErr.Position, instrErr)
	}

	state := State{
		assembler:    assembler,
		scopes:       scopes,
		registers:    registers,
		environment:  environment,
		function:     function,
		tokens:       tokens,
		instructions: instructions,
		optimize:     optimize,
		verbose:      verbose,
	}

	err = state.CompileInstructions()

	if err != nil {
		return nil, function.Error(state.tokenCursor, err)
	}

	for _, variable := range scopes.Unused() {
		return nil, function.Error(variable.Position, &errors.UnusedVariable{VariableName: variable.Name})
	}

	assembler.Return()
	return assembler, nil
}

// declareParameters declares the given parameters as variables inside the scope.
// It also assigns a register to each variable.
func declareParameters(function *Function, scopes *ScopeStack, registers *register.Manager) error {
	for i, parameter := range function.Parameters {
		if i >= len(registers.CallRegisters) {
			return errors.ExceededMaxParameters
		}

		register := registers.CallRegisters[i]

		variable := &Variable{
			Name:     parameter.Name,
			Position: 0,
		}

		variable.BindRegister(register)
		scopes.Add(variable)
	}

	return nil
}
