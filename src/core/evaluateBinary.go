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
	leftValue, err := f.evaluate(left)

	if err != nil {
		return nil, err
	}

	right := expr.Children[1]
	rightValue, err := f.evaluate(right)

	if err != nil {
		return nil, err
	}

	_, leftIsStruct := leftValue.Type().(*types.Struct)
	_, rightIsStruct := rightValue.Type().(*types.Struct)

	if leftIsStruct || rightIsStruct {
		return nil, errors.New(InvalidStructOperation, f.File, expr.Token.Position)
	}

	v := &ssa.BinaryOp{
		Left:   leftValue,
		Right:  rightValue,
		Op:     expr.Token.Kind,
		Source: expr.Source(),
	}

	err = f.lintBinaryOp(v)

	if err != nil {
		return nil, err
	}

	if v.Op.IsComparison() {
		f.Block().Append(v)
		return v, nil
	}

	return f.Append(v), nil
}