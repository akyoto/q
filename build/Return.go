package build

import "github.com/akyoto/q/build/token"

// Return handles return statements.
func (state *State) Return(tokens []token.Token) error {
	expression := tokens[1:]

	if len(expression) > 0 {
		err := state.TokensToRegister(expression, state.registers.ReturnValue[0])

		if err != nil {
			return err
		}
	}

	state.assembler.Return()
	return nil
}
