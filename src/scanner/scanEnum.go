package scanner

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// scanEnum scans a namespaced block of constants.
func (s *scanner) scanEnum(file *fs.File, tokens token.List, i int) (int, error) {
	enumName := tokens[i].StringFrom(file.Bytes)
	i += 2

	if tokens[i].Kind != token.BlockStart {
		return i, errors.NewAt(MissingBlockStart, file, tokens[i].Position)
	}

	enum := types.NewEnum(file.Package, enumName, file)
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

				enum.AddMember(name, value)
			}

			if tokens[i].Kind == token.BlockEnd {
				s.items <- enum
				return i, nil
			}

			start = -1
		}

		i++
	}

	return i, errors.NewAt(MissingBlockEnd, file, tokens[i].Position)
}