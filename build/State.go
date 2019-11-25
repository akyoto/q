package build

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/build/errors"
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
	forLoop      ForState
	optimize     bool
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

	case instruction.IfStart:
		return state.IfStart(instr.Tokens)

	case instruction.IfEnd:
		return state.IfEnd()

	case instruction.ForStart:
		return state.ForStart(instr.Tokens)

	case instruction.ForEnd:
		return state.ForEnd()

	case instruction.LoopStart:
		return state.LoopStart()

	case instruction.LoopEnd:
		return state.LoopEnd()

	case instruction.Return:
		return state.Return(instr.Tokens)

	case instruction.Invalid:
		return state.Invalid(instr.Tokens)

	default:
		return nil
	}
}

// CompareExpression compares a register with the result of the expression.
// If the expression needs to be stored in a temporary register, it will return it.
func (state *State) CompareExpression(register *register.Register, expression []token.Token, labelBeforeComparison string) (*register.Register, error) {
	if len(expression) == 1 {
		state.assembler.AddLabel(labelBeforeComparison)

		if state.verbose {
			log.Asm.Printf("%s:\n", labelBeforeComparison)
		}

		switch expression[0].Kind {
		case token.Identifier:
			variableName := expression[0].Text()
			variable := state.scopes.Get(variableName)

			if variable == nil {
				return nil, &errors.UnknownVariable{VariableName: variableName}
			}

			variable.AliveUntil = state.instrCursor + 1
			state.assembler.CompareRegisterRegister(register.Name, variable.Register.Name)

			if state.verbose {
				log.Asm.Printf("cmp %s, %s\n", register, variable.Register)
			}

			return nil, nil

		case token.Number:
			numberString := expression[0].Text()
			number, err := state.ParseInt(numberString)

			if err != nil {
				return nil, err
			}

			state.assembler.CompareRegisterNumber(register.Name, uint64(number))

			if state.verbose {
				log.Asm.Printf("cmp %s, %d\n", register, uint64(number))
			}

			return nil, nil

		default:
			return nil, errors.InvalidExpression
		}
	}

	temporary := state.registers.FindFreeRegister()

	if temporary == nil {
		return nil, errors.ExceededMaxVariables
	}

	err := state.TokensToRegister(expression, temporary)

	if err != nil {
		return nil, err
	}

	state.assembler.AddLabel(labelBeforeComparison)

	if state.verbose {
		log.Asm.Printf("%s:\n", labelBeforeComparison)
	}

	state.assembler.CompareRegisterRegister(register.Name, temporary.Name)

	if state.verbose {
		log.Asm.Printf("cmp %s, %s\n", register, temporary)
	}

	return temporary, nil
}

// IfStart handles the start of if conditions.
func (state *State) IfStart(tokens []token.Token) error {
	state.Expect(token.Keyword)
	expression := tokens[1:]
	variableName := expression[0].Text()
	variable := state.scopes.Get(variableName)

	if variable == nil {
		return &errors.UnknownVariable{VariableName: variableName}
	}

	numberString := expression[len(expression)-1].Text()
	number, err := state.ParseInt(numberString)

	if err != nil {
		return err
	}

	label := "if_1_end"
	state.assembler.CompareRegisterNumber(variable.Register.Name, uint64(number))

	if state.verbose {
		log.Asm.Printf("cmp %s, %d\n", variable.Register.Name, number)
	}

	operator := expression[1].Text()

	switch operator {
	case ">=":
		state.assembler.JumpIfLess(label)

		if state.verbose {
			log.Asm.Printf("jl %s\n", label)
		}

	case ">":
		state.assembler.JumpIfLessOrEqual(label)

		if state.verbose {
			log.Asm.Printf("jle %s\n", label)
		}

	case "<=":
		state.assembler.JumpIfGreater(label)

		if state.verbose {
			log.Asm.Printf("jg %s\n", label)
		}

	case "<":
		state.assembler.JumpIfGreaterOrEqual(label)

		if state.verbose {
			log.Asm.Printf("jle %s\n", label)
		}

	case "==":
		state.assembler.JumpIfNotEqual(label)

		if state.verbose {
			log.Asm.Printf("jne %s\n", label)
		}

	case "!=":
		state.assembler.JumpIfEqual(label)

		if state.verbose {
			log.Asm.Printf("je %s\n", label)
		}
	}

	return nil
}

// IfEnd handles the end of if conditions.
func (state *State) IfEnd() error {
	label := "if_1_end"
	state.assembler.AddLabel(label)

	if state.verbose {
		log.Asm.Printf("%s:\n", label)
	}

	return nil
}

