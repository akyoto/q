package core

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
)

// evaluateBinary converts a binary expression to an SSA value.
func (f *Function) evaluateBinary(expr *expression.Expression) (ssa.Value, error) {
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