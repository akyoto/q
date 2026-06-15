package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/optimizer"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// evaluateBinary converts a binary expression to an SSA value.
func (f *Function) evaluateBinary(expr *expression.Expression) (ssa.Value, error) {
	left := expr.Children[0]
	right := expr.Children[1]

	if expr.Token.Kind.IsCommutative() {
		leftComplexity := optimizer.Complexity(left)
		rightComplexity := optimizer.Complexity(right)

		if rightComplexity > leftComplexity {
			left, right = right, left
		}
	}

	leftValue, err := f.evaluateRight(left)

	if err != nil {
		return nil, err
	}

	rightValue, err := f.evaluateRight(right)

	if err != nil {
		return nil, err
	}

	_, leftIsStruct := leftValue.Type().(*types.Struct)
	_, rightIsStruct := rightValue.Type().(*types.Struct)

	if leftIsStruct && rightIsStruct && expr.Token.Kind == token.Concat {
		return f.evaluateConcat(leftValue, rightValue, expr.Source())
	}

	if leftIsStruct || rightIsStruct {
		return nil, errors.New(InvalidStructOperation, f.File, expr.Token)
	}

	v := &ssa.BinaryOp{
		Left:   leftValue,
		Right:  rightValue,
		Op:     expr.Token.Kind,
		Source: expr.Source(),
	}

	if v.Op.IsComparison() {
		f.Block().Append(v)
		return v, nil
	}

	return f.Append(v), nil
}