// ForStart handles the start of for loops.
func (state *State) ForStart(tokens []token.Token) error {
	state.Expect(token.Keyword)
	state.scopes.Push()
	expression := tokens[1:]

	rangePos := token.Index(expression, token.Range)

	if rangePos == -1 {
		return errors.MissingRange
	}

	operatorPos := token.Index(expression, token.Operator)
	var register *register.Register

	if operatorPos == -1 {
		register = state.registers.FindFreeRegister()

		if register == nil {
			return errors.ExceededMaxVariables
		}

		err := state.TokensToRegister(expression[:rangePos], register)

		if err != nil {
			return err
		}
	} else {
		assignment := expression[:rangePos]
		variable, err := state.AssignVariable(assignment)

		if err != nil {
			return err
		}

		register = variable.Register
	}

	state.forLoop.counter++

	labelStart := fmt.Sprintf("for_%d", state.forLoop.counter)
	labelEnd := fmt.Sprintf("for_%d_end", state.forLoop.counter)

	state.forLoop.labels = append(state.forLoop.labels, labelStart)
	state.forLoop.registers = append(state.forLoop.registers, register)

	upperLimit := expression[rangePos+1:]

	if len(upperLimit) == 0 {
		return errors.MissingRangeLimit
	}

	state.tokenCursor++
	temporary, err := state.CompareExpression(register, upperLimit, labelStart)

	if err != nil {
		return err
	}

	state.forLoop.temporaries = append(state.forLoop.temporaries, temporary)
	state.assembler.JumpIfEqual(labelEnd)

	if state.verbose {
		log.Asm.Printf("je %s\n", labelEnd)
	}

	return nil
}

// ForEnd handles the end of for loops.
func (state *State) ForEnd() error {
	err := state.PopScope()

	if err != nil {
		return err
	}

	label := state.forLoop.labels[len(state.forLoop.labels)-1]
	register := state.forLoop.registers[len(state.forLoop.registers)-1]
	temporary := state.forLoop.temporaries[len(state.forLoop.temporaries)-1]

	state.forLoop.labels = state.forLoop.labels[:len(state.forLoop.labels)-1]
	state.forLoop.registers = state.forLoop.registers[:len(state.forLoop.registers)-1]
	state.forLoop.temporaries = state.forLoop.temporaries[:len(state.forLoop.temporaries)-1]

	state.assembler.IncreaseRegister(register.Name)
	state.assembler.Jump(label)
	state.assembler.AddLabel(label + "_end")

	if state.verbose {
		log.Asm.Printf("inc %s\n", register)
		log.Asm.Printf("jmp %s\n", label)
		log.Asm.Printf("%s:\n", label+"_end")
	}

	register.UsedBy = nil

	if temporary != nil {
		temporary.UsedBy = nil
	}

	return nil
}

// LoopStart handles the start of loops.
func (state *State) LoopStart() error {
	state.scopes.Push()
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
	state.scopes.Pop()
	label := state.loop.labels[len(state.loop.labels)-1]
	state.assembler.Jump(label)
	state.loop.labels = state.loop.labels[:len(state.loop.labels)-1]

	if state.verbose {
		log.Asm.Printf("jmp %s\n", label)
	}

	return nil
}

// PopScope pops the last scope on the stack and returns
// an error if there were any unused variables.
func (state *State) PopScope() error {
	for _, variable := range state.scopes.Unused() {
		return state.function.Error(variable.Position, &errors.UnusedVariable{VariableName: variable.Name})
	}

	state.scopes.Pop()
	return nil
}

// Return handles return statements.
func (state *State) Return(tokens []token.Token) error {
	state.assembler.Return()

	if state.verbose {
		log.Asm.Print("ret")
	}

	return nil
}

// Assignment handles assignment instructions.
func (state *State) Assignment(tokens []token.Token) error {
	_, err := state.AssignVariable(tokens)
	return err
}

