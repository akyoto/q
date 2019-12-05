package build

import (
	"github.com/akyoto/q/build/token"
)

// Ensure specifies a condition that must be true for return values.
func (state *State) Ensure(tokens []token.Token) error {
	if state.ignoreContracts {
		return nil
	}

	state.Expect(token.Keyword)
	// condition := tokens[1:]
	return nil
}
