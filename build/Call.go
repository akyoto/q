package build

import (
	"fmt"
	"strconv"
	"sync/atomic"

	"github.com/akyoto/asm/syscall"
	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/expression"
	"github.com/akyoto/q/build/register"
	"github.com/akyoto/q/build/token"
)

// Call handles function calls.
func (state *State) Call(tokens []token.Token) error {
	firstToken := tokens[0]

	if firstToken.Kind != token.Identifier {
		return errors.New(errors.MissingFunctionName)
	}

	lastToken := tokens[len(tokens)-1]

	if lastToken.Kind != token.GroupEnd {
		return errors.New(&errors.MissingCharacter{Character: ")"})
	}

	_, err := state.TokensToRegister(tokens, nil)
	return err
}

// CallExpression executes a function call.
func (state *State) CallExpression(expr *expression.Expression) error {
	parameters := expr.Children
	functionName := expr.Token.Text()
	functionName = PolymorphName(functionName, len(parameters))
	function := state.environment.Functions[functionName]
	isBuiltin := false

	if function == nil {
		function = BuiltinFunctions[functionName]
		isBuiltin = true
	}

	if function == nil {
		typ := state.environment.Types[functionName]

		if typ != nil {
			state.assembler.MoveRegisterNumber(state.registers.Syscall[0], 9)
			state.assembler.MoveRegisterNumber(state.registers.Syscall[1], 0)
			state.assembler.MoveRegisterNumber(state.registers.Syscall[2], uint64(typ.Size))
			state.assembler.MoveRegisterNumber(state.registers.Syscall[3], 3)
			state.assembler.MoveRegisterNumber(state.registers.Syscall[4], 290)
			state.assembler.Syscall()

			if expr.Register != state.registers.ReturnValue[0] {
				state.assembler.MoveRegisterRegister(expr.Register, state.registers.ReturnValue[0])
			}

			expr.Type = typ
			return nil
		}

		return errors.New(state.environment.UnknownFunctionError(functionName))
	}

	// Calling a function with side effects causes our function to have side effects
	if atomic.LoadInt32(&function.SideEffects) > 0 {
		atomic.AddInt32(&state.function.SideEffects, 1)
	}

	// Parameter check
	if !function.NoParameterCheck && len(parameters) != len(function.Parameters) {
		return errors.New(&errors.ParameterCount{
			FunctionName:  function.Name,
			CountGiven:    len(parameters),
			CountRequired: len(function.Parameters),
		})
	}

	if isBuiltin {
		switch functionName {
		case BuiltinPrint:
			parameter := parameters[0]

			if parameter.Token.Kind != token.Text {
				return fmt.Errorf("'%s' requires a text parameter instead of '%s'", function.Name, parameter.Token.Text())
			}

			state.printLn(parameter.Token.Text())
			return nil

		case BuiltinStore:
			variableName := parameters[0].Token.Text()
			offsetString := parameters[1].Token.Text()
			byteCountString := parameters[2].Token.Text()
			valueString := parameters[3].Token.Text()

			variable := state.scopes.Get(variableName)
			offset, _ := strconv.Atoi(offsetString)
			byteCount, _ := strconv.Atoi(byteCountString)
			value, _ := strconv.Atoi(valueString)

			state.UseVariable(variable)
			state.assembler.StoreNumber(variable.Register(), byte(offset), byte(byteCount), uint64(value))
			return nil
		}
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

		// Inline the function call if it's a little function
		if function.CanInline() {
			function.InlineInto(state.function)
		} else {
			state.assembler.Call(functionName)
		}

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

	if function.HasReturnValue() {
		expr.Type = function.ReturnTypes[0]
	}

	return nil
}

// BeforeCall pushes parameters into registers.
func (state *State) BeforeCall(function *Function, parameters []*expression.Expression) (register.List, register.List, error) {
	// nolint:prealloc
	var pushRegisters []*register.Register
	var usedRegisterIDs []register.ID

	if function == state.function {
		// Recursive call.
		// We can't determine the used registers for recursive calls
		// so we'll assume that every register has been used.
		// This is obviously bad for performance.
		// NOTE: We could save a recursive call reference here
		// and revisit it later after the function has been compiled.
		for _, reg := range state.registers.All {
			usedRegisterIDs = append(usedRegisterIDs, reg.ID)
		}
	} else {
		// Wait for function compilation to finish
		function.Wait()

		// If the function failed with a compilation error,
		// we're done here.
		if function.Error != nil {
			return nil, nil, function.Error
		}

		usedRegisterIDs = function.UsedRegisterIDs()
	}

	// Determine the registers we need to save
	for _, registerID := range usedRegisterIDs {
		callModifiedRegister := state.registers.ByID(registerID)

		if callModifiedRegister.IsFree() {
			continue
		}

		variable := callModifiedRegister.User().(*Variable)

		// Don't push variables that are going to die after this instruction
		if variable.AliveUntil < state.InstructionEndPosition() {
			continue
		}

		pushRegisters = append(pushRegisters, callModifiedRegister)
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
				return nil, nil, errors.New(errors.ExceededMaxVariables)
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
		typ, err := state.ExpressionToRegister(parameter, callRegister)

		if err != nil {
			return nil, nil, err
		}

		if !function.NoParameterCheck && typ != function.Parameters[i].Type {
			return nil, nil, errors.New(&errors.InvalidType{
				Type:          typ.String(),
				Expected:      function.Parameters[i].Type.String(),
				ParameterName: function.Parameters[i].Name,
			})
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

// printLn adds instructions to print a message to the console.
func (state *State) printLn(text string) {
	text += "\n"
	address := state.assembler.AddString(text)
	state.assembler.MoveRegisterNumber(state.registers.Syscall[0], uint64(syscall.Write))
	state.assembler.MoveRegisterNumber(state.registers.Syscall[1], 1)
	state.assembler.MoveRegisterAddress(state.registers.Syscall[2], address)
	state.assembler.MoveRegisterNumber(state.registers.Syscall[3], uint64(len(text)))
	state.assembler.Syscall()
}
