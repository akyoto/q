package ast

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

func parseLoop(tokens token.List, file *fs.File) (Node, error) {
	blockStart, _, body, err := block(tokens, file)
	headTokens := tokens[1:blockStart]

	loop := &Loop{
		Head: nil,
		Body: body,
	}

	if len(headTokens) > 0 {
		loop.Head = expression.Parse(headTokens)
	}

	return loop, err
}