package scanner

import (
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

// scanConst scans a block of constants.
func (s *scanner) scanConst(file *fs.File, tokens token.List, i int) (int, error) {
	i++

	if tokens[i].Kind != token.BlockStart {
		return i, errors.NewAt(MissingBlockStart, file, tokens[i].Position)
	}

	i++
	start := -1
	blockLevel := 1

	for i < len(tokens) {
		switch tokens[i].Kind {
		case token.Identifier:
			if start == -1 {
				start = i
			}

		case token.BlockStart:
			blockLevel++

		case token.NewLine, token.BlockEnd:
			if tokens[i].Kind == token.BlockEnd {
				blockLevel--

				if blockLevel > 0 {
					break
				}
			}

			if start != -1 {
				name, value, err := scanAssignment(tokens, start, i, file)

				if err != nil {
					return i, err
				}

				s.items <- &core.Constant{
					Name:  name,
					File:  file,
					Value: value,
				}
			}

			if tokens[i].Kind == token.BlockEnd {
				return i, nil
			}

			start = -1
		}

		i++
	}

	return i, errors.NewAt(MissingBlockEnd, file, tokens[i].Position)
}