package core

import (
	"unsafe"

	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
)

// evaluateCast converts a type cast to an SSA value.
func (f *Function) evaluateCast(expr *expression.Expression) (ssa.Value, error) {
	left := expr.Children[0]
	leftValue, err := f.evaluate(left)

	if err != nil {
		return nil, err
	}

	right := (*expression.TypeExpression)(unsafe.Pointer(expr.Children[1]))
	typ, err := TypeFromTokens(right.Tokens, f.File, f.Env)

	if err != nil {
		return nil, err
	}

	if leftValue.Type() == typ {
		return nil, errors.New(UnnecessaryCast, f.File, expr.Token.Position)
	}

	v := f.Append(&ssa.Copy{
		Value:  leftValue,
		Typ:    typ,
		Source: left.Source(),
	})

	return v, nil
}