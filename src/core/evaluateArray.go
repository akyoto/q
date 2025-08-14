package core

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
)

// evaluateArray converts a array indexing expression to an SSA value.
func (f *Function) evaluateArray(expr *expression.Expression) (ssa.Value, error) {
	address := expr.Children[0]
	index := expr.Children[1]
	addressValue, err := f.evaluate(address)

	if err != nil {
		return nil, err
	}

	indexValue, err := f.evaluate(index)

	if err != nil {
		return nil, err
	}

	v := f.Append(&ssa.Load{
		Address: addressValue,
		Index:   indexValue,
		Source:  ssa.Source(expr.Source()),
	})

	return v, nil
}