package build

import (
	"github.com/akyoto/q/build/expression"
	"github.com/akyoto/q/build/token"
)

// Require specifies a condition that must be true for parameters.
func (state *State) Require(tokens []token.Token) error {
	if state.ignoreContracts {
		return nil
	}

	state.Expect(token.Keyword)
	condition := tokens[1:]
	expr, err := expression.FromTokens(condition)

	if err != nil {
		return err
	}

	state.function.Require = append(state.function.Require, expr)
	return nil
}
