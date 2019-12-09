package build

import (
	"fmt"

	"github.com/akyoto/q/build/token"
)

// EnsureState handles the state of ensure compilation.
type EnsureState struct {
	counter int
	list    []Ensure
}

// Ensure represents an ensure statement.
type Ensure struct {
	condition []token.Token
	failLabel string
}

// Ensure specifies a condition that must be true for parameters.
func (state *State) Ensure(tokens []token.Token) error {
	if state.ignoreContracts {
		return nil
	}

	state.Skip(token.Keyword)
	condition := tokens[1:]

	state.ensureState.counter++
	failLabel := fmt.Sprintf("ensure_%d_fail", state.ensureState.counter)

	state.ensureState.list = append(state.ensureState.list, Ensure{
		condition: condition,
		failLabel: failLabel,
	})

	return nil
}
