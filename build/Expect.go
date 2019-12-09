package build

import (
	"fmt"

	"github.com/akyoto/q/build/token"
)

// ExpectState handles the state of expect compilation.
type ExpectState struct {
	counter int
	list    []Expect
}

// Expect represents an expect statement.
type Expect struct {
	condition []token.Token
	failLabel string
}

// Expect specifies a condition that must be true for parameters.
func (state *State) Expect(tokens []token.Token) error {
	if state.ignoreContracts {
		return nil
	}

	state.Skip(token.Keyword)
	condition := tokens[1:]

	state.expectState.counter++
	failLabel := fmt.Sprintf("expect_%d_fail", state.expectState.counter)

	state.expectState.list = append(state.expectState.list, Expect{
		condition: condition,
		failLabel: failLabel,
	})

	return state.Condition(condition, failLabel)
}
