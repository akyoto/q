package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
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

	_, leftIsStruct := leftValue.Type().(*types.Struct)
	_, rightIsStruct := rightValue.Type().(*types.Struct)

	if leftIsStruct || rightIsStruct {
		return nil, errors.New(InvalidStructOperation, f.File, expr.Token.Position)
	}

	v := f.Append(&ssa.BinaryOp{
		Left:   leftValue,
		Right:  rightValue,
		Op:     expr.Token.Kind,
		Source: ssa.Source(expr.Source()),
	})

	return v, nil
}