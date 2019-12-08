package build

import (
	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/token"
)

// Return handles return statements.
func (state *State) Return(tokens []token.Token) error {
	state.Expect(token.Keyword)
	expression := tokens[1:]

	if len(expression) > 0 {
		if len(state.function.ReturnTypes) == 0 {
			return errors.ReturnWithoutFunctionType
		}

		err := state.TokensToRegister(expression, state.registers.ReturnValue[0])

		if err != nil {
			return err
		}
	} else if len(state.function.ReturnTypes) > 0 {
		return &errors.MissingReturnValue{ReturnType: state.function.ReturnTypes[0].Name}
	}

	if state.ensureState.counter == 0 {
		state.assembler.Return()
		return nil
	}

	state.assembler.Jump("return")
	return nil
}
