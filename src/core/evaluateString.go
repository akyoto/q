package core

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// evaluateString converts a string expression to an SSA value.
func (f *Function) evaluateString(expr *expression.Expression) (ssa.Value, error) {
	data, err := unescape(expr.Token, f.File)

	if err != nil {
		return nil, err
	}

	length := f.Append(&ssa.Int{
		Int:    len(data),
		Source: expr.Source(),
	})

	pointer := f.Append(&ssa.Bytes{
		Bytes:  data,
		Source: expr.Source(),
	})

	v := f.makeStruct(types.String, []ssa.Value{pointer, length}, expr.Source())
	return v, nil
}