package build

import (
	"fmt"
	"io/ioutil"

	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/log"
	"github.com/akyoto/q/build/token"
)

// File represents a single source file.
type File struct {
	contents      []byte
	tokens        []token.Token
	path          string
	functionCount int64
	verbose       bool
}

// NewFile creates a new compiler for a single file.
func NewFile(inputFile string) *File {
	file := &File{
		path:   inputFile,
		tokens: make([]token.Token, 1, 8192),
	}

	file.tokens[0].Kind = token.NewLine
	return file
}

// Tokenize converts the entire file contents to a list of tokens.
func (file *File) Tokenize() error {
	var (
		err       error
		processed int
	)

	file.contents, err = ioutil.ReadFile(file.path)

	if err != nil {
		return err
	}

	file.tokens, processed = token.Tokenize(file.contents, file.tokens)

	if processed != len(file.contents) {
		remaining := file.contents[processed:]
		until := len(remaining)

		for i, c := range remaining {
			if c == '\n' {
				until = i
				break
			}
		}

		return NewError(fmt.Errorf("Unknown expression: %s", string(remaining[:until])), file.path, file.tokens)
	}

	return nil
}

// Scan scans the input file.
func (file *File) Scan(functions chan<- *Function) error {
	var (
		function   *Function
		groupLevel = 0
		blockLevel = 0
		index      token.Position
		t          token.Token
	)

	for index, t = range file.tokens {
		switch t.Kind {
		case token.Identifier:
			if function != nil {
				continue
			}

			functionName := t.Text()

			if functionName == "func" || functionName == "fn" {
				return NewError(errors.InvalidFunctionName, file.path, file.tokens[:index+1])
			}

			if index+1 >= len(file.tokens) || file.tokens[index+1].Kind != token.GroupStart {
				return NewError(errors.ParameterOpeningBracket, file.path, file.tokens[:index+2])
			}

			function = &Function{
				Name:           functionName,
				File:           file,
				parameterStart: index + 2,
			}

			if file.verbose {
				log.Info.Println("Function:", function.Name)
			}

			file.functionCount++

		case token.BlockStart:
			if groupLevel > 0 {
				return NewError(&errors.MissingCharacter{Character: ")"}, file.path, file.tokens[:index+1])
			}

			blockLevel++

			if function.TokenStart != 0 {
				continue
			}

			function.TokenStart = index + 1

		case token.BlockEnd:
			blockLevel--

			if blockLevel != 0 {
				continue
			}

			function.TokenEnd = index
			functions <- function
			function = nil

		case token.GroupStart:
			groupLevel++

		case token.GroupEnd:
			groupLevel--

			if groupLevel != 0 {
				continue
			}

			if function.TokenStart != 0 {
				continue
			}

			if function.parameterStart < index {
				parameter := file.tokens[function.parameterStart:index]
				parameterName := parameter[0]

				function.Parameters = append(function.Parameters, Variable{
					Name: parameterName.Text(),
				})

				function.parameterStart = -1
			}

		case token.Separator:
			if function == nil || function.TokenStart != 0 || groupLevel != 1 {
				continue
			}

			if function.parameterStart < index {
				parameter := file.tokens[function.parameterStart:index]
				parameterName := parameter[0]

				function.Parameters = append(function.Parameters, Variable{
					Name: parameterName.Text(),
				})

				function.parameterStart = index + 1
			}

		case token.NewLine:
			// OK.

		default:
			if function == nil {
				return NewError(errors.TopLevel, file.path, file.tokens[:index+1])
			}
		}
	}

	return nil
}

// Tokens returns the complete list of tokens.
func (file *File) Tokens() []token.Token {
	return file.tokens
}

// Close frees up the memory.
func (file *File) Close() {
	file.tokens = nil
	file.contents = nil
}
