package core

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// evaluateString converts a string expression to an SSA value.
func (f *Function) evaluateString(expr *expression.Expression) (ssa.Value, error) {
	data := expr.Token.Bytes(f.File.Bytes)
	data = unescape(data)

	v := &ssa.Struct{
		Typ:    types.String,
		Source: expr.Source(),
	}

	length := f.Append(&ssa.Int{
		Int:       len(data),
		Structure: v,
		Source:    expr.Source(),
	})

	pointer := f.Append(&ssa.Bytes{
		Bytes:     data,
		Structure: v,
		Source:    expr.Source(),
	})

	v.Arguments = []ssa.Value{pointer, length}
	return v, nil
}