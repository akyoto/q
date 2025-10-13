package scanner

import (
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
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

	for i < len(tokens) {
		switch tokens[i].Kind {
		case token.Identifier:
			if start == -1 {
				start = i
			}

		case token.NewLine, token.BlockEnd:
			if start != -1 {
				name := tokens[start].StringFrom(file.Bytes)

				if tokens[start+1].Kind != token.Assign {
					return i, errors.NewAt(MissingAssign, file, tokens[start+1].Position)
				}

				valueTokens := tokens[start+2 : i]

				if len(valueTokens) == 0 {
					return i, errors.NewAt(MissingExpression, file, tokens[start+1].End())
				}

				value := expression.Parse(valueTokens)

				if value.Token.Kind == token.Invalid {
					return i, errors.New(InvalidExpression, file, valueTokens)
				}

				s.constants <- &core.Constant{
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