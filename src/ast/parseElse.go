package ast

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

func parseElse(tokens token.List, file *fs.File, nodes AST) (Node, error) {
	_, _, body, err := block(tokens, file)

	if len(nodes) == 0 {
		return nil, errors.New(ExpectedIfBeforeElse, file, tokens[0].Position)
	}

	last := nodes[len(nodes)-1]
	ifNode, exists := last.(*If)

	if !exists {
		return nil, errors.New(ExpectedIfBeforeElse, file, tokens[0].Position)
	}

	ifNode.Else = body
	return nil, err
}