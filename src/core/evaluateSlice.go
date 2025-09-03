package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// evaluateSlice converts a slice expression to an SSA value.
func (f *Function) evaluateSlice(expr *expression.Expression, index *expression.Expression, addressValue ssa.Value, length ssa.Value) (ssa.Value, error) {
	switch len(index.Children) {
	case 1:
		from, err := f.evaluate(index.Children[0])

		if err != nil {
			return nil, err
		}

		newPointer := f.Append(&ssa.BinaryOp{
			Op:    token.Add,
			Left:  addressValue,
			Right: from,
		})

		newLength := f.Append(&ssa.BinaryOp{
			Op:    token.Sub,
			Left:  length,
			Right: from,
		})

		slice := &ssa.Struct{
			Typ:       types.String,
			Arguments: []ssa.Value{newPointer, newLength},
			Source:    expr.Source(),
		}

		return slice, nil

	case 2:
		if index.Children[0].Token.Kind == token.Invalid {
			newLength, err := f.evaluate(index.Children[1])

			if err != nil {
				return nil, err
			}

			slice := &ssa.Struct{
				Typ:       types.String,
				Arguments: []ssa.Value{addressValue, newLength},
				Source:    expr.Source(),
			}

			return slice, nil
		}

		from, err := f.evaluate(index.Children[0])

		if err != nil {
			return nil, err
		}

		to, err := f.evaluate(index.Children[1])

		if err != nil {
			return nil, err
		}

		newPointer := f.Append(&ssa.BinaryOp{
			Op:    token.Add,
			Left:  addressValue,
			Right: from,
		})

		newLength := f.Append(&ssa.BinaryOp{
			Op:    token.Sub,
			Left:  to,
			Right: from,
		})

		slice := &ssa.Struct{
			Typ:       types.String,
			Arguments: []ssa.Value{newPointer, newLength},
			Source:    expr.Source(),
		}

		return slice, nil
	}

	return nil, errors.New(InvalidExpression, f.File, expr.Source().StartPos)
}