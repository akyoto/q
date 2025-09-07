package core

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
)

// evaluateNumber converts a numeric expression to an SSA value.
func (f *Function) evaluateNumber(expr *expression.Expression) (ssa.Value, error) {
	number, err := toNumber(expr.Token, f.File)

	if err != nil {
		return nil, err
	}

	v := f.Append(&ssa.Int{
		Int:    number,
		Source: expr.Source(),
	})

	return v, nil
}