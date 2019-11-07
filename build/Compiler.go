package build

import (
	"fmt"
	"sort"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/build/similarity"
	"github.com/akyoto/q/spec"
	"github.com/akyoto/q/token"
)

// Compiler handles the compilation of tokens.
type Compiler struct {
	cursor        int
	file          *File
	tokenStart    int
	tokenEnd      int
	tokens        []token.Token
	environment   *Environment
	groups        []spec.Group
	blocks        []spec.Block
	functionStack []*Function
	assembler     *asm.Assembler
}

// NewCompiler creates a new compiler.
func NewCompiler(file *File, tokenStart int, tokenEnd int) *Compiler {
	return &Compiler{
		file:       file,
		tokenStart: tokenStart,
		tokenEnd:   tokenEnd,
		tokens:     file.tokens[tokenStart:tokenEnd],
		assembler:  assemblerPool.Get().(*asm.Assembler),
	}
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
		functionCall.ProcessedTokens = compiler.cursor + 1
		compiler.groups = append(compiler.groups, functionCall)

	case token.GroupEnd:
		if len(compiler.groups) == 0 {
			return compiler.Error("Missing opening bracket '('")
		}

		call := compiler.groups[len(compiler.groups)-1].(*FunctionCall)

		// Add the last parameter
		if call.ProcessedTokens < compiler.cursor {
			call.Parameters = append(call.Parameters, compiler.tokens[call.ProcessedTokens:])
		}

		// print builtin
		if call.Function.Name == "print" {
			parameters := call.Parameters

			if len(parameters) < len(call.Function.Parameters) {
				return compiler.Error(fmt.Sprintf("Too few arguments in '%s' call", call.Function.Name))
			}

			if len(parameters) > len(call.Function.Parameters) {
				return compiler.Error(fmt.Sprintf("Too many arguments in '%s' call", call.Function.Name))
			}

			parameter := parameters[0][0]

			if parameter.Kind != token.Text {
				return compiler.Error(fmt.Sprintf("'%s' requires a text parameter instead of '%s'", call.Function.Name, parameter.String()))
			}

			text := parameter.String()
			compiler.assembler.Println(text)
		} else {
			compiler.assembler.Call(call.Function.Name)
		}

		call.Reset()
		functionCallPool.Put(call)

		compiler.groups = compiler.groups[:len(compiler.groups)-1]

	case token.BlockStart:
		compiler.blocks = append(compiler.blocks, nil)

	case token.BlockEnd:
		if len(compiler.blocks) == 0 {
			return compiler.Error("Missing opening bracket '{'")
		}

		compiler.blocks = compiler.blocks[:len(compiler.blocks)-1]

	case token.NewLine:
		if len(compiler.groups) > 0 {
			return compiler.Error("Missing closing bracket ')'")
		}
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
