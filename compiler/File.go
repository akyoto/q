package compiler

import (
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/similarity"
	"github.com/akyoto/q/spec"
	"github.com/akyoto/q/token"
)

// File represents a single source file.
type File struct {
	token.Producer
	name      string
	assembler *asm.Assembler
	groups    []spec.Group
	blocks    []spec.Block
	functions map[string]*spec.Function
}

// NewFile creates a new compiler for a single file.
func NewFile(inputFile string, assembler *asm.Assembler) *File {
	file := &File{
		name:      inputFile,
		assembler: assembler,
		functions: map[string]*spec.Function{},
		Producer: token.Producer{
			Tokens: *tokenBufferPool.Get().(*[]token.Token),
		},
	}

	file.Tokens = append(file.Tokens, token.Token{
		Kind: token.NewLine,
		Text: "",
	})

	return file
}

// Compile compiles the input file.
func (file *File) Compile() error {
	fd, err := os.Open(file.name)

	if err != nil {
		return err
	}

	defer fd.Close()

	var (
		buffer      [16384]byte
		unprocessed = make([]byte, 0, len(buffer))
		final       []byte
	)

	for {
		n, err := fd.Read(buffer[:])

		if n > 0 {
			if len(unprocessed) > 0 {
				final = append(unprocessed, buffer[:n]...) // nolint:gocritic
				unprocessed = unprocessed[:0]
			} else {
				final = buffer[:n]
			}

			processedBytes, compilerError := token.Tokenize(final, file.handleToken)

			if compilerError != nil {
				return compilerError
			}

			if processedBytes < len(final) {
				unprocessed = append(unprocessed, final[processedBytes:]...)
			}
		}

		if err == nil {
			continue
		}

		if err == io.EOF {
			if len(unprocessed) > 0 {
				return file.Error(fmt.Sprintf("Unknown expression: %s", string(unprocessed)))
			}

			break
		}

		return err
	}

	return nil
}

// Error generates an error message at the current token position.
// The error message is clickable in popular editors and leads you
// directly to the faulty file at the given line and position.
func (file *File) Error(message string) error {
	lineCount := 0
	column := 1

	for _, oldToken := range file.Tokens {
		if oldToken.Kind == token.NewLine {
			lineCount++
			column = 1
		} else {
			column += len(oldToken.Text)
		}
	}

	return fmt.Errorf("%s:%d:%d: %s", file.name, lineCount, column, message)
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
	// fmt.Printf("%-10s | %s\n", t.Kind, string(t.Text))

	switch t.Kind {
	case token.Text:
		// text := string(token.Text[1 : len(token.Text)-1])
		// file.assembler.Println(text)

	case token.NewLine:
		if len(file.groups) > 0 {
			return file.Error("Missing closing bracket ')'")
		}

	case token.BlockStart:
		file.blocks = append(file.blocks, nil)

	case token.BlockEnd:
		if len(file.blocks) == 0 {
			return file.Error("Missing opening bracket '{'")
		}

		file.assembler.Return()
		file.blocks = file.blocks[:len(file.blocks)-1]

	case token.GroupStart:
		previous := file.PreviousToken(-1)

		if previous.Kind != token.Identifier {
			return file.Error("Expected function name before '('")
		}

		isDefinition := false

		if len(file.Tokens) >= 3 {
			whiteSpace := file.PreviousToken(-2)
			functionKeyword := file.PreviousToken(-3)

			if whiteSpace.Kind == token.WhiteSpace && functionKeyword.Kind == token.Keyword && functionKeyword.Text == "func" {
				isDefinition = true
			}
		}

		functionName := previous.Text

		if isDefinition {
			function := &spec.Function{
				Name: functionName,
			}

			file.functions[functionName] = function
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
			file.assembler.AddLabel(function.Name)

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
					return file.Error(fmt.Sprintf("'%s' requires a text parameter instead of '%s'", call.Function.Name, parameter.Text))
				}

				text := parameter.Text
				text = text[1 : len(text)-1]
				file.assembler.Println(text)
			}

			call.Reset()
			functionCallPool.Put(call)
		}

		file.groups = file.groups[:len(file.groups)-1]
	}

	file.Tokens = append(file.Tokens, t)
	return nil
}

// Close frees up all resources.
func (file *File) Close() {
	file.Tokens = file.Tokens[:0]
	tokenBufferPool.Put(&file.Tokens)
}
