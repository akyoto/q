package scanner

import (
	"os"

	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

// scanFile scans a single file.
func (s *scanner) scanFile(path string, pkg string) error {
	contents, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	tokens := token.Tokenize(contents)

	file := &fs.File{
		Path:    path,
		Package: pkg,
		Bytes:   contents,
		Tokens:  tokens,
	}

	s.files <- file

	for i := 0; i < len(tokens); i++ {
		switch tokens[i].Kind {
		case token.NewLine:
		case token.Comment:
		case token.Identifier:
			i, err = s.scanFunction(file, tokens, i)
		case token.Extern:
			i, err = s.scanExtern(file, tokens, i)
		case token.Import:
			i, err = s.scanImport(file, tokens, i)
		case token.EOF:
			return nil
		case token.Invalid:
			return errors.New(&InvalidCharacter{Character: tokens[i].String(file.Bytes)}, file, tokens[i].Position)
		default:
			return errors.New(&InvalidTopLevel{Instruction: tokens[i].String(file.Bytes)}, file, tokens[i].Position)
		}

		if err != nil {
			return err
		}
	}

	return nil
}