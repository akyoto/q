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
func Compile(function *Function, environment *Environment, optimize bool, verbose bool) {
	defer func() {
		function.Finished.L.Lock()
		function.IsFinished = true
		function.Finished.Broadcast()
		function.Finished.L.Unlock()
	}()

	scopes := &ScopeStack{}
	scopes.Push()

	registers := register.NewManager()
	tokens := function.Tokens()
	identifierLifeTime := IdentifierLifeTimeMap(tokens)

	// Parameters
	err := declareParameters(function, scopes, registers, identifierLifeTime)

	if err != nil {
		function.Error = err
		return
	}

	// Instructions
	instructions, instrErr := instruction.FromTokens(tokens)

	if instrErr != nil {
		function.Error = function.NewError(instrErr.Position, instrErr)
		return
	}

	// Assembler
	assembler := assembler.New(verbose)
	assembler.AddLabel(function.Name)
	function.assembler = assembler

	// State
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

	// Return types
	if len(function.ReturnTypeTokens) > 0 {
		typeName := function.ReturnTypeTokens[0].Text()
		typ := environment.Types[typeName]

		if typ == nil {
			err = errors.New(state.environment.UnknownTypeError(typeName))
			function.Error = NewError(err, function.File.path, function.File.tokens[:function.returnTypeStart+1], function)
			return
		}

		function.ReturnTypes = append(function.ReturnTypes, typ)
	}

	// Compile the function
	err = state.CompileInstructions()

	if err != nil {
		function.Error = function.NewError(state.tokenCursor, err)
		return
	}

	// Check for mistakes in variable usage
	err = state.PopScope(false)

	if err != nil {
		function.Error = err
		return
	}

	// Return
	if state.ensureState.counter > 0 {
		assembler.AddLabel("return")

		underscore := &Variable{Name: "_"}
		underscore.ForceSetRegister(registers.ReturnValue[0])
		state.scopes.Push()
		state.scopes.Add(underscore)

		for _, ensure := range state.ensureState.list {
			err := state.Condition(ensure.condition, ensure.failLabel)

			if err != nil {
				function.Error = err
				return
			}
		}

		registers.ReturnValue[0].Free()
	}

	assembler.Return()

	// Contract expect failures
	for _, expect := range state.expectState.list {
		assembler.AddLabel(expect.failLabel)
		state.printLn(fmt.Sprintf("%s: expect %v", state.function.Name, expect.condition))
		state.assembler.MoveRegisterNumber(state.registers.Syscall[0], 60)
		state.assembler.MoveRegisterNumber(state.registers.Syscall[1], 1)
		state.assembler.Syscall()
	}

	// Contract ensure failures
	for _, ensure := range state.ensureState.list {
		assembler.AddLabel(ensure.failLabel)
		state.printLn(fmt.Sprintf("%s: ensure %v", state.function.Name, ensure.condition))
		state.assembler.MoveRegisterNumber(state.registers.Syscall[0], 60)
		state.assembler.MoveRegisterNumber(state.registers.Syscall[1], 1)
		state.assembler.Syscall()
	}

	// Optimize assembly code
	state.assembler.Optimize()
}

// declareParameters declares the given parameters as variables inside the scope.
// It also assigns a register to each variable.
func declareParameters(function *Function, scopes *ScopeStack, registers *register.Manager, identifierLifeTime map[string]token.Position) error {
	for i, parameter := range function.Parameters {
		if i >= len(registers.Call) {
			return errors.New(errors.ExceededMaxParameters)
		}

		register := registers.Call[i]
		typeName := parameter.TypeTokens[0].Text()
		file := function.File
		parameter.Type = file.environment.Types[typeName]

		if parameter.Type == nil {
			return NewError(errors.New(&errors.UnknownType{Name: typeName}), file.path, file.tokens[:parameter.Position+2], function)
		}

		variable := &Variable{
			Name:       parameter.Name,
			Type:       parameter.Type,
			Position:   0,
			AliveUntil: identifierLifeTime[parameter.Name],
		}

		_ = variable.SetRegister(register)
		scopes.Add(variable)
	}

	return nil
}
