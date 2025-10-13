package scanner

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// scanStruct scans a struct.
func (s *scanner) scanStruct(file *fs.File, tokens token.List, i int) (int, error) {
	structName := tokens[i].StringFrom(file.Bytes)
	structure := types.NewStruct(file, file.Package, structName)
	i++

	if tokens[i].Kind != token.BlockStart {
		return i, errors.NewAt(MissingBlockStart, file, tokens[i].Position)
	}

	i++
	start := -1

	for i < len(tokens) {
		switch tokens[i].Kind {
		case token.Identifier:
			if start == -1 {
				start = i
			}

		case token.NewLine, token.BlockEnd:
			if start != -1 {
				name := tokens[start].StringFrom(file.Bytes)

				structure.AddField(&types.Field{
					Name:   name,
					Tokens: tokens[start:i],
				})
			}

			if tokens[i].Kind == token.BlockEnd {
				s.structs <- structure
				return i, nil
			}

			start = -1
		}

		i++
	}

	return i, errors.NewAt(MissingBlockEnd, file, tokens[i].Position)
}