package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// selectFunction selects the correct variant of a function based on the input types.
func (f *Function) selectFunction(fn *Function, values []ssa.Value, identifier *expression.Expression) (*Function, error) {
	if fn.Next == nil {
		if len(values) != len(fn.Input) {
			return nil, errors.NewAt(&ParameterCountMismatch{Function: fn.FullName, Count: len(values), ExpectedCount: len(fn.Input)}, f.File, identifier.Source().End())
		}

		return fn, nil
	}

match:
	for fn != nil {
		if len(values) != len(fn.Input) {
			fn = fn.Next
			continue match
		}

		for i, value := range values {
			if !types.Is(value.Type(), fn.Input[i].Typ) {
				fn = fn.Next
				continue match
			}
		}

		return fn, nil
	}

	return nil, nil
}