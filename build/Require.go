package build

import (
	"fmt"

	"github.com/akyoto/q/build/token"
)

// RequireState handles the state of require compilation.
type RequireState struct {
	counter int
	list    []Require
}

// Require represents a require statement.
type Require struct {
	condition []token.Token
	failLabel string
}

// Require specifies a condition that must be true for parameters.
func (state *State) Require(tokens []token.Token) error {
	if state.ignoreContracts {
		return nil
	}

	state.Expect(token.Keyword)
	condition := tokens[1:]

	state.requireState.counter++
	failLabel := fmt.Sprintf("require_%d_fail", state.requireState.counter)

	state.requireState.list = append(state.requireState.list, Require{
		condition: condition,
		failLabel: failLabel,
	})

	return state.Condition(condition, failLabel)
}
