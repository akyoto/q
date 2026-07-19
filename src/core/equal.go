package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// equal returns the binary operation to compare the left with the right value.
func (f *Function) equal(left ssa.Value, right ssa.Value, source ssa.Source) (ssa.Value, error) {
	leftStructType, leftIsStruct := types.Unwrap(left.Type()).(*types.Struct)
	rightStructType, rightIsStruct := types.Unwrap(right.Type()).(*types.Struct)

	if leftIsStruct && rightIsStruct && leftStructType == types.String && rightStructType == types.String {
		return f.evaluateStringOp("equal", left, right, source)
	}

	if leftIsStruct || rightIsStruct {
		return nil, errors.New(InvalidStructOperation, f.File, source)
	}

	comparison := f.Append(&ssa.BinaryOp{
		Left:   left,
		Right:  right,
		Op:     token.Equal,
		Source: source,
	})

	return comparison, nil
}