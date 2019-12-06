package build

import (
	"fmt"
	"io/ioutil"
	"strings"
	"sync/atomic"

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
	imports       map[string]*Import
	verbose       bool
}

// NewFile creates a new compiler for a single file.
func NewFile(inputFile string) *File {
	file := &File{
		path:    inputFile,
		tokens:  make([]token.Token, 1, 8192),
		imports: make(map[string]*Import),
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

		err = &errors.UnknownExpression{
			Expression: string(remaining[:until]),
		}

		return NewError(err, file.path, file.tokens, nil)
	}

	return nil
}

// Scan scans the input file.
func (file *File) Scan(imports chan<- *Import, functions chan<- *Function) error {
	var (
		function   *Function
		groupLevel                = 0
		blockLevel                = 0
		tokens                    = file.tokens
		index      token.Position = 0
		t          token.Token
	)

begin:
	for ; index < len(tokens); index++ {
		t = tokens[index]

		switch t.Kind {
		case token.Identifier:
			if function != nil {
				continue
			}

			functionName := t.Text()

			if functionName == "func" || functionName == "fn" {
				return NewError(errors.InvalidFunctionName, file.path, tokens[:index+1], function)
			}

			if index+1 >= len(tokens) || tokens[index+1].Kind != token.GroupStart {
				return NewError(errors.ParameterOpeningBracket, file.path, tokens[:index+2], function)
			}

			function = &Function{
				Name:           functionName,
				File:           file,
				Finished:       make(chan struct{}),
				parameterStart: index + 2,
			}

			if functionName == "main" {
				function.CallCount = 1
			}

			if file.verbose {
				log.Info.Println("Function:", function.Name)
			}

			file.functionCount++

		case token.BlockStart:
			if groupLevel > 0 {
				return NewError(&errors.MissingCharacter{Character: ")"}, file.path, tokens[:index+1], function)
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
				parameter := tokens[function.parameterStart:index]
				parameterName := parameter[0]

				function.Parameters = append(function.Parameters, &Parameter{
					Name: parameterName.Text(),
				})

				function.parameterStart = -1
			}

		case token.Separator:
			if function == nil || function.TokenStart != 0 || groupLevel != 1 {
				continue
			}

			if function.parameterStart < index {
				parameter := tokens[function.parameterStart:index]
				parameterName := parameter[0]

				function.Parameters = append(function.Parameters, &Parameter{
					Name: parameterName.Text(),
				})

				function.parameterStart = index + 1
			}

		case token.Keyword:
			if t.Text() == "import" {
				stdLib, err := stdLibPath()

				if err != nil {
					return err
				}

				position := index
				fullImportPath := strings.Builder{}
				fullImportPath.WriteString(stdLib)
				fullImportPath.WriteByte('/')
				importPath := strings.Builder{}
				index++

				for ; index < len(tokens); index++ {
					t = tokens[index]

					switch t.Kind {
					case token.Identifier:
						fullImportPath.WriteString(t.Text())
						importPath.WriteString(t.Text())

					case token.Operator:
						if t.Text() != "." {
							return NewError(&errors.InvalidCharacter{Character: t.Text()}, file.path, tokens[:index+1], function)
						}

						fullImportPath.WriteByte('/')
						importPath.WriteByte('.')

					case token.NewLine:
						imp := &Import{
							Path:     importPath.String(),
							FullPath: fullImportPath.String(),
							Position: position,
							Used:     0,
						}

						file.imports[imp.Path] = imp
						imports <- imp
						index++
						goto begin
					}
				}
			}

			if function == nil {
				return NewError(errors.TopLevel, file.path, tokens[:index+1], function)
			}

		case token.NewLine:
			// OK.

		case token.Comment:
			// OK.

		default:
			if function == nil {
				return NewError(errors.TopLevel, file.path, tokens[:index+1], function)
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
