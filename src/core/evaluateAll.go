package core

import (
	"slices"

	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
)

// evaluateAll evalutes all expressions and returns the list.
func (f *Function) evaluateAll(expressions []*expression.Expression) ([]ssa.Value, error) {
	values := make([]ssa.Value, len(expressions))

	for i, expr := range slices.Backward(expressions) {
		value, err := f.evaluateRight(expr)

		if err != nil {
			return nil, err
		}

		values[i] = value
	}

	return values, nil
}