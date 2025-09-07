package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// evaluateLeaf converts a leaf expression to an SSA value.
func (f *Function) evaluateLeaf(expr *expression.Expression) (ssa.Value, error) {
	switch expr.Token.Kind {
	case token.Identifier:
		return f.evaluateIdentifier(expr)

	case token.Number, token.Rune:
		return f.evaluateNumber(expr)

	case token.String:
		return f.evaluateString(expr)
	}

	return nil, errors.New(InvalidExpression, f.File, expr.Token.Position)
}