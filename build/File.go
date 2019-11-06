package build

import (
	"fmt"
	"io/ioutil"
	"sort"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/build/similarity"
	"github.com/akyoto/q/spec"
	"github.com/akyoto/q/token"
)

// File represents a single source file.
type File struct {
	token.Producer
	path      string
	contents  []byte
	build     *Build
	groups    []spec.Group
	blocks    []spec.Block
	functions []*spec.Function
}

// NewFile creates a new compiler for a single file.
func NewFile(inputFile string) *File {
	file := &File{
		path: inputFile,
		Producer: token.Producer{
			Tokens: *tokenBufferPool.Get().(*[]token.Token),
		},
	}

	file.Tokens = append(file.Tokens, token.Token{
		Kind: token.NewLine,
	})

	return file
}

// Compile compiles the input file.
func (file *File) Compile() error {
	var err error
	file.contents, err = ioutil.ReadFile(file.path)

	if err != nil {
		return err
	}

	processed, err := token.Tokenize(file.contents, file.handleToken)

	if err != nil {
		return err
	}

	if processed != len(file.contents) {
		return file.Error(fmt.Sprintf("Unknown expression: %s", string(file.contents[:processed])))
	}

	return nil
}

// Error generates an error message at the current token position.
// The error message is clickable in popular editors and leads you
// directly to the faulty file at the given line and position.
func (file *File) Error(message string) error {
	lineCount := 0
	lineStart := 0

	for _, oldToken := range file.Tokens {
		if oldToken.Kind == token.NewLine {
			lineCount++
			lineStart = oldToken.Position
		}
	}

	lastToken := file.PreviousToken(-1)
	column := lastToken.Position - lineStart
	return fmt.Errorf("%s:%d:%d: %s", file.path, lineCount, column, message)
}

// UnknownFunctionError produces an unknown function error
// and tries to guess which function the user was trying to type.
func (file *File) UnknownFunctionError(functionName string) error {
	knownFunctions := []string{"print"}

	// Suggest a function name based on the similarity to known functions
	sort.Slice(knownFunctions, func(a, b int) bool {
		aSimilarity := similarity.Default(functionName, knownFunctions[a])
		bSimilarity := similarity.Default(functionName, knownFunctions[b])
		return aSimilarity > bSimilarity
	})

	if similarity.Default(functionName, knownFunctions[0]) > 0.9 {
		return file.Error(fmt.Sprintf("Unknown function '%s', did you mean '%s'?", functionName, knownFunctions[0]))
	}

	return file.Error(fmt.Sprintf("Unknown function '%s'", functionName))
}

// handleToken processes a single token.
func (file *File) handleToken(t token.Token) error {
	switch t.Kind {
	case token.BlockStart:
		file.blocks = append(file.blocks, nil)

	case token.BlockEnd:
		if len(file.blocks) == 0 {
			return file.Error("Missing opening bracket '{'")
		}

		if file.build != nil {
			file.Assembler().Return()
		}

		file.blocks = file.blocks[:len(file.blocks)-1]

	case token.GroupStart:
		previous := file.PreviousToken(-1)

		if previous.Kind != token.Identifier {
			return file.Error("Expected function name before '('")
		}

		isDefinition := false

		if len(file.Tokens) >= 2 {
			functionKeyword := file.PreviousToken(-2)

			if functionKeyword.Kind == token.Keyword && functionKeyword.String() == "func" {
				isDefinition = true
			}
		}

		functionName := previous.String()

		if isDefinition {
			function := &spec.Function{
				Name:      functionName,
				Assembler: assemblerPool.Get().(*asm.Assembler),
			}

			if file.build != nil {
				file.build.functions.Store(functionName, function)
			}

			file.functions = append(file.functions, function)
			file.groups = append(file.groups, function)
		} else {
			function := spec.Functions[functionName]

			if function == nil {
				return file.UnknownFunctionError(functionName)
			}

			functionCall := functionCallPool.Get().(*FunctionCall)
			functionCall.Function = function
			functionCall.ProcessedTokens = len(file.Tokens) + 1
			file.groups = append(file.groups, functionCall)
		}

	case token.GroupEnd:
		if len(file.groups) == 0 {
			return file.Error("Missing opening bracket '('")
		}

		switch group := file.groups[len(file.groups)-1].(type) {
		case *spec.Function:
			function := group

			if file.build != nil {
				file.Assembler().AddLabel(function.Name)
			}

		case *FunctionCall:
			call := group

			// Add the last parameter
			if call.ProcessedTokens < len(file.Tokens) {
				call.Parameters = append(call.Parameters, file.Tokens[call.ProcessedTokens:])
			}

			// print builtin
			if call.Function.Name == "print" {
				parameters := call.Parameters

				if len(parameters) < len(call.Function.Parameters) {
					return file.Error(fmt.Sprintf("Too few arguments in '%s' call", call.Function.Name))
				}

				if len(parameters) > len(call.Function.Parameters) {
					return file.Error(fmt.Sprintf("Too many arguments in '%s' call", call.Function.Name))
				}

				parameter := parameters[0][0]

				if parameter.Kind != token.Text {
					return file.Error(fmt.Sprintf("'%s' requires a text parameter instead of '%s'", call.Function.Name, parameter.String()))
				}

				text := parameter.String()

				if file.build != nil {
					file.Assembler().Println(text)
				}
			}

			call.Reset()
			functionCallPool.Put(call)
		}

		file.groups = file.groups[:len(file.groups)-1]

	case token.NewLine:
		if len(file.groups) > 0 {
			return file.Error("Missing closing bracket ')'")
		}
	}

	file.Tokens = append(file.Tokens, t)
	return nil
}

// Assembler returns the current assembler.
func (file *File) Assembler() *asm.Assembler {
	return file.CurrentFunction().Assembler
}

// CurrentFunction returns the current function.
func (file *File) CurrentFunction() *spec.Function {
	return file.functions[len(file.functions)-1]
}

// Close frees up all resources.
func (file *File) Close() {
	file.Tokens = file.Tokens[:0]
	tokenBufferPool.Put(&file.Tokens)
}
