package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
)

// selectFunction selects the correct variant of a function based on the input types.
func (f *Function) selectFunction(fn *Function, inputExpressions []*expression.Expression, identifier *expression.Expression) (*Function, error) {
	if fn.Next == nil {
		if len(inputExpressions) != len(fn.Input) {
			return nil, errors.New(&ParameterCountMismatch{Function: fn.FullName, Count: len(inputExpressions), ExpectedCount: len(fn.Input)}, f.File, identifier.Source().EndPos)
		}

		return fn, nil
	}

	for fn != nil {
		if len(inputExpressions) != len(fn.Input) {
			fn = fn.Next
			continue
		}

		reset := len(f.Block().Instructions)
		matches, err := f.matchesType(inputExpressions, fn.Input)
		f.Block().Instructions = f.Block().Instructions[:reset]

		if err != nil {
			return nil, err
		}

		if matches {
			return fn, nil
		}

		fn = fn.Next
	}

	return nil, nil
}