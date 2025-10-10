package ast

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

func parseLoop(tokens token.List, file *fs.File) (Node, error) {
	if len(tokens) < 3 {
		return nil, errors.New(&InvalidInstruction{Instruction: tokens.String(file.Bytes)}, file, tokens[0].Position)
	}

	if tokens[1].Kind == token.Dot {
		control := &LoopControl{
			Expression: expression.Parse(tokens[2:]),
		}

		return control, nil
	}

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