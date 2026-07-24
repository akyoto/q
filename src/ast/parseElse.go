package ast

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

func parseElse(tokens token.List, file *fs.File, nodes AST) (Node, error) {
	if len(nodes) == 0 {
		return nil, errors.New(ExpectedIfBeforeElse, file, tokens[0])
	}

	blockStart, _, body, err := block(tokens, file)

	if err != nil {
		return nil, err
	}

	if blockStart != 1 {
		if tokens[1].Kind == token.If {
			return nil, errors.New(NoElseIf, file, tokens[:2])
		}

		return nil, errors.New(ExpectedBlock, file, tokens[1:blockStart])
	}

	last := nodes[len(nodes)-1]
	ifNode, exists := last.(*If)

	if !exists {
		return nil, errors.New(ExpectedIfBeforeElse, file, tokens[0])
	}

	ifNode.Else = body
	return nil, nil
}