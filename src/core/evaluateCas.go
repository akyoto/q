package core

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
)

// evaluateCas converts a CAS call to an SSA value.
func (f *Function) evaluateCas(expr *expression.Expression) (ssa.Value, error) {
	args, err := f.decompose(expr.Children[1:], nil, false)

	if err != nil {
		return nil, err
	}

	cas := f.Append(&ssa.Cas{
		Arguments: args,
		Source:    expr.Source(),
	})

	return cas, nil
}