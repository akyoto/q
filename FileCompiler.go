package main

import (
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/similarity"
)

// FileCompiler is a single-threaded file compiler.
type FileCompiler struct {
	assembler *asm.Assembler
	fileName  string
	tokens    []Token
	calls     []FunctionCall
}

func NewFileCompiler(inputFile string, assembler *asm.Assembler) *FileCompiler {
	return &FileCompiler{
		assembler: assembler,
		fileName:  inputFile,
		tokens: []Token{
			{TokenStartOfLine, nil},
		},
	}
}

// Compile reads the input file.
func (compiler *FileCompiler) Compile() error {
	file, err := os.Open(compiler.fileName)

	if err != nil {
		return err
	}

	defer file.Close()

	var (
		buffer      [16384]byte
		unprocessed = make([]byte, 0, len(buffer))
		final       []byte
	)

	for {
		n, err := file.Read(buffer[:])

		if n > 0 {
			if len(unprocessed) > 0 {
				final = append(unprocessed, buffer[:n]...) // nolint:gocritic
				unprocessed = unprocessed[:0]
			} else {
				final = buffer[:n]
			}

			processedBytes, compilerError := compiler.processBuffer(final)

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
				return compiler.Error(fmt.Sprintf("Unknown expression: %s", string(unprocessed)))
			}

			break
		}

		return err
	}

	return nil
}

// processBuffer processes the partial read and returns how many bytes were processed.
// The remaining bytes will be prepended to the next call.
func (compiler *FileCompiler) processBuffer(buffer []byte) (int, error) {
	var (
		i              int
		c              byte
		processedBytes int
	)

	for i < len(buffer) {
		c = buffer[i]

		// Identifiers
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			processedBytes = i

			for (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
				i++

				if i >= len(buffer) {
					return processedBytes, nil
				}

				c = buffer[i]
			}

			token := Token{TokenIdentifier, buffer[processedBytes:i]}
			err := compiler.processToken(token)

			if err != nil {
				return 0, err
			}

			processedBytes = i + 1
			continue
		}

		// Texts
		if c == '"' {
			processedBytes = i

			for {
				i++

				if i >= len(buffer) {
					return processedBytes, nil
				}

				c = buffer[i]

				if c == '"' {
					i++
					break
				}
			}

			token := Token{TokenText, buffer[processedBytes:i]}
			err := compiler.processToken(token)

			if err != nil {
				return 0, err
			}

			processedBytes = i + 1
			continue
		}

		// Bracket start
		if c == '(' {
			token := Token{TokenBracketStart, buffer[i : i+1]}
			err := compiler.processToken(token)

			if err != nil {
				return 0, err
			}

			i++
			processedBytes = i
			continue
		}

		// Bracket end
		if c == ')' {
			token := Token{TokenBracketEnd, buffer[i : i+1]}
			err := compiler.processToken(token)

			if err != nil {
				return 0, err
			}

			i++
			processedBytes = i
			continue
		}

		// End of line
		if c == '\n' {
			token := Token{TokenStartOfLine, nil}
			err := compiler.processToken(token)

			if err != nil {
				return 0, err
			}

			i++
			processedBytes = i
			continue
		}

		// Whitespace
		if c == ' ' || c == '\t' {
			processedBytes = i

			for {
				i++

				if i >= len(buffer) {
					return processedBytes, nil
				}

				c = buffer[i]

				if c != ' ' && c != '\t' {
					i++
					break
				}
			}

			token := Token{TokenWhiteSpace, buffer[processedBytes:i]}
			err := compiler.processToken(token)

			if err != nil {
				return 0, err
			}

			processedBytes = i + 1
			continue
		}

		i++
	}

	return processedBytes, nil
}

// processToken processes a single token.
func (compiler *FileCompiler) processToken(token Token) error {
	switch token.Kind {
	// "Hello"
	case TokenText:
		// text := string(token.Text[1 : len(token.Text)-1])
		// compiler.assembler.Println(text)

	case TokenStartOfLine:
		if len(compiler.calls) > 0 {
			return compiler.Error("Missing closing bracket ')'")
		}

	// '('
	case TokenBracketStart:
		previous := compiler.previousToken()

		if previous.Kind != TokenIdentifier {
			return compiler.Error("Expected function name before '('")
		}

		compiler.calls = append(compiler.calls, FunctionCall{
			FunctionName:   string(previous.Text),
			ParameterStart: len(compiler.tokens) + 1,
		})

	// ')'
	case TokenBracketEnd:
		if len(compiler.calls) == 0 {
			return compiler.Error("Missing opening bracket '('")
		}

		call := compiler.calls[len(compiler.calls)-1]

		// Add the last parameter
		if call.ParameterStart < len(compiler.tokens) {
			call.Parameters = append(call.Parameters, compiler.tokens[call.ParameterStart:])
		}

		knownFunctions := []string{"print"}

		switch call.FunctionName {
		case "print":
			parameters := call.Parameters
			expectedParameters := 1

			if len(parameters) < expectedParameters {
				return compiler.Error(fmt.Sprintf("Too few arguments in '%s' call", call.FunctionName))
			}

			if len(parameters) > expectedParameters {
				return compiler.Error(fmt.Sprintf("Too many arguments in '%s' call", call.FunctionName))
			}

			parameter := parameters[0][0]

			if parameter.Kind != TokenText {
				return compiler.Error(fmt.Sprintf("'%s' requires a text parameter instead of '%s'", call.FunctionName, string(parameter.Text)))
			}

			text := parameter.Text
			text = text[1 : len(text)-1]
			compiler.assembler.Println(string(text))
			compiler.calls = compiler.calls[:len(compiler.calls)-1]

		default:
			// Suggest a function name based on the similarity to known functions
			sort.Slice(knownFunctions, func(a, b int) bool {
				aSimilarity := similarity.Default(call.FunctionName, knownFunctions[a])
				bSimilarity := similarity.Default(call.FunctionName, knownFunctions[b])
				return aSimilarity > bSimilarity
			})

			if similarity.Default(call.FunctionName, knownFunctions[0]) > 0.9 {
				return compiler.Error(fmt.Sprintf("Unknown function '%s', did you mean '%s'?", call.FunctionName, knownFunctions[0]))
			}

			return compiler.Error(fmt.Sprintf("Unknown function '%s'", call.FunctionName))
		}
	}

	compiler.tokens = append(compiler.tokens, token)
	return nil
}

// previousToken returns the last token.
func (compiler *FileCompiler) previousToken() Token {
	return compiler.tokens[len(compiler.tokens)-1]
}

// Error generates an error message at the current token position.
// The error message is clickable in popular editors and leads you
// directly to the faulty file at the given line and position.
func (compiler *FileCompiler) Error(message string) error {
	lineCount := 0
	column := 1

	for _, oldToken := range compiler.tokens {
		if oldToken.Kind == TokenStartOfLine {
			lineCount++
			column = 1
		} else {
			column += len(oldToken.Text)
		}
	}

	return fmt.Errorf("%s:%d:%d: %s", compiler.fileName, lineCount, column, message)
}
