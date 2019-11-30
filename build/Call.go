package build

import (
	"fmt"
	"sync/atomic"

	"github.com/akyoto/asm/syscall"
	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/expression"
	"github.com/akyoto/q/build/register"
	"github.com/akyoto/q/build/token"
)

// CallExpression executes a function call.
func (state *State) CallExpression(expr *expression.Expression) error {
	functionName := expr.Token.Text()
	function := state.environment.Functions[functionName]
	isBuiltin := false

	if function == nil {
		function = BuiltinFunctions[functionName]
		isBuiltin = true
	}

	if function == nil {
		return state.UnknownFunctionError(functionName)
	}

	parameters := expr.Children

	// Calling a function with side effects causes our function to have side effects
	if atomic.LoadInt32(&function.SideEffects) > 0 {
		atomic.AddInt32(&state.function.SideEffects, 1)
	}

	// Parameter check
	if !function.NoParameterCheck && len(parameters) != len(function.Parameters) {
		return &errors.ParameterCount{
			FunctionName:  function.Name,
			CountGiven:    len(parameters),
			CountRequired: len(function.Parameters),
		}
	}

	// print is a little special
	if isBuiltin && functionName == "print" {
		parameter := parameters[0]

		if parameter.Token.Kind != token.Text {
			return fmt.Errorf("'%s' requires a text parameter instead of '%s'", function.Name, parameter.Token.Text())
		}

		text := parameter.Token.Text() + "\n"
		address := state.assembler.AddString(text)
		state.assembler.MoveRegisterNumber(state.registers.Syscall[0], uint64(syscall.Write))
		state.assembler.MoveRegisterNumber(state.registers.Syscall[1], 1)
		state.assembler.MoveRegisterAddress(state.registers.Syscall[2], address)
		state.assembler.MoveRegisterNumber(state.registers.Syscall[3], uint64(len(text)))
		state.assembler.Syscall()
		return nil
	}

	// Call the function
	if functionName == BuiltinSyscall {
		pushRegisters, callRegisters, err := state.BeforeCall(function, parameters)

		if err != nil {
			return err
		}

		state.assembler.Syscall()
		state.AfterCall(function, pushRegisters, callRegisters)
	} else {
		pushRegisters, callRegisters, err := state.BeforeCall(function, parameters)

		if err != nil {
			return err
		}

		state.assembler.Call(functionName)
		state.AfterCall(function, pushRegisters, callRegisters)
	}

	// Mark return value register temporarily as used for better assembly output
	returnValueRegister := state.registers.ReturnValue[0]
	err := returnValueRegister.Use(expr)

	if err != nil {
		return err
	}

	// Save return value in temporary register
	if expr.Register != returnValueRegister {
		if expr.Register != nil {
			state.assembler.MoveRegisterRegister(expr.Register, returnValueRegister)
		}

		returnValueRegister.Free()
	}

	return nil
}

// Call handles function calls.
func (state *State) Call(tokens []token.Token) error {
	firstToken := tokens[0]

	if firstToken.Kind != token.Identifier {
		return errors.MissingFunctionName
	}

	lastToken := tokens[len(tokens)-1]

	if lastToken.Kind != token.GroupEnd {
		return &errors.MissingCharacter{Character: ")"}
	}

	return state.TokensToRegister(tokens, nil)
}

// BeforeCall pushes parameters into registers.
func (state *State) BeforeCall(function *Function, parameters []*expression.Expression) (register.List, register.List, error) {
	var pushRegisters []*register.Register

	// Wait for function compilation to finish
	function.Wait()

	// Determine the registers we need to save
	for registerName := range function.UsedRegisterNames() {
		callModifiedRegister := state.registers.All.ByName(registerName)

		if !callModifiedRegister.IsFree() {
			pushRegisters = append(pushRegisters, callModifiedRegister)
		}
	}

	// Save registers
	for _, reg := range pushRegisters {
		state.assembler.PushRegister(reg)
	}

	// Determine which registers to use for our parameters
	var callRegisters register.List

	if function.Name == BuiltinSyscall {
		callRegisters = state.registers.Syscall
	} else {
		callRegisters = state.registers.Call
	}

	// Move parameters into registers
	for i, parameter := range parameters {
		callRegister := callRegisters[i]

		// Check if we can skip the move entirely in case our
		// variable is already inside the correct register.
		if parameter.IsLeaf() && parameter.Token.Kind == token.Identifier {
			variable := state.scopes.Get(parameter.Token.Text())

			if variable != nil && variable.Register() == callRegister {
				state.UseVariable(variable)
				continue
			}
		}

		// If one of the call registers is already in use,
		// move the current user of the register to another one.
		if !callRegister.IsFree() {
			freeRegister := state.registers.General.FindFree()

			if freeRegister == nil {
				return pushRegisters, callRegisters, errors.ExceededMaxVariables
			}

			state.assembler.MoveRegisterRegister(freeRegister, callRegister)
			variable, isVariable := callRegister.User().(*Variable)

			if isVariable {
				_ = variable.SetRegister(freeRegister)
			} else {
				panic("This should never happen")
			}

			callRegister.Free()
		}

		_ = callRegister.Use(function.Parameters[i])

		// Save the parameter in the call register
		err := state.ExpressionToRegister(parameter, callRegister)

		if err != nil {
			return pushRegisters, callRegisters, err
		}
	}

	return pushRegisters, callRegisters, nil
}

// AfterCall restores saved registers from the stack.
func (state *State) AfterCall(function *Function, pushedRegisters []*register.Register, callRegisters []*register.Register) {
	atomic.AddInt32(&function.CallCount, 1)

	// Restore saved registers
	for i := len(pushedRegisters) - 1; i >= 0; i-- {
		state.assembler.PopRegister(pushedRegisters[i])
	}

	// Free the call registers
	for _, callRegister := range callRegisters {
		callRegister.Free()
	}
}
