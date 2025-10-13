package scanner

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

// scanFile scans a single file.
func (s *scanner) scanFile(path string, pkg string) error {
	contents, err := fs.ReadFile(path)

	if err != nil {
		return err
	}

	tokens := token.Tokenize(contents)

	file := &fs.File{
		Path:    path,
		Package: pkg,
		Bytes:   contents,
		Tokens:  tokens,
		Imports: make(map[string]*fs.Import, 4),
	}

	s.files <- file

	for i := 0; i < len(tokens); i++ {
		switch tokens[i].Kind {
		case token.NewLine:
		case token.Comment:
		case token.Identifier:
			if i+1 >= len(tokens) {
				return errors.NewAt(InvalidFunctionDefinition, file, tokens[i].End())
			}

			next := tokens[i+1]

			switch next.Kind {
			case token.GroupStart:
				i, err = s.scanFunction(file, tokens, i)
			case token.BlockStart:
				i, err = s.scanStruct(file, tokens, i)
			case token.GroupEnd:
				return errors.NewAt(MissingGroupStart, file, next.Position)
			case token.BlockEnd:
				return errors.NewAt(MissingBlockStart, file, next.Position)
			case token.Invalid:
				return errors.New(&InvalidCharacter{Character: next.StringFrom(file.Bytes)}, file, next)
			default:
				return errors.NewAt(InvalidFunctionDefinition, file, next.Position)
			}
		case token.Const:
			i, err = s.scanConst(file, tokens, i)
		case token.Extern:
			i, err = s.scanExtern(file, tokens, i)
		case token.Import:
			i, err = s.scanImport(file, tokens, i)
		case token.Global:
			i, err = s.scanGlobal(file, tokens, i)
		case token.EOF:
			return nil
		case token.Invalid:
			return errors.New(&InvalidCharacter{Character: tokens[i].StringFrom(file.Bytes)}, file, tokens[i])
		case token.Script:
		default:
			return errors.New(InvalidTopLevel, file, tokens[i])
		}

		if err != nil {
			return err
		}
	}

	return nil
}