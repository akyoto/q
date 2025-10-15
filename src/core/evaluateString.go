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

	length := f.Append(&ssa.Int{
		Int:    len(data),
		Source: expr.Source(),
	})

	pointer := f.Append(&ssa.Bytes{
		Bytes:  data,
		Source: expr.Source(),
	})

	v := f.makeStruct(types.String, expr.Source(), []ssa.Value{pointer, length})
	return v, nil
}