package core

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// matchesType returns true if the nodes match the given parameter types.
func (f *Function) matchesType(nodes []*expression.Expression, parameters []*ssa.Parameter) (bool, error) {
	for i, node := range nodes {
		value, err := f.evaluate(node)

		if err != nil {
			return false, err
		}

		if parameters != nil && !types.Is(value.Type(), parameters[i].Typ) {
			return false, nil
		}
	}

	return true, nil
}