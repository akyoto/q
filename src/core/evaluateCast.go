package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// evaluateCast converts a type cast to an SSA value.
func (f *Function) evaluateCast(expr *expression.Expression) (ssa.Value, error) {
	left := expr.Children[0]
	leftValue, err := f.evaluateRight(left)

	if err != nil {
		return nil, err
	}

	right := expr.Children[1]
	typ, err := f.Env.TypeFromToken(right.Token, f.File)

	if err != nil {
		return nil, err
	}

	if leftValue.Type() == typ {
		return nil, errors.New(UnnecessaryCast, f.File, expr.Token)
	}

	fromType := types.Unwrap(leftValue.Type())
	toType := types.Unwrap(typ)

	if !types.IsCastable(fromType, toType) {
		return nil, errors.New(&TypeCastNotAllowed{From: fromType.Name(), To: toType.Name()}, f.File, right.Source())
	}

	v := f.Append(&ssa.Copy{
		Value:  leftValue,
		Typ:    typ,
		Source: left.Source(),
	})

	return v, nil
}