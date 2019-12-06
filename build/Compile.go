package build

import (
	"fmt"

	"github.com/akyoto/q/build/assembler"
	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/instruction"
	"github.com/akyoto/q/build/register"
	"github.com/akyoto/q/build/token"
)

// Compile turns a function into machine code.
// It is executed for all function bodies.
func Compile(function *Function, environment *Environment, optimize bool, verbose bool) error {
	defer close(function.Finished)

	scopes := &ScopeStack{}
	scopes.Push()

	registers := register.NewManager()
	tokens := function.Tokens()
	identifierLifeTime := IdentifierLifeTimeMap(tokens)
	err := declareParameters(function, scopes, registers, identifierLifeTime)

	if err != nil {
		return err
	}

	instructions, instrErr := instruction.FromTokens(tokens)

	if instrErr != nil {
		return function.Error(instrErr.Position, instrErr)
	}

	assembler := assembler.New(verbose)
	assembler.AddLabel(function.Name)
	function.assembler = assembler

	state := State{
		assembler:          assembler,
		scopes:             scopes,
		registers:          registers,
		environment:        environment,
		function:           function,
		tokens:             tokens,
		instructions:       instructions,
		identifierLifeTime: identifierLifeTime,
		ignoreContracts:    false,
	}

	if optimize {
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

	// Contract failures
	for _, requirement := range state.requireState.list {
		assembler.AddLabel(requirement.failLabel)
		state.printLn(fmt.Sprintf("%s: require %v", state.function.Name, requirement.condition))
		state.assembler.MoveRegisterNumber(state.registers.Syscall[0], 60)
		state.assembler.MoveRegisterNumber(state.registers.Syscall[1], 1)
		state.assembler.Syscall()
	}

	// Optimize assembly code
	state.assembler.Optimize()

	return nil
}

// declareParameters declares the given parameters as variables inside the scope.
// It also assigns a register to each variable.
func declareParameters(function *Function, scopes *ScopeStack, registers *register.Manager, identifierLifeTime map[string]token.Position) error {
	for i, parameter := range function.Parameters {
		if i >= len(registers.Call) {
			return errors.ExceededMaxParameters
		}

		register := registers.Call[i]

		variable := &Variable{
			Name:       parameter.Name,
			Position:   0,
			AliveUntil: identifierLifeTime[parameter.Name],
		}

		_ = variable.SetRegister(register)
		scopes.Add(variable)
	}

	return nil
}
