package build

import (
	"fmt"
	"log"
	"os"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/build/assembler"
	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/instruction"
	"github.com/akyoto/q/build/register"
)

// Compile turns a function into machine code.
// It is executed for all function bodies.
func Compile(function *Function, environment *Environment, optimize bool, verbose bool) (*asm.Assembler, error) {
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

	var logger *log.Logger

	if verbose {
		logPrefix := fmt.Sprintf("[%s] ", function.Name)
		logger = log.New(os.Stdout, logPrefix, 0)
	}

	assembler := assembler.New(logger)
	assembler.AddLabel(function.Name)

	state := State{
		assembler:    assembler,
		scopes:       scopes,
		registers:    registers,
		environment:  environment,
		function:     function,
		tokens:       tokens,
		instructions: instructions,
		log:          logger,
		optimize:     optimize,
	}

	err = state.CompileInstructions()

	if err != nil {
		return nil, function.Error(state.tokenCursor, err)
	}

	for _, variable := range scopes.Unused() {
		return nil, function.Error(variable.Position, &errors.UnusedVariable{VariableName: variable.Name})
	}

	assembler.Return()
	return assembler.Finalize(state.log), nil
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

		variable.BindRegister(register)
		scopes.Add(variable)
	}

	return nil
}
