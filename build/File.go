package build

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/akyoto/q/token"
)

// File represents a single source file.
type File struct {
	path     string
	contents []byte
	tokens   []token.Token
	build    *Build
}

// NewFile creates a new compiler for a single file.
func NewFile(inputFile string, build *Build) *File {
	file := &File{
		path:   inputFile,
		tokens: *tokenBufferPool.Get().(*[]token.Token),
		build:  build,
	}

	file.tokens = append(file.tokens, token.Token{
		Kind: token.NewLine,
	})

	return file
}

// Scan scans the input file.
func (file *File) Scan() error {
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

	var (
		function   *Function
		blockLevel = 0
	)

	for index, t := range file.tokens {
		switch t.Kind {
		case token.Keyword:
			if t.String() != "func" {
				continue
			}

			if index+1 >= len(file.tokens) {
				return errors.New("Expected function name")
			}

			nameToken := file.tokens[index+1]

			if nameToken.Kind != token.Identifier {
				return errors.New("Function name must be a valid identifier")
			}

			function = &Function{
				File: file,
				Name: nameToken.String(),
			}

			if file.build.Verbose {
				fmt.Println("Function:", function.Name)
			}

		case token.BlockStart:
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
			file.build.functions.Store(function.Name, function)
		}
	}

	return nil
}

// Close frees up all resources.
func (file *File) Close() {
	file.tokens = file.tokens[:0]
	tokenBufferPool.Put(&file.tokens)
}
