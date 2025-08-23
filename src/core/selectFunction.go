package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
)

// selectFunction selects the correct variant of a function based on the input types.
func (f *Function) selectFunction(variants []*Function, inputExpressions []*expression.Expression, identifier *expression.Expression) (*Function, error) {
	if len(variants) == 1 {
		fn := variants[0]

		if len(inputExpressions) != len(fn.Input) {
			return nil, errors.New(&ParameterCountMismatch{Function: fn.FullName, Count: len(inputExpressions), ExpectedCount: len(fn.Input)}, f.File, identifier.Token.End())
		}

		return fn, nil
	}

	for _, variant := range variants {
		if len(inputExpressions) != len(variant.Input) {
			continue
		}

		reset := len(f.Block().Instructions)
		matches, err := f.matchesType(inputExpressions, variant.Input)
		f.Block().Instructions = f.Block().Instructions[:reset]

		if err != nil {
			return nil, err
		}

		if matches {
			return variant, nil
		}
	}

	return nil, nil
}