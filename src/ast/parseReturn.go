package ast

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

func parseReturn(tokens token.List, file *fs.File) (Node, error) {
	if len(tokens) == 1 {
		return &Return{Token: tokens[0]}, nil
	}

	values := expression.NewList(tokens[1:])
	return &Return{Values: values, Token: tokens[0]}, nil
}