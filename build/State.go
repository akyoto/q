package build

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/build/expression"
	"github.com/akyoto/q/build/instruction"
	"github.com/akyoto/q/build/log"
	"github.com/akyoto/q/build/register"
	"github.com/akyoto/q/build/token"
	"github.com/akyoto/stringutils/similarity"
)

// State encapsulates a compiler's state.
// Every compilation requires a fresh state.
type State struct {
	instructions []instruction.Instruction
	tokens       []token.Token
	assembler    *asm.Assembler
	scopes       *ScopeStack
	registers    *register.Manager
	function     *Function
	environment  *Environment
	tokenCursor  token.Position
	instrCursor  instruction.Position
	loop         LoopState
	verbose      bool
}

// CompileInstructions compiles all instructions.
func (state *State) CompileInstructions() error {
	for index, instr := range state.instructions {
		err := state.Instruction(instr, index)

		if err != nil {
			return err
		}
	}

	return nil
}

// Instruction generates machine code for the given instruction.
func (state *State) Instruction(instr instruction.Instruction, index instruction.Position) error {
	state.tokenCursor = instr.Position
	state.instrCursor = index

	switch instr.Kind {
	case instruction.Assignment:
		return state.Assignment(instr.Tokens)

	case instruction.Call:
		return state.Call(instr.Tokens)

	case instruction.LoopStart:
		return state.LoopStart()

	case instruction.LoopEnd:
		return state.LoopEnd()

	case instruction.Invalid:
		return state.Invalid(instr.Tokens)

	default:
		return nil
	}
}

// LoopStart handles the start of loops.
func (state *State) LoopStart() error {
	state.loop.counter++
	label := fmt.Sprintf("loop_%d", state.loop.counter)
	state.loop.labels = append(state.loop.labels, label)
	state.assembler.AddLabel(label)

	if state.verbose {
		log.Asm.Printf("%s:\n", label)
	}

	return nil
}

// LoopEnd handles the end of loops.
func (state *State) LoopEnd() error {
	label := state.loop.labels[len(state.loop.labels)-1]
	state.assembler.Jump(label)
	state.loop.labels = state.loop.labels[:len(state.loop.labels)-1]

	if state.verbose {
		log.Asm.Printf("jmp %s\n", label)
	}

	return nil
}

// Assignment handles assignment instructions.
func (state *State) Assignment(tokens []token.Token) error {
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
			Position: state.tokenCursor,
		}

		variable.BindRegister(register)
		state.scopes.Add(variable)
	}

	// Skip variable name and operator
	expressionStart := 2
	state.tokenCursor += expressionStart
	value := tokens[expressionStart:]
	return state.TokensToRegister(value, variable.Register)
}

// Call handles function calls.
func (state *State) Call(tokens []token.Token) error {
	firstToken := tokens[0]

	if firstToken.Kind != token.Identifier {
		return state.Error("Expected function name before '('")
	}

	lastToken := tokens[len(tokens)-1]

	if lastToken.Kind != token.GroupEnd {
		return state.Error("Missing closing bracket ')'")
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

	call := FunctionCall{
		Function: function,
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
				return state.Error("Missing parameter")
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

			if state.verbose {
				log.Asm.Printf("print \"%s\"", text)
			}

		case "syscall":
			err := state.BeforeCall(&call, state.registers.SyscallRegisters)

			if err != nil {
				return err
			}

			state.assembler.Syscall()

			if state.verbose {
				log.Asm.Println("syscall")
			}

			state.AfterCall(&call)
		}

		return nil
	}

	err := state.BeforeCall(&call, state.registers.SyscallRegisters)

	if err != nil {
		return err
	}

	state.assembler.Call(call.Function.Name)

	if state.verbose {
		log.Asm.Printf("call %s\n", call.Function.Name)
	}

	state.AfterCall(&call)
	return nil
}

// Keyword handles keywords.
func (state *State) Keyword(tokens []token.Token) error {
	return state.Error("Not implemented")
}

