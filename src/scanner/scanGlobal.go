package scanner

import (
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

// scanGlobal scans a global.
func (s *scanner) scanGlobal(file *fs.File, tokens token.List, i int) (int, error) {
	i++

	if tokens[i].Kind != token.BlockStart {
		return i, errors.New(MissingBlockStart, file, tokens[i].Position)
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
				global := &core.Global{
					Name:   tokens[start].String(file.Bytes),
					Tokens: tokens[start:i],
					File:   file,
				}

				s.globals <- global
			}

			if tokens[i].Kind == token.BlockEnd {
				return i, nil
			}

			start = -1
		}

		i++
	}

	return i, errors.New(MissingBlockEnd, file, tokens[i].Position)
}