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
		from, err := f.evaluateRight(index.Children[0])

		if err != nil {
			return nil, err
		}

		if !types.Is(from.Type(), types.AnyInt) {
			return nil, errors.New(&TypeMismatch{Encountered: from.Type().Name(), Expected: types.AnyInt.Name()}, f.File, index.Children[0].Source().StartPos)
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
			to, err := f.evaluateRight(index.Children[1])

			if err != nil {
				return nil, err
			}

			if !types.Is(to.Type(), types.AnyInt) {
				return nil, errors.New(&TypeMismatch{Encountered: to.Type().Name(), Expected: types.AnyInt.Name()}, f.File, index.Children[1].Source().StartPos)
			}

			slice := &ssa.Struct{
				Typ:       types.String,
				Arguments: []ssa.Value{addressValue, to},
				Source:    expr.Source(),
			}

			return slice, nil
		}

		from, err := f.evaluateRight(index.Children[0])

		if err != nil {
			return nil, err
		}

		if !types.Is(from.Type(), types.AnyInt) {
			return nil, errors.New(&TypeMismatch{Encountered: from.Type().Name(), Expected: types.AnyInt.Name()}, f.File, index.Children[0].Source().StartPos)
		}

		to, err := f.evaluateRight(index.Children[1])

		if err != nil {
			return nil, err
		}

		if !types.Is(to.Type(), types.AnyInt) {
			return nil, errors.New(&TypeMismatch{Encountered: to.Type().Name(), Expected: types.AnyInt.Name()}, f.File, index.Children[1].Source().StartPos)
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