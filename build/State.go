package build

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/akyoto/q/build/assembler"
	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/instruction"
	"github.com/akyoto/q/build/register"
	"github.com/akyoto/q/build/token"
	"github.com/akyoto/stringutils/similarity"
)

// State encapsulates a compiler's state.
// Every compilation requires a fresh state.
type State struct {
	instructions    []instruction.Instruction
	tokens          []token.Token
	assembler       *assembler.Assembler
	log             *log.Logger
	scopes          *ScopeStack
	registers       *register.Manager
	function        *Function
	environment     *Environment
	tokenCursor     token.Position
	instrCursor     instruction.Position
	loop            LoopState
	forLoop         ForState
	ignoreContracts bool
	useIncDec       bool
}

// CompileInstructions compiles all instructions.
func (state *State) CompileInstructions() error {
	identifiers := state.identifiersLastUse()
	lastKillPos := 0

	for index, instr := range state.instructions {
		state.killVariables(identifiers, lastKillPos, instr.Position)
		lastKillPos = instr.Position

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

	case instruction.Require:
		return state.Require(instr.Tokens)

	case instruction.Ensure:
		return state.Ensure(instr.Tokens)

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
		if labelBeforeComparison != "" {
			state.assembler.AddLabel(labelBeforeComparison)
		}

		switch expression[0].Kind {
		case token.Identifier:
			variableName := expression[0].Text()
			variable := state.scopes.Get(variableName)

			if variable == nil {
				return nil, &errors.UnknownVariable{VariableName: variableName}
			}

			state.UseVariable(variable)
			state.assembler.CompareRegisterRegister(register, variable.Register())

			return nil, nil

		case token.Number:
			numberString := expression[0].Text()
			number, err := state.ParseInt(numberString)

			if err != nil {
				return nil, err
			}

			state.assembler.CompareRegisterNumber(register, uint64(number))
			return nil, nil

		default:
			return nil, errors.InvalidExpression
		}
	}

	temporary := state.registers.General.FindFree()

	if temporary == nil {
		return nil, errors.ExceededMaxVariables
	}

	err := state.TokensToRegister(expression, temporary)

	if err != nil {
		return nil, err
	}

	state.assembler.AddLabel(labelBeforeComparison)
	state.assembler.CompareRegisterRegister(register, temporary)
	return temporary, nil
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

// UseVariable marks the variable as used and should always
// be called when the variable value is required.
func (state *State) UseVariable(variable *Variable) {
	variable.AliveUntil = state.instrCursor + 1
	variable.LastAssignUsed = true
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
		return &errors.UnknownFunction{
			FunctionName: functionName,
			CorrectName:  knownFunctions[0],
		}
	}

	return &errors.UnknownFunction{
		FunctionName: functionName,
	}
}
