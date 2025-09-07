package ast

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

func parseIf(tokens token.List, file *fs.File) (Node, error) {
	blockStart, _, body, err := block(tokens, file)
	condition := expression.Parse(tokens[1:blockStart])
	return &If{Condition: condition, Body: body}, err
}