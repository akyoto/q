package build

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/build/log"
	"github.com/akyoto/q/build/register"
	"github.com/akyoto/q/build/similarity"
	"github.com/akyoto/q/spec"
	"github.com/akyoto/q/token"
)

// State encapsulates a compiler's state.
// Every compilation requires a fresh state.
type State struct {
	assembler   *asm.Assembler
	scopes      *ScopeStack
	registers   *register.Manager
	function    *Function
	environment *Environment
	tokens      []token.Token
	cursor      token.Position
	groups      []spec.Group
	blocks      []spec.Block
	verbose     bool
}

// ProcessTokens processes all tokens and fills the assembler with machine code.
func (state *State) ProcessTokens() error {
	for state.cursor < len(state.tokens) {
		err := state.Token(state.tokens[state.cursor])

		if err != nil {
			return err
		}

		state.cursor++
	}

	return nil
}

// Token processes a token.
func (state *State) Token(t token.Token) error {
	switch t.Kind {
	case token.GroupStart:
		return state.GroupStart()

	case token.GroupEnd:
		return state.GroupEnd()

	case token.Separator:
		return state.Separator()

	case token.BlockStart:
		return state.BlockStart()

	case token.BlockEnd:
		return state.BlockEnd()

	case token.Operator:
		return state.Operator(t)

	case token.NewLine:
		return state.NewLine()

	default:
		return nil
	}
}

// GroupStart processes a GroupStart token.
func (state *State) GroupStart() error {
	previous := state.TokenAtOffset(-1)

	if previous.Kind != token.Identifier {
		return state.Error("Expected function name before '('")
	}

	functionName := previous.Text()
	function := Functions[functionName]

	if function == nil && state.environment != nil {
		function = state.environment.functions[functionName]
	}

	if function == nil {
		return state.UnknownFunctionError(functionName)
	}

	functionCall := functionCallPool.Get().(*FunctionCall)
	functionCall.Function = function
	functionCall.ParameterStart = state.cursor + 1
	state.groups = append(state.groups, functionCall)
	return nil
}

// GroupEnd processes a GroupEnd token.
func (state *State) GroupEnd() error {
	if len(state.groups) == 0 {
		return state.Error("Missing opening bracket '('")
	}

	call := state.groups[len(state.groups)-1].(*FunctionCall)

	// Add the last parameter
	if call.ParameterStart < state.cursor {
		call.Parameters = append(call.Parameters, state.tokens[call.ParameterStart:state.cursor+1])
		call.ParameterStart = -1
	}

	parameters := call.Parameters

	// Builtin functions
	builtin := Functions[call.Function.Name]

	if builtin == nil || !builtin.NoParameterCheck {
		if len(parameters) < len(call.Function.Parameters) {
			return state.Error(fmt.Sprintf("Too few arguments in '%s' call", call.Function.Name))
		}

		if len(parameters) > len(call.Function.Parameters) {
			return state.Error(fmt.Sprintf("Too many arguments in '%s' call", call.Function.Name))
		}
	}

	if builtin != nil {
		switch builtin.Name {
		case "print":
			parameter := parameters[0][0]

			if parameter.Kind != token.Text {
				return state.Error(fmt.Sprintf("'%s' requires a text parameter instead of '%s'", call.Function.Name, parameter.Text()))
			}

			text := parameter.Text()
			state.assembler.Println(text)

		case "syscall":
			err := state.BeforeCall(call, state.registers.SyscallRegisters)

			if err != nil {
				return err
			}

			state.assembler.Syscall()
			state.AfterCall(call)
		}
	} else {
		err := state.BeforeCall(call, state.registers.SyscallRegisters)

		if err != nil {
			return err
		}

		state.assembler.Call(call.Function.Name)
		state.AfterCall(call)
	}

	call.Reset()
	functionCallPool.Put(call)

	state.groups = state.groups[:len(state.groups)-1]
	return nil
}

// Separator processes a Separator token.
func (state *State) Separator() error {
	if len(state.groups) == 0 {
		return state.Error("Invalid use of comma ',' without a function call")
	}

	call := state.groups[len(state.groups)-1].(*FunctionCall)

	// Add the parameter
	if call.ParameterStart < state.cursor {
		call.Parameters = append(call.Parameters, state.tokens[call.ParameterStart:state.cursor+1])
		call.ParameterStart = state.cursor + 1
	}

	return nil
}

// BlockStart processes a BlockStart token.
func (state *State) BlockStart() error {
	state.blocks = append(state.blocks, nil)
	state.scopes.Push()
	return nil
}

// BlockEnd processes a BlockEnd token.
func (state *State) BlockEnd() error {
	if len(state.blocks) == 0 {
		return state.Error("Missing opening bracket '{'")
	}

	state.blocks = state.blocks[:len(state.blocks)-1]
	state.scopes.Pop()
	return nil
}

// Operator processes a Operator token.
func (state *State) Operator(t token.Token) error {
	switch t.Text() {
	case "=":
		return state.OperatorAssign()

	default:
		return state.Error(fmt.Sprintf("Operator %s has not been implemented yet", t.Text()))
	}
}

// NewLine processes a NewLine token.
func (state *State) NewLine() error {
	if len(state.groups) > 0 {
		return state.Error("Missing closing bracket ')'")
	}

	return nil
}

// BeforeCall pushes parameters into registers.
func (state *State) BeforeCall(call *FunctionCall, registers []*register.Register) error {
	for index, expression := range call.Parameters {
		register := registers[index]
		err := state.SaveExpressionInRegister(register, expression)

		if err != nil {
			return err
		}
	}

	return nil
}

// AfterCall restores saved registers from the stack.
func (state *State) AfterCall(call *FunctionCall) {
	// Noop.
}

// OperatorAssign handles assignment instructions.
func (state *State) OperatorAssign() error {
	left := state.TokenAtOffset(-1)
	variableName := left.Text()
	variable := state.scopes.Get(variableName)

	if variable == nil {
		register := state.registers.FindFreeRegister()

		if register == nil {
			return state.Error(fmt.Sprintf("Exceeded maximum limit of %d variables", len(state.registers.Registers)))
		}

		variable = &Variable{
			Name:     variableName,
			Register: register,
		}

		register.UsedBy = variable
		state.scopes.Add(variable)
	}

	var expression Expression
	expressionStart := state.cursor + 1
	tokenIndex := expressionStart

	for {
		if tokenIndex >= len(state.tokens) {
			return state.Error("Invalid expression")
		}

		t := state.tokens[tokenIndex]

		if t.Kind == token.NewLine {
			expression = state.tokens[expressionStart:tokenIndex]
			break
		}

		tokenIndex++
	}

	return state.SaveExpressionInRegister(variable.Register, expression)
}

// SaveExpressionInRegister moves the result of an expression to the given register.
func (state *State) SaveExpressionInRegister(register *register.Register, expression Expression) error {
	singleToken := expression[0]

	switch singleToken.Kind {
	case token.Identifier:
		variableName := singleToken.Text()
		variable := state.scopes.Get(variableName)

		if variable == nil {
			return state.Error(fmt.Sprintf("Unknown variable %s", variableName))
		}

		if variable.Register != register {
			state.assembler.MoveRegisterRegister(register.Name, variable.Register.Name)

			if state.verbose {
				log.Info.Printf("mov %s, %s\n", register, variable.Register)
			}
		}

	case token.Number:
		numberAsString := singleToken.Text()
		number, err := strconv.ParseInt(numberAsString, 10, 64)

		if err != nil {
			return state.Error(fmt.Sprintf("Not a number: %s", numberAsString))
		}

		state.assembler.MoveRegisterNumber(register.Name, uint64(number))

		if state.verbose {
			log.Info.Printf("mov %s, %d\n", register, number)
		}

	case token.Text:
		address := state.assembler.Strings.Add(singleToken.Text())
		state.assembler.MoveRegisterAddress(register.Name, address)

		if state.verbose {
			log.Info.Printf("mov %s, <%d>\n", register, address)
		}

	default:
		return state.Error("Invalid expression")
	}

	return nil
}

// TokenAtOffset returns the token at the given offset relative to the cursor.
func (state *State) TokenAtOffset(offset token.Position) token.Token {
	return state.tokens[state.cursor+offset]
}

// Error generates an error message at the current token position.
// The error message is clickable in popular editors and leads you
// directly to the faulty file at the given line and position.
func (state *State) Error(message string) error {
	function := state.function
	until := function.TokenStart + state.cursor + 1
	return NewError(message, function.File.path, function.File.tokens[:until])
}

// UnknownFunctionError produces an unknown function error
// and tries to guess which function the user was trying to type.
func (state *State) UnknownFunctionError(functionName string) error {
	knownFunctions := []string{"print"}

	// Suggest a function name based on the similarity to known functions
	sort.Slice(knownFunctions, func(a, b int) bool {
		aSimilarity := similarity.Default(functionName, knownFunctions[a])
		bSimilarity := similarity.Default(functionName, knownFunctions[b])
		return aSimilarity > bSimilarity
	})

	if similarity.Default(functionName, knownFunctions[0]) > 0.9 {
		return state.Error(fmt.Sprintf("Unknown function '%s', did you mean '%s'?", functionName, knownFunctions[0]))
	}

	return state.Error(fmt.Sprintf("Unknown function '%s'", functionName))
}
