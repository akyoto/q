package core

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
)

// evaluateDelete converts a delete call to an SSA value.
func (f *Function) evaluateDelete(expr *expression.Expression) (ssa.Value, error) {
	value, err := f.evaluateRight(expr.Children[1])

	if err != nil {
		return nil, err
	}

	return f.delete(value)
}