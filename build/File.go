package build

import (
	"fmt"
	"io"
	"os"
	"sync/atomic"

	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/token"
	"github.com/akyoto/q/build/types"
)

// File represents a single source file.
type File struct {
	tokens        []token.Token
	imports       map[string]*Import
	environment   *Environment
	pkg           *Package
	path          string
	functionCount int64
	Error         error
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
		fd            *os.File
		err           error
		processed     uint16
		contents      []byte
		bufferOnStack [4096]byte
		buffer        = bufferOnStack[:]
	)

	fd, err = os.Open(file.path)

	if err != nil {
		return err
	}

	defer fd.Close()

	for {
		n, err := fd.Read(buffer)

		if n > 0 {
			contents = append(contents, buffer[:n]...)
		}

		if err != nil {
			if err == io.EOF {
				break
			}

			return err
		}
	}

	// Dividing the file length by 2 is a good approximation
	// of the number of tokens in the file.
	// For very small files we add a minimum of 4 reserved tokens.
	guessTokenCount := len(contents)/2 + 4
	tokens := make([]token.Token, 0, guessTokenCount)

	// Process tokens
	tokens, processed = token.Tokenize(contents, tokens)

	// Warn about missing final newline.
	if contents[len(contents)-1] != '\n' {
		return NewError(errors.New(errors.MissingEndingNewline), file.path, tokens, nil)
	}

	// If we didn't process everything, there's some error in the tokenization.
	if processed != uint16(len(contents)) {
		remaining := contents[processed:]
		until := len(remaining)

		for i, c := range remaining {
			if c == '\n' {
				until = i
				break
			}
		}

		err = errors.New(&errors.UnknownExpression{
			Expression: string(remaining[:until]),
		})

		return NewError(err, file.path, tokens, nil)
	}

	file.tokens = tokens
	return nil
}

// Tokens returns the complete list of tokens.
func (file *File) Tokens() []token.Token {
	return file.tokens
}

// Close frees up the memory.
func (file *File) Close() {
	for _, imp := range file.imports {
		if atomic.LoadInt32(&imp.Used) == 0 {
			file.Error = NewError(fmt.Errorf("Import '%s' has never been used", imp.Path), file.path, file.tokens[:imp.Position+1], nil)
			break
		}
	}

	file.tokens = nil
}

// Type returns the type with the given name.
func (file *File) Type(name string) *types.Type {
	t := file.environment.Types[name]

	if t == nil {
		prefix := file.pkg.Name + "."
		t = file.environment.Types[prefix+name]
	}

	return t
}
