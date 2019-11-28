package build

import (
	"sync/atomic"

	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/expression"
	"github.com/akyoto/q/build/token"
)

// Call represents a function call in the source code.
type Call struct {
	Function   *Function
	Parameters [][]token.Token
}

// Call handles function calls.
func (state *State) Call(tokens []token.Token) error {
	firstToken := tokens[0]

	if firstToken.Kind != token.Identifier {
		return errors.MissingFunctionName
	}

	lastToken := tokens[len(tokens)-1]

	if lastToken.Kind != token.GroupEnd {
		return &errors.MissingCharacter{Character: ")"}
	}

	return state.TokensToRegister(tokens, nil)
}

// BeforeCall pushes parameters into registers.
func (state *State) BeforeCall(parameters []*expression.Expression) error {
	for i, expr := range parameters {
		err := state.ExpressionToRegister(expr, state.registers.Call[i])

		if err != nil {
			return err
		}
	}

	return nil
}

// AfterCall restores saved registers from the stack.
func (state *State) AfterCall(function *Function) {
	atomic.AddInt64(&function.CallCount, 1)
}
