package main

import (
	"fmt"
	"io"
	"os"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/token"
)

// File represents a single source file.
type File struct {
	token.Producer
	assembler *asm.Assembler
	fileName  string
	funcCalls []FunctionCall
}

// NewFile creates a new compiler for a single file.
func NewFile(inputFile string, assembler *asm.Assembler) *File {
	return &File{
		assembler: assembler,
		fileName:  inputFile,
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
func (compiler *File) Compile() error {
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

			processedBytes, compilerError := token.Tokenize(final, compiler.handleToken)

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

// Error generates an error message at the current token position.
// The error message is clickable in popular editors and leads you
// directly to the faulty file at the given line and position.
func (compiler *File) Error(message string) error {
	lineCount := 0
	column := 1

	for _, oldToken := range compiler.Tokens {
		if oldToken.Kind == token.StartOfLine {
			lineCount++
			column = 1
		} else {
			column += len(oldToken.Text)
		}
	}

	return fmt.Errorf("%s:%d:%d: %s", compiler.fileName, lineCount, column, message)
}