// Invalid handles invalid instructions.
func (state *State) Invalid(tokens []token.Token) error {
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
func (state *State) AfterCall(call *FunctionCall) {
	call.Function.Used = true
}

// TokenToRegister moves a token into a register.
// It only works with identifiers, numbers and texts.
func (state *State) TokenToRegister(singleToken token.Token, register *register.Register) error {
	switch singleToken.Kind {
	case token.Identifier:
		variableName := singleToken.Text()
		variable := state.scopes.Get(variableName)

		if variable == nil {
			return state.Errorf("Unknown variable %s", variableName)
		}

		variable.AliveUntil = state.instrCursor + 1

		// Moving a variable into its own register is pointless
		if variable.Register == register {
			return nil
		}

		state.assembler.MoveRegisterRegister(register.Name, variable.Register.Name)

		if state.verbose {
			log.Asm.Printf("mov %s, %s\n", register, variable.Register)
		}

	case token.Number:
		numberString := singleToken.Text()
		number, err := strconv.ParseInt(numberString, 10, 64)

		if err != nil {
			return state.Errorf("Not a number: %s", numberString)
		}

		state.assembler.MoveRegisterNumber(register.Name, uint64(number))

		if state.verbose {
			log.Asm.Printf("mov %s, %d\n", register, number)
		}

	case token.Text:
		address := state.assembler.Strings.Add(singleToken.Text())
		state.assembler.MoveRegisterAddress(register.Name, address)

		if state.verbose {
			log.Asm.Printf("mov %s, <%d>\n", register, address)
		}
	}

	return nil
}

// TokensToRegister moves the result of a token expression into the given register.
func (state *State) TokensToRegister(tokens []token.Token, register *register.Register) error {
	if len(tokens) == 1 {
		return state.TokenToRegister(tokens[0], register)
	}

	expr, err := expression.FromTokens(tokens)

	if err != nil {
		return state.Error(err.Error())
	}

	return state.ExpressionToRegister(expr, register)
}

// ExpressionToRegister moves the result of an expression into the given register.
func (state *State) ExpressionToRegister(expr *expression.Expression, register *register.Register) error {
	if expr.IsLeaf() {
		return state.TokenToRegister(expr.Value, register)
	}

	expr.SortByRegisterCount()
	tmp := expr

	for len(tmp.Children) > 0 {
		tmp.Children[0].Register = register
		tmp = tmp.Children[0]
	}

	err := expr.EachOperation(func(sub *expression.Expression) error {
		left := sub.Children[0]
		right := sub.Children[1]

		if left.Register == nil {
			left.Register = state.registers.FindFreeRegister()
		}

		sub.Register = left.Register

		if sub.Register.UsedBy == nil {
			sub.Register.UsedBy = sub
		}

		// Left operand
		if left.IsLeaf() {
			err := state.TokenToRegister(left.Value, sub.Register)

			if err != nil {
				return err
			}
		} else if sub.Register != left.Register {
			state.assembler.MoveRegisterRegister(sub.Register.Name, left.Register.Name)

			if state.verbose {
				log.Asm.Printf("mov %s, %s\n", sub.Register, left.Register)
			}
		}

		// Operator
		operator := sub.Value.Text()

		// Right operand is a leaf node
		if right.IsLeaf() {
			switch right.Value.Kind {
			case token.Identifier:
				variableName := right.Value.Text()
				variable := state.scopes.Get(variableName)

				if variable == nil {
					return state.Errorf("Unknown variable %s", variableName)
				}

				variable.AliveUntil = state.instrCursor + 1
				return state.CalculateRegisterRegister(operator, sub.Register, variable.Register)

			case token.Number:
				return state.CalculateRegisterNumber(operator, sub.Register, right.Value.Text())

			default:
				return state.Errorf("Invalid operand %s", right.Value)
			}
		}

		// Right operand is an expression
		err := state.CalculateRegisterRegister(operator, sub.Register, right.Register)

		if right.Register != nil {
			right.Register.UsedBy = nil
		}

		return err
	})

	if err != nil {
		return err
	}

	_ = expr.EachOperation(func(expr *expression.Expression) error {
		if expr.Register != register {
			expr.Register.UsedBy = nil
		}

		return nil
	})

	return nil
}

// CalculateRegisterNumber performs an operation on a register and a number.
func (state *State) CalculateRegisterNumber(operation string, register *register.Register, operand string) error {
	number, err := strconv.ParseInt(operand, 10, 64)

	if err != nil {
		return state.Errorf("Not a number: %s", operand)
	}

	switch operation {
	case "+":
		state.assembler.AddRegisterNumber(register.Name, uint64(number))

		if state.verbose {
			log.Asm.Printf("add %s, %d\n", register, number)
		}

	case "-":
		state.assembler.SubRegisterNumber(register.Name, uint64(number))

		if state.verbose {
			log.Asm.Printf("sub %s, %d\n", register, number)
		}

	case "*":
		state.assembler.MulRegisterNumber(register.Name, uint64(number))

		if state.verbose {
			log.Asm.Printf("imul %s, %d\n", register, number)
		}

	default:
		return state.Error("Not implemented")
	}

	return nil
}

// CalculateRegisterRegister performs an operation on two registers.
func (state *State) CalculateRegisterRegister(operation string, registerTo *register.Register, registerFrom *register.Register) error {
	switch operation {
	case "+":
		state.assembler.AddRegisterRegister(registerTo.Name, registerFrom.Name)

		if state.verbose {
			log.Asm.Printf("add %s, %s\n", registerTo, registerFrom)
		}

	case "-":
		state.assembler.SubRegisterRegister(registerTo.Name, registerFrom.Name)

		if state.verbose {
			log.Asm.Printf("sub %s, %s\n", registerTo, registerFrom)
		}

	case "*":
		state.assembler.MulRegisterRegister(registerTo.Name, registerFrom.Name)

		if state.verbose {
			log.Asm.Printf("imul %s, %s\n", registerTo, registerFrom)
		}

	default:
		return state.Error("Not implemented")
	}

	return nil
}

// Error generates an error message at the current token position.
func (state *State) Error(message string) error {
	return state.function.Error(state.tokenCursor, message)
}

// Errorf generates a formatted error message at the current token position.
func (state *State) Errorf(message string, args ...interface{}) error {
	return state.function.Errorf(state.tokenCursor, message, args...)
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
		aSimilarity := similarity.JaroWinkler(functionName, knownFunctions[a])
		bSimilarity := similarity.JaroWinkler(functionName, knownFunctions[b])
		return aSimilarity > bSimilarity
	})

	if similarity.JaroWinkler(functionName, knownFunctions[0]) > 0.9 {
		return state.Errorf("Unknown function '%s', did you mean '%s'?", functionName, knownFunctions[0])
	}

	return state.Errorf("Unknown function '%s'", functionName)
}
