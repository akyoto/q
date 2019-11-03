package main

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
	groups    []Group
}

// NewFile creates a new compiler for a single file.
func NewFile(inputFile string, assembler *asm.Assembler) *File {
	return &File{
		name:      inputFile,
		assembler: assembler,
		Producer: token.Producer{
			Tokens: []token.Token{
				{
					Kind: token.StartOfLine,
					Text: nil,
				},
			},
		},
	}
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
		if oldToken.Kind == token.StartOfLine {
			lineCount++
			column = 1
		} else {
			column += len(oldToken.Text)
		}
	}

	return fmt.Errorf("%s:%d:%d: %s", file.name, lineCount, column, message)
}

// handleToken processes a single token.
func (file *File) handleToken(t token.Token) error {
	fmt.Println(t.Kind, string(t.Text))

	switch t.Kind {
	// "Hello"
	case token.Text:
		// text := string(token.Text[1 : len(token.Text)-1])
		// file.assembler.Println(text)

	case token.StartOfLine:
		if len(file.groups) > 0 {
			return file.Error("Missing closing bracket ')'")
		}

	// '('
	case token.GroupStart:
		previous := file.PreviousToken(-1)

		if previous.Kind != token.Identifier {
			return file.Error("Expected function name before '('")
		}

		isDefinition := false

		if len(file.Tokens) >= 3 {
			whiteSpace := file.PreviousToken(-2)
			functionKeyword := file.PreviousToken(-3)

			if whiteSpace.Kind == token.WhiteSpace && functionKeyword.Kind == token.Keyword && string(functionKeyword.Text) == "func" {
				isDefinition = true
			}
		}

		if isDefinition {

		} else {
			functionName := string(previous.Text)

			file.groups = append(file.groups, &FunctionCall{
				Function:        spec.Functions[functionName],
				ProcessedTokens: len(file.Tokens) + 1,
			})
		}

	// ')'
	case token.GroupEnd:
		if len(file.groups) == 0 {
			return file.Error("Missing opening bracket '('")
		}

		call := file.groups[len(file.groups)-1].(*FunctionCall)

		// Add the last parameter
		if call.ProcessedTokens < len(file.Tokens) {
			call.Parameters = append(call.Parameters, file.Tokens[call.ProcessedTokens:])
		}

		// Currently, we only allow builtin functions
		if spec.Functions[call.Function.Name] == nil {
			knownFunctions := []string{"print"}

			// Suggest a function name based on the similarity to known functions
			sort.Slice(knownFunctions, func(a, b int) bool {
				aSimilarity := similarity.Default(call.Function.Name, knownFunctions[a])
				bSimilarity := similarity.Default(call.Function.Name, knownFunctions[b])
				return aSimilarity > bSimilarity
			})

			if similarity.Default(call.Function.Name, knownFunctions[0]) > 0.9 {
				return file.Error(fmt.Sprintf("Unknown function '%s', did you mean '%s'?", call.Function.Name, knownFunctions[0]))
			}

			return file.Error(fmt.Sprintf("Unknown function '%s'", call.Function.Name))
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
				return file.Error(fmt.Sprintf("'%s' requires a text parameter instead of '%s'", call.Function.Name, string(parameter.Text)))
			}

			text := parameter.Text
			text = text[1 : len(text)-1]
			file.assembler.Println(string(text))
			file.groups = file.groups[:len(file.groups)-1]
		}
	}

	file.Tokens = append(file.Tokens, t)
	return nil
}
