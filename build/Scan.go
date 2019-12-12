package build

import (
	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/log"
	"github.com/akyoto/q/build/token"
)

// Scan scans the input file.
func (file *File) Scan(imports chan<- *Import, functions chan<- *Function) error {
	var (
		tokens                  = file.tokens
		newlines                = 0
		index    token.Position = 0
		t        token.Token
	)

begin:
	for ; index < len(tokens); index++ {
		t = tokens[index]

		if t.Kind != token.NewLine {
			newlines = 0
		}

		switch t.Kind {
		case token.Identifier:
			var function *Function
			var err error
			function, index, err = file.scanFunction(tokens, index)

			if err != nil {
				return err
			}

			functions <- function

		case token.Keyword:
			if t.Text() == "import" {
				var imp *Import
				var err error

				imp, index, err = file.scanImport(tokens, index)

				if err != nil {
					return err
				}

				file.imports[imp.BaseName] = imp
				imports <- imp
				goto begin
			}

			if t.Text() == "struct" {
				var structure *Struct
				var err error

				structure, index, err = file.scanStruct(tokens, index)

				if err != nil {
					return err
				}

				log.Info.Println("struct", structure, structure.Fields, structure.Size)
				continue
			}

			return NewError(errors.TopLevel, file.path, tokens[:index+1], nil)

		case token.NewLine:
			newlines++

			if newlines == 3 {
				return NewError(errors.UnnecessaryNewlines, file.path, tokens[:index+1], nil)
			}

		case token.Comment:
			// OK.

		default:
			return NewError(errors.TopLevel, file.path, tokens[:index+1], nil)
		}
	}

	return nil
}
