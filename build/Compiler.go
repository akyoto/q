package build

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/akyoto/asm"
	"github.com/akyoto/asm/syscall"
	"github.com/akyoto/q/build/similarity"
	"github.com/akyoto/q/spec"
	"github.com/akyoto/q/token"
)

// Compiler handles the compilation of tokens.
// Each function has its own compiler.
type Compiler struct {
	cursor          int
	file            *File
	tokenStart      int
	tokenEnd        int
	tokens          []token.Token
	environment     *Environment
	groups          []spec.Group
	blocks          []spec.Block
	functionStack   []*Function
	scopes          ScopeStack
	assembler       *asm.Assembler
	registerCounter int
}

// NewCompiler creates a new compiler.
func NewCompiler(file *File, tokenStart int, tokenEnd int) *Compiler {
	compiler := &Compiler{
		file:       file,
		tokenStart: tokenStart,
		tokenEnd:   tokenEnd,
		tokens:     file.tokens[tokenStart:tokenEnd],
		assembler:  assemblerPool.Get().(*asm.Assembler),
	}

	compiler.scopes.Push()
	return compiler
}

// Run compiles the set of tokens to machine code.
func (compiler *Compiler) Run() error {
	for cursor, t := range compiler.tokens {
		compiler.cursor = cursor
		err := compiler.handleToken(t)

		if err != nil {
			return err
		}
	}

	return nil
}

// handleToken processes a single token.
func (compiler *Compiler) handleToken(t token.Token) error {
	switch t.Kind {
	case token.GroupStart:
		previous := compiler.TokenAtOffset(-1)

		if previous.Kind != token.Identifier {
			return compiler.Error("Expected function name before '('")
		}

		functionName := previous.String()
		function := Functions[functionName]

		if function == nil && compiler.environment != nil {
			obj, exists := compiler.environment.functions.Load(functionName)

			if exists {
				function = obj.(*Function)
			}
		}

		if function == nil {
			return compiler.UnknownFunctionError(functionName)
		}

		functionCall := functionCallPool.Get().(*FunctionCall)
		functionCall.Function = function
		functionCall.ParameterStart = compiler.cursor + 1
		compiler.groups = append(compiler.groups, functionCall)

	case token.Separator:
		if len(compiler.groups) == 0 {
			return compiler.Error("Invalid use of comma ',' without a function call")
		}

		call := compiler.groups[len(compiler.groups)-1].(*FunctionCall)

		// Add the parameter
		if call.ParameterStart < compiler.cursor {
			call.Parameters = append(call.Parameters, compiler.tokens[call.ParameterStart:compiler.cursor+1])
			call.ParameterStart = compiler.cursor + 1
		}

	case token.GroupEnd:
		if len(compiler.groups) == 0 {
			return compiler.Error("Missing opening bracket '('")
		}

		call := compiler.groups[len(compiler.groups)-1].(*FunctionCall)

		// Add the last parameter
		if call.ParameterStart < compiler.cursor {
			call.Parameters = append(call.Parameters, compiler.tokens[call.ParameterStart:compiler.cursor+1])
			call.ParameterStart = -1
		}

		parameters := call.Parameters

		// Builtin functions
		builtin := Functions[call.Function.Name]

		if builtin == nil || !builtin.NoParameterCheck {
			if len(parameters) < len(call.Function.Parameters) {
				return compiler.Error(fmt.Sprintf("Too few arguments in '%s' call", call.Function.Name))
			}

			if len(parameters) > len(call.Function.Parameters) {
				return compiler.Error(fmt.Sprintf("Too many arguments in '%s' call", call.Function.Name))
			}
		}

		if builtin != nil {
			switch builtin.Name {
			case "print":
				parameter := parameters[0][0]

				if parameter.Kind != token.Text {
					return compiler.Error(fmt.Sprintf("'%s' requires a text parameter instead of '%s'", call.Function.Name, parameter.String()))
				}

				text := parameter.String()
				compiler.assembler.Println(text)

			case "syscall":
				err := compiler.beforeCall(call)

				if err != nil {
					return err
				}

				compiler.assembler.Syscall()
				compiler.afterCall(call)
			}
		} else {
			err := compiler.beforeCall(call)

			if err != nil {
				return err
			}

			compiler.assembler.Call(call.Function.Name)
			compiler.afterCall(call)
		}

		call.Reset()
		functionCallPool.Put(call)

		compiler.groups = compiler.groups[:len(compiler.groups)-1]

	case token.BlockStart:
		compiler.blocks = append(compiler.blocks, nil)
		compiler.scopes.Push()

	case token.BlockEnd:
		if len(compiler.blocks) == 0 {
			return compiler.Error("Missing opening bracket '{'")
		}

		compiler.blocks = compiler.blocks[:len(compiler.blocks)-1]
		compiler.scopes.Pop()

	case token.Operator:
		// Assignments
		if t.String() == "=" {
			err := compiler.handleAssignment()

			if err != nil {
				return err
			}
		}

	case token.NewLine:
		if len(compiler.groups) > 0 {
			return compiler.Error("Missing closing bracket ')'")
		}
	}

	return nil
}

