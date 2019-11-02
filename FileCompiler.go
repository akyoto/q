package main

import (
	"fmt"
	"io"
	"os"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/token"
)

// FileCompiler is a single-threaded file compiler.
type FileCompiler struct {
	assembler *asm.Assembler
	fileName  string
	tokens    []token.Token
	funcCalls []FunctionCall
}

func NewFileCompiler(inputFile string, assembler *asm.Assembler) *FileCompiler {
	return &FileCompiler{
		assembler: assembler,
		fileName:  inputFile,
		tokens: []token.Token{
			{
				Kind: token.StartOfLine,
				Text: nil,
			},
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

// previousToken returns the last token.
func (compiler *FileCompiler) previousToken() token.Token {
	return compiler.tokens[len(compiler.tokens)-1]
}

// Error generates an error message at the current token position.
// The error message is clickable in popular editors and leads you
// directly to the faulty file at the given line and position.
func (compiler *FileCompiler) Error(message string) error {
	lineCount := 0
	column := 1

	for _, oldToken := range compiler.tokens {
		if oldToken.Kind == token.StartOfLine {
			lineCount++
			column = 1
		} else {
			column += len(oldToken.Text)
		}
	}

	return fmt.Errorf("%s:%d:%d: %s", compiler.fileName, lineCount, column, message)
}
