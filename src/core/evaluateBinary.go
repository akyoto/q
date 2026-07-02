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

	leftStructType, leftIsStruct := types.Unwrap(leftValue.Type()).(*types.Struct)
	rightStructType, rightIsStruct := types.Unwrap(rightValue.Type()).(*types.Struct)

	if leftIsStruct && rightIsStruct && leftStructType == types.String && rightStructType == types.String {
		switch expr.Token.Kind {
		case token.Add:
			return f.evaluateStringOp("concat", leftValue, rightValue, expr.Source())
		case token.Equal:
			return f.evaluateStringOp("equal", leftValue, rightValue, expr.Source())
		case token.NotEqual:
			equal, err := f.evaluateStringOp("equal", leftValue, rightValue, expr.Source())

			if err != nil {
				return nil, err
			}

			v := f.Append(&ssa.UnaryOp{
				Op:      token.Not,
				Operand: equal,
				Source:  expr.Source(),
			})

			return v, nil
		}
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