package build

import (
	"fmt"
	"io/ioutil"

	"github.com/akyoto/q/token"
)

// File represents a single source file.
type File struct {
	path     string
	contents []byte
	tokens   []token.Token
	verbose  bool
}

// NewFile creates a new compiler for a single file.
func NewFile(inputFile string) *File {
	file := &File{
		path:   inputFile,
		tokens: *tokenBufferPool.Get().(*[]token.Token),
	}

	file.tokens = append(file.tokens, token.Token{
		Kind: token.NewLine,
	})

	return file
}

// Read reads the entire file contents.
func (file *File) Read() error {
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
		return fmt.Errorf("Unknown expression: %s", string(file.contents[processed:]))
	}

	return nil
}

// Scan scans the input file.
func (file *File) Scan(handleFunction func(*Function)) error {
	var (
		function   *Function
		groupLevel = 0
		blockLevel = 0
	)

	for index, t := range file.tokens {
		switch t.Kind {
		case token.Identifier:
			if function != nil {
				continue
			}

			functionName := t.String()

			if functionName == "func" || functionName == "fn" {
				return NewError("A function can not be named 'func' or 'fn'", file.path, file.tokens[:index+1])
			}

			if index+1 >= len(file.tokens) || file.tokens[index+1].Kind != token.GroupStart {
				return NewError("Missing opening bracket '(' after the function name", file.path, file.tokens[:index+2])
			}

			function = &Function{
				Name: functionName,
			}

			if file.verbose {
				fmt.Println("Function:", function.Name)
			}

		case token.BlockStart:
			if groupLevel > 0 {
				return NewError("Missing closing bracket ')'", file.path, file.tokens[:index+1])
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
			function.compiler = NewCompiler(file, function.TokenStart, function.TokenEnd)
			handleFunction(function)
			function = nil

		case token.GroupStart:
			groupLevel++

		case token.GroupEnd:
			groupLevel--

		case token.NewLine:
			// OK.

		default:
			if function == nil {
				return NewError("Only function definitions are allowed at the top level", file.path, file.tokens[:index+1])
			}
		}
	}

	return nil
}

// Tokens returns the complete list of tokens.
func (file *File) Tokens() []token.Token {
	return file.tokens
}

// Close frees up all resources.
func (file *File) Close() {
	file.tokens = file.tokens[:0]
	tokenBufferPool.Put(&file.tokens)
}
