package scanner

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

// scanAssignment scans a single assignment in a const block.
func scanAssignment(tokens token.List, start int, i int, file *fs.File) (string, *expression.Expression, error) {
	name := tokens[start].StringFrom(file.Bytes)

	if tokens[start+1].Kind != token.Assign {
		return "", nil, errors.NewAt(MissingAssign, file, tokens[start+1].Position)
	}

	valueTokens := tokens[start+2 : i]

	if len(valueTokens) == 0 {
		return "", nil, errors.NewAt(MissingExpression, file, tokens[start+1].End())
	}

	value := expression.Parse(valueTokens)

	if value.Token.Kind == token.Invalid {
		return "", nil, errors.New(InvalidExpression, file, valueTokens)
	}

	return name, value, nil
}