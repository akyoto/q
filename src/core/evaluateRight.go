package core

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
)

// evaluateRight evaluates the expression and automatically dereferences it if needed.
func (f *Function) evaluateRight(expr *expression.Expression) (ssa.Value, error) {
	value, err := f.evaluate(expr)

	if err != nil {
		return nil, err
	}

	return f.dereference(value), nil
}