package build

import (
	"sort"
	"strconv"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/build/log"
	"github.com/akyoto/q/build/register"
	"github.com/akyoto/q/build/similarity"
	"github.com/akyoto/q/instruction"
	"github.com/akyoto/q/token"
)

// State encapsulates a compiler's state.
// Every compilation requires a fresh state.
type State struct {
	assembler    *asm.Assembler
	scopes       *ScopeStack
	registers    *register.Manager
	function     *Function
	environment  *Environment
	tokens       []token.Token
	instructions []instruction.Instruction
	cursor       token.Position
	verbose      bool
}

// CompileInstructions compiles all instructions.
func (state *State) CompileInstructions() error {
	for _, instr := range state.instructions {
		err := state.Instruction(instr)

		if err != nil {
			return err
		}
	}

	return nil
}

// Instruction generates machine code for the given instruction.
func (state *State) Instruction(instr instruction.Instruction) error {
	state.cursor = instr.Position

	switch instr.Kind {
	case instruction.Assignment:
		return state.Assignment(instr.Tokens)

	case instruction.Call:
		return state.Call(instr.Tokens)

	case instruction.Keyword:
		return state.Keyword(instr.Tokens)

	case instruction.Invalid:
		return state.Invalid(instr.Tokens)

	default:
		return nil
	}
}

// Assignment handles assignment instructions.
func (state *State) Assignment(tokens Expression) error {
	left := tokens[0]

	if left.Kind != token.Identifier {
		return state.Error("Expected variable on the left side of the assignment")
	}

	variableName := left.Text()
	variable := state.scopes.Get(variableName)

	if variable == nil {
		register := state.registers.FindFreeRegister()

		if register == nil {
			return state.Errorf("Exceeded maximum limit of %d variables", len(state.registers.Registers))
		}

		variable = &Variable{
			Name:     variableName,
			Position: state.cursor,
		}

		variable.BindRegister(register)
		state.scopes.Add(variable)
	}

	// Skip variable name and operator
	expressionStart := 2
	state.cursor += expressionStart
	expression := tokens[expressionStart:]

	return state.SaveExpressionInRegister(variable.Register, expression)
}

// Call handles function calls.
func (state *State) Call(tokens Expression) error {
	left := tokens[0]

	if left.Kind != token.Identifier {
		return state.Error("Expected function name before '('")
	}

	right := tokens[len(tokens)-1]

	if right.Kind != token.GroupEnd {
		return state.Error("Missing closing bracket ')'")
	}

	functionName := left.Text()
	function := state.environment.Functions[functionName]
	isBuiltin := false

	if function == nil {
		function = BuiltinFunctions[functionName]
		isBuiltin = true
	}

	if function == nil {
		return state.UnknownFunctionError(functionName)
	}

	call := FunctionCall{
		Function: function,
	}

	bracketPos := 1
	parameterStart := bracketPos + 1
	pos := parameterStart

	for pos < len(tokens) {
		t := tokens[pos]

		switch t.Kind {
		case token.Separator:
			if pos == parameterStart {
				state.cursor += pos
				return state.Error("Missing parameter")
			}

			call.Parameters = append(call.Parameters, tokens[parameterStart:pos])
			parameterStart = pos + 1

		case token.GroupEnd:
			if pos == parameterStart {
				// Call with no parameters
				break
			}

			call.Parameters = append(call.Parameters, tokens[parameterStart:pos])
			parameterStart = pos + 1
		}

		pos++
	}

	// Parameter check
	if !function.NoParameterCheck {
		if len(call.Parameters) < len(call.Function.Parameters) {
			return state.Errorf("Too few arguments in '%s' call", call.Function.Name)
		}

		if len(call.Parameters) > len(call.Function.Parameters) {
			return state.Errorf("Too many arguments in '%s' call", call.Function.Name)
		}
	}

	if isBuiltin {
		switch functionName {
		case "print":
			parameter := call.Parameters[0][0]

			if parameter.Kind != token.Text {
				return state.Errorf("'%s' requires a text parameter instead of '%s'", call.Function.Name, parameter.Text())
			}

			text := parameter.Text()
			state.assembler.Println(text)

		case "syscall":
			err := state.BeforeCall(&call, state.registers.SyscallRegisters)

			if err != nil {
				return err
			}

			state.assembler.Syscall()
			state.AfterCall(&call)
		}

		return nil
	}

	err := state.BeforeCall(&call, state.registers.SyscallRegisters)

	if err != nil {
		return err
	}

	state.assembler.Call(call.Function.Name)
	state.AfterCall(&call)
	return nil
}

// Keyword handles keywords.
func (state *State) Keyword(tokens Expression) error {
	return state.Error("Not implemented")
}

// Invalid handles invalid instructions.
func (state *State) Invalid(tokens Expression) error {
	openingBrackets := token.Count(tokens, token.GroupStart)
	closingBrackets := token.Count(tokens, token.GroupEnd)

	if openingBrackets < closingBrackets {
		return state.Error("Missing opening bracket '('")
	}

	if openingBrackets > closingBrackets {
		return state.Error("Missing closing bracket ')'")
	}

	return state.Error("Invalid instruction")
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

// SaveExpressionInRegister moves the result of an expression to the given register.
func (state *State) SaveExpressionInRegister(register *register.Register, expression Expression) error {
	singleToken := expression[0]

	switch singleToken.Kind {
	case token.Identifier:
		variableName := singleToken.Text()
		variable := state.scopes.Get(variableName)

		if variable == nil {
			return state.Errorf("Unknown variable %s", variableName)
		}

		variable.TimesUsed++

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
			return state.Errorf("Not a number: %s", numberAsString)
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

// Error generates an error message at the current token position.
func (state *State) Error(message string) error {
	return state.function.Error(state.cursor, message)
}

// Errorf generates a formatted error message at the current token position.
func (state *State) Errorf(message string, args ...interface{}) error {
	return state.function.Errorf(state.cursor, message, args...)
}

// UnknownFunctionError produces an unknown function error
// and tries to guess which function the user was trying to type.
func (state *State) UnknownFunctionError(functionName string) error {
	knownFunctions := make([]string, 0, len(state.environment.Functions)+len(BuiltinFunctions))

	for builtin := range BuiltinFunctions {
		knownFunctions = append(knownFunctions, builtin)
	}

	for function := range state.environment.Functions {
		knownFunctions = append(knownFunctions, function)
	}

	// Suggest a function name based on the similarity to known functions
	sort.Slice(knownFunctions, func(a, b int) bool {
		aSimilarity := similarity.Default(functionName, knownFunctions[a])
		bSimilarity := similarity.Default(functionName, knownFunctions[b])
		return aSimilarity > bSimilarity
	})

	if similarity.Default(functionName, knownFunctions[0]) > 0.9 {
		return state.Errorf("Unknown function '%s', did you mean '%s'?", functionName, knownFunctions[0])
	}

	return state.Errorf("Unknown function '%s'", functionName)
}
