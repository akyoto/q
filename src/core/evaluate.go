package core

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// evaluate converts an expression to an SSA value.
func (f *Function) evaluate(expr *expression.Expression) (ssa.Value, error) {
	if expr.IsLeaf() {
		return f.evaluateLeaf(expr)
	}

	switch expr.Token.Kind {
	case token.Call:
		return f.evaluateCall(expr)

	case token.Dot:
		return f.evaluateDot(expr)

	case token.Array:
		panic("not implemented")
	}

	if expr.Token.Kind.IsUnaryOperator() {
		left := expr.Children[0]

		leftValue, err := f.evaluate(left)

		if err != nil {
			return nil, err
		}

		v := f.Append(&ssa.UnaryOp{
			Operand: leftValue,
			Op:      expr.Token.Kind,
			Source:  ssa.Source(expr.Source()),
		})

		return v, nil
	}

	left := expr.Children[0]
	right := expr.Children[1]

	leftValue, err := f.evaluate(left)

	if err != nil {
		return nil, err
	}

	rightValue, err := f.evaluate(right)

	if err != nil {
		return nil, err
	}

	v := f.Append(&ssa.BinaryOp{
		Left:   leftValue,
		Right:  rightValue,
		Op:     expr.Token.Kind,
		Source: ssa.Source(expr.Source()),
	})

	return v, nil
}