package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// evaluateUnary converts a unary expression to an SSA value.
func (f *Function) evaluateUnary(expr *expression.Expression) (ssa.Value, error) {
	if !expr.Token.Kind.IsUnaryOperator() {
		return nil, errors.New(MissingOperand, f.File, expr.Token.End())
	}

	left := expr.Children[0]
	leftValue, err := f.evaluateRight(left)

	if err != nil {
		return nil, err
	}

	_, isStruct := leftValue.Type().(*types.Struct)

	if isStruct {
		return nil, errors.New(InvalidStructOperation, f.File, expr.Token.Position)
	}

	v := f.Append(&ssa.UnaryOp{
		Operand: leftValue,
		Op:      expr.Token.Kind,
		Source:  expr.Source(),
	})

	return v, nil
}