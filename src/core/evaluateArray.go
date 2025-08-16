package core

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// evaluateArray converts a array indexing expression to an SSA value.
func (f *Function) evaluateArray(expr *expression.Expression) (ssa.Value, error) {

	address := expr.Children[0]
	addressValue, err := f.evaluate(address)

	if err != nil {
		return nil, err
	}

	var indexValue ssa.Value

	if len(expr.Children) > 1 {
		index := expr.Children[1]
		indexValue, err = f.evaluate(index)

		if err != nil {
			return nil, err
		}
	} else {
		indexValue = f.Append(&ssa.Int{Int: 0})
	}

	v := f.Append(&ssa.Load{
		Typ:     addressValue.Type().(*types.Pointer).To,
		Address: addressValue,
		Index:   indexValue,
		Source:  ssa.Source(expr.Source()),
	})

	return v, nil
}