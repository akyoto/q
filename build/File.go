package build

import (
	"fmt"
	"io/ioutil"
	"sync/atomic"

	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/token"
)

// File represents a single source file.
type File struct {
	contents      []byte
	tokens        []token.Token
	path          string
	functionCount int64
	imports       map[string]*Import
	environment   *Environment
	verbose       bool
}

// NewFile creates a new compiler for a single file.
func NewFile(inputFile string) *File {
	file := &File{
		path:    inputFile,
		imports: make(map[string]*Import),
	}

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

	// Dividing the file length by 2 is a good approximation
	// of the number of tokens in the file.
	tokens := make([]token.Token, 0, len(file.contents)/2)

	// Process tokens
	file.tokens, processed = token.Tokenize(file.contents, tokens)

	// Warn about missing final newline.
	if file.contents[len(file.contents)-1] != '\n' {
		return NewError(errors.MissingEndingNewline, file.path, file.tokens, nil)
	}

	// If we didn't process everything, there's some error in the tokenization.
	if processed != len(file.contents) {
		remaining := file.contents[processed:]
		until := len(remaining)

		for i, c := range remaining {
			if c == '\n' {
				until = i
				break
			}
		}

		err = &errors.UnknownExpression{
			Expression: string(remaining[:until]),
		}

		return NewError(err, file.path, file.tokens, nil)
	}

	return nil
}

// Tokens returns the complete list of tokens.
func (file *File) Tokens() []token.Token {
	return file.tokens
}

// Close frees up the memory.
func (file *File) Close() error {
	for _, imp := range file.imports {
		if atomic.LoadInt32(&imp.Used) == 0 {
			return NewError(fmt.Errorf("Import '%s' has never been used", imp.Path), file.path, file.tokens[:imp.Position+1], nil)
		}
	}

	file.tokens = nil
	file.contents = nil
	return nil
}
