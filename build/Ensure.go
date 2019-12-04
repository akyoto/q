package build

import (
	"github.com/akyoto/q/build/expression"
	"github.com/akyoto/q/build/token"
)

// Ensure specifies a condition that must be true for return values.
func (state *State) Ensure(tokens []token.Token) error {
	if state.ignoreContracts {
		return nil
	}

	state.Expect(token.Keyword)
	condition := tokens[1:]
	expr, err := expression.FromTokens(condition)

	if err != nil {
		return err
	}

	state.function.Ensure = append(state.function.Ensure, expr)
	return nil
}
