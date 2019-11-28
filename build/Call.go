package build

import (
	"fmt"
	"sync/atomic"

	"github.com/akyoto/asm/syscall"
	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/register"
	"github.com/akyoto/q/build/token"
)

// Call represents a function call in the source code.
type Call struct {
	Function   *Function
	Parameters [][]token.Token
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

	functionName := firstToken.Text()
	function := state.environment.Functions[functionName]
	isBuiltin := false

	if function == nil {
		function = BuiltinFunctions[functionName]
		isBuiltin = true
	}

	if function == nil {
		return state.UnknownFunctionError(functionName)
	}

	call := Call{
		Function: function,
	}

	// Calling a function with side effects causes our function to have side effects
	if atomic.LoadInt32(&function.SideEffects) > 0 {
		atomic.AddInt32(&state.function.SideEffects, 1)
	}

	bracketPos := 1
	parameterStart := bracketPos + 1
	state.tokenCursor += bracketPos
	pos := parameterStart

	for pos < len(tokens) {
		t := tokens[pos]

		switch t.Kind {
		case token.Separator:
			if pos == parameterStart {
				return errors.MissingParameter
			}

			parameterTokens := tokens[parameterStart:pos]
			call.Parameters = append(call.Parameters, parameterTokens)
			parameterStart = pos + 1

		case token.GroupEnd:
			if pos == parameterStart {
				// Call with no parameters
				break
			}

			parameterTokens := tokens[parameterStart:pos]
			call.Parameters = append(call.Parameters, parameterTokens)
			parameterStart = pos + 1
		}

		state.tokenCursor++
		pos++
	}

	// Parameter check
	if !function.NoParameterCheck && len(call.Parameters) != len(call.Function.Parameters) {
		return &errors.ParameterCount{
			FunctionName:  call.Function.Name,
			CountGiven:    len(call.Parameters),
			CountRequired: len(call.Function.Parameters),
		}
	}

	if isBuiltin {
		switch functionName {
		case "print":
			parameter := call.Parameters[0][0]

			if parameter.Kind != token.Text {
				return fmt.Errorf("'%s' requires a text parameter instead of '%s'", call.Function.Name, parameter.Text())
			}

			text := parameter.Text() + "\n"
			address := state.assembler.AddString(text)
			state.assembler.MoveRegisterNumber(state.registers.Syscall[0], uint64(syscall.Write))
			state.assembler.MoveRegisterNumber(state.registers.Syscall[1], 1)
			state.assembler.MoveRegisterAddress(state.registers.Syscall[2], address)
			state.assembler.MoveRegisterNumber(state.registers.Syscall[3], uint64(len(text)))
			state.assembler.Syscall()

		case "syscall":
			err := state.BeforeCall(&call, state.registers.Syscall)

			if err != nil {
				return err
			}

			state.assembler.Syscall()
			state.AfterCall(&call)
		}

		return nil
	}

	err := state.BeforeCall(&call, state.registers.Call)

	if err != nil {
		return err
	}

	state.assembler.Call(call.Function.Name)
	state.AfterCall(&call)
	return nil
}

// BeforeCall pushes parameters into registers.
func (state *State) BeforeCall(call *Call, registers []*register.Register) error {
	for index, tokens := range call.Parameters {
		register := registers[index]
		err := state.TokensToRegister(tokens, register)

		if err != nil {
			return err
		}
	}

	return nil
}

// AfterCall restores saved registers from the stack.
func (state *State) AfterCall(call *Call) {
	call.Function.Used = true
}