// beforeCall pushes parameters into registers.
func (compiler *Compiler) beforeCall(call *FunctionCall) error {
	for index, expression := range call.Parameters {
		register := syscall.Registers[index]
		err := compiler.saveExpressionInRegister(register, expression)

		if err != nil {
			return err
		}
	}

	return nil
}

// afterCall restores saved registers from the stack.
func (compiler *Compiler) afterCall(call *FunctionCall) {
	// Currently a noop.
}

// handleAssignment handles assignment instructions.
func (compiler *Compiler) handleAssignment() error {
	left := compiler.TokenAtOffset(-1)
	variableName := left.String()
	variable := compiler.scopes.Get(variableName)

	if variable == nil {
		if compiler.registerCounter == len(variableRegisters) {
			return compiler.Error(fmt.Sprintf("Exceeded maximum limit of %d variables", len(variableRegisters)))
		}

		variable = &Variable{
			Name:     variableName,
			Register: variableRegisters[compiler.registerCounter],
		}

		compiler.scopes.Add(variable)
		compiler.registerCounter++
	}

	var expression Expression
	expressionStart := compiler.cursor + 1
	tokenIndex := expressionStart

	for {
		if tokenIndex >= len(compiler.tokens) {
			return compiler.Error("Invalid expression")
		}

		t := compiler.tokens[tokenIndex]

		if t.Kind == token.NewLine {
			expression = compiler.tokens[expressionStart:tokenIndex]
			break
		}

		tokenIndex++
	}

	return compiler.saveExpressionInRegister(variable.Register, expression)
}

// saveExpressionInRegister moves the result of an expression to the given register.
func (compiler *Compiler) saveExpressionInRegister(register string, expression Expression) error {
	singleToken := expression[0]

	switch singleToken.Kind {
	case token.Identifier:
		variableName := singleToken.String()
		variable := compiler.scopes.Get(variableName)

		if variable == nil {
			return compiler.Error(fmt.Sprintf("Unknown variable %s", variableName))
		}

		if variable.Register != register {
			compiler.assembler.MoveRegisterRegister(register, variable.Register)
		}

	case token.Number:
		numberAsString := singleToken.String()
		number, err := strconv.ParseInt(numberAsString, 10, 64)

		if err != nil {
			return compiler.Error(fmt.Sprintf("Not a number: %s", numberAsString))
		}

		compiler.assembler.MoveRegisterNumber(register, uint64(number))

	case token.Text:
		address := compiler.assembler.Strings.Add(singleToken.String())
		compiler.assembler.MoveRegisterAddress(register, address)

	default:
		return compiler.Error("Invalid expression")
	}

	return nil
}

// TokenAtOffset returns the token at the given offset relative to the cursor.
func (compiler *Compiler) TokenAtOffset(offset int) token.Token {
	return compiler.tokens[compiler.cursor+offset]
}

// Function returns the current function.
func (compiler *Compiler) Function() *Function {
	return compiler.functionStack[len(compiler.functionStack)-1]
}

// Error generates an error message at the current token position.
// The error message is clickable in popular editors and leads you
// directly to the faulty file at the given line and position.
func (compiler *Compiler) Error(message string) error {
	return NewError(message, compiler.file.path, compiler.file.tokens[:compiler.tokenStart+compiler.cursor+1])
}

// UnknownFunctionError produces an unknown function error
// and tries to guess which function the user was trying to type.
func (compiler *Compiler) UnknownFunctionError(functionName string) error {
	knownFunctions := []string{"print"}

	// Suggest a function name based on the similarity to known functions
	sort.Slice(knownFunctions, func(a, b int) bool {
		aSimilarity := similarity.Default(functionName, knownFunctions[a])
		bSimilarity := similarity.Default(functionName, knownFunctions[b])
		return aSimilarity > bSimilarity
	})

	if similarity.Default(functionName, knownFunctions[0]) > 0.9 {
		return compiler.Error(fmt.Sprintf("Unknown function '%s', did you mean '%s'?", functionName, knownFunctions[0]))
	}

	return compiler.Error(fmt.Sprintf("Unknown function '%s'", functionName))
}