// AssignVariable handles assignment instructions and also returns the referenced variable.
func (state *State) AssignVariable(tokens []token.Token) (*Variable, error) {
	cursor := 0
	mutable := false
	left := tokens[cursor]

	if left.Kind == token.Keyword && left.Text() == "mut" {
		mutable = true
		cursor++
		state.tokenCursor++
		left = tokens[cursor]
	}

	if left.Kind != token.Identifier {
		return nil, errors.ExpectedVariable
	}

	variableName := left.Text()
	variable := state.scopes.Get(variableName)

	if variable == nil {
		register := state.registers.FindFreeRegister()

		if register == nil {
			return nil, errors.ExceededMaxVariables
		}

		variable = &Variable{
			Name:     variableName,
			Position: state.tokenCursor,
			Mutable:  mutable,
		}

		variable.BindRegister(register)
		state.scopes.Add(variable)
	} else if !variable.Mutable {
		return variable, &errors.ImmutableVariable{VariableName: variable.Name}
	}

	// Operator
	cursor++
	state.tokenCursor++
	operator := tokens[cursor]

	if operator.Kind != token.Operator {
		return variable, errors.MissingAssignmentOperator
	}

	// Expression
	cursor++
	state.tokenCursor++
	expression := tokens[cursor:]

	if len(expression) == 0 {
		return variable, errors.MissingAssignmentExpression
	}

	err := state.TokensToRegister(expression, variable.Register)

	if err != nil {
		return variable, err
	}

	state.tokenCursor += len(expression)
	return variable, nil
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

	call := FunctionCall{
		Function: function,
	}

	// Calling a function with side effects causes our function to have side effects
	if function.HasSideEffects {
		state.function.HasSideEffects = true
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

			text := parameter.Text()
			state.assembler.Println(text)

			if state.verbose {
				log.Asm.Printf("mov %s, 1", state.registers.SyscallRegisters[0])
				log.Asm.Printf("mov %s, 1", state.registers.SyscallRegisters[1])
				log.Asm.Printf("mov %s, \"%s\"", state.registers.SyscallRegisters[2], text)
				log.Asm.Printf("mov %s, %d", state.registers.SyscallRegisters[3], len(text)+1)
				log.Asm.Println("syscall")
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

// Invalid handles invalid instructions.
func (state *State) Invalid(tokens []token.Token) error {
	openingBrackets := token.Count(tokens, token.GroupStart)
	closingBrackets := token.Count(tokens, token.GroupEnd)

	if openingBrackets < closingBrackets {
		return &errors.MissingCharacter{Character: "("}
	}

	if openingBrackets > closingBrackets {
		return &errors.MissingCharacter{Character: ")"}
	}

	return errors.InvalidInstruction
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
			return fmt.Errorf("Unknown variable %s", variableName)
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
		number, err := state.ParseInt(numberString)

		if err != nil {
			return err
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
		return err
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
					return fmt.Errorf("Unknown variable %s", variableName)
				}

				variable.AliveUntil = state.instrCursor + 1
				return state.CalculateRegisterRegister(operator, sub.Register, variable.Register)

			case token.Number:
				return state.CalculateRegisterNumber(operator, sub.Register, right.Value.Text())

			default:
				return fmt.Errorf("Invalid operand %s", right.Value)
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
	number, err := state.ParseInt(operand)

	if err != nil {
		return err
	}

	switch operation {
	case "+":
		if number == 1 && state.optimize {
			state.assembler.IncreaseRegister(register.Name)

			if state.verbose {
				log.Asm.Printf("inc %s\n", register)
			}

			return nil
		}

		state.assembler.AddRegisterNumber(register.Name, uint64(number))

		if state.verbose {
			log.Asm.Printf("add %s, %d\n", register, number)
		}

	case "-":
		if number == 1 && state.optimize {
			state.assembler.DecreaseRegister(register.Name)

			if state.verbose {
				log.Asm.Printf("dec %s\n", register)
			}

			return nil
		}

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
		return errors.NotImplemented
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
		return errors.NotImplemented
	}

	return nil
}

// ParseInt parses an integer number.
func (state *State) ParseInt(numberString string) (int64, error) {
	number, err := strconv.ParseInt(numberString, 10, 64)

	if err != nil {
		return 0, &errors.NotANumber{
			Expression: numberString,
		}
	}

	return number, nil
}

// Expect asserts that the token at the current cursor position has the given kind.
// If the comparison was successful, it will increment the cursor and return the token.
// If the expectation is not met, it will panic.
func (state *State) Expect(expectedKind token.Kind) token.Token {
	actual := state.tokens[state.tokenCursor]

	if actual.Kind != expectedKind {
		panic(fmt.Errorf("Expected '%s' instead of '%s'", expectedKind, actual))
	}

	state.tokenCursor++
	return actual
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
		return fmt.Errorf("Unknown function '%s', did you mean '%s'?", functionName, knownFunctions[0])
	}

	return fmt.Errorf("Unknown function '%s'", functionName)
}
