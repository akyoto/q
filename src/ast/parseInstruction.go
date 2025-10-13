package ast

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/token"
)

// parseInstruction generates an AST node from an instruction.
func parseInstruction(tokens token.List, file *fs.File, nodes AST) (Node, error) {
	switch {
	case tokens[0].Kind.IsKeyword():
		return parseKeyword(tokens, file, nodes)

	case tokens[0].Kind == token.Comment:
		return nil, nil
	}

	expr := expression.Parse(tokens)

	switch {
	case expr.Token.Kind == token.Call:
		return &Call{Expression: expr}, nil

	case expr.Token.Kind == token.Define:
		if len(expr.Children) < 2 {
			return nil, errors.NewAt(MissingOperand, file, expr.Token.End())
		}

		return &Define{Expression: expr}, nil

	case expr.Token.Kind.IsAssignment():
		if len(expr.Children) < 2 {
			return nil, errors.NewAt(MissingOperand, file, expr.Token.End())
		}

		return &Assign{Expression: expr}, nil

	default:
		return nil, errors.New(InvalidInstruction, file, tokens)
	}
}