package ast

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

func parseAssert(tokens token.List, file *fs.File) (Node, error) {
	if len(tokens) == 1 {
		return nil, errors.NewAt(MissingExpression, file, tokens[0].End())
	}

	condition := expression.Parse(tokens[1:])
	return &Assert{Condition: condition}, nil
}