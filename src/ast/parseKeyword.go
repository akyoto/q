package ast

import (
	"fmt"

	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

// parseKeyword generates a keyword node from an instruction.
func parseKeyword(tokens token.List, file *fs.File, nodes AST) (Node, error) {
	switch tokens[0].Kind {
	case token.Assert:
		if len(tokens) == 1 {
			return nil, errors.New(MissingExpression, file, tokens[0].End())
		}

		condition := expression.Parse(tokens[1:])
		return &Assert{Condition: condition}, nil

	case token.If:
		blockStart, _, body, err := block(tokens, file)
		condition := expression.Parse(tokens[1:blockStart])
		return &If{Condition: condition, Body: body}, err

	case token.Else:
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

	case token.Loop:
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

	case token.Return:
		if len(tokens) == 1 {
			return &Return{}, nil
		}

		values := expression.NewList(tokens[1:])
		return &Return{Values: values}, nil

	case token.Switch:
		blockStart := tokens.IndexKind(token.BlockStart)
		blockEnd := tokens.LastIndexKind(token.BlockEnd)

		if blockStart == -1 {
			return nil, errors.New(MissingBlockStart, file, tokens[0].End())
		}

		if blockEnd == -1 {
			return nil, errors.New(MissingBlockEnd, file, tokens[len(tokens)-1].End())
		}

		body := tokens[blockStart+1 : blockEnd]

		if len(body) == 0 {
			return nil, errors.New(EmptySwitch, file, tokens[0].Position)
		}

		cases, err := parseCases(body, file)
		return &Switch{Cases: cases}, err

	default:
		panic(fmt.Sprintf("keyword not implemented: %s", tokens[0].String(file.Bytes)))
	}
}