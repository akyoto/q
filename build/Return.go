package build

import "github.com/akyoto/q/build/token"

// Return handles return statements.
func (state *State) Return(tokens []token.Token) error {
	state.Expect(token.Keyword)
	expression := tokens[1:]

	if len(expression) > 0 {
		err := state.TokensToRegister(expression, state.registers.ReturnValue[0])

		if err != nil {
			return err
		}
	}

	if state.ensureState.counter == 0 {
		state.assembler.Return()
		return nil
	}

	state.assembler.Jump("return")
	return nil
}
