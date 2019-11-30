package build

import (
	"github.com/akyoto/q/build/token"
)

// identifiersLastUse returns a map of variable names
// mapped to the position they were last used.
func (state *State) identifiersLastUse() map[string]token.Position {
	// scopeLevel := 0
	identifiers := map[string]token.Position{}

	for i := len(state.tokens) - 1; i >= 0; i-- {
		t := state.tokens[i]

		if t.Kind != token.Identifier {
			continue
		}

		identifier := t.Text()
		_, exists := identifiers[identifier]

		if exists {
			continue
		}

		identifiers[identifier] = i
	}

	return identifiers
}

// killVariables frees the registers of all variables that die in the given token range.
func (state *State) killVariables(identifiers map[string]token.Position, from int, until int) {
	for identifier, deathPos := range identifiers {
		if deathPos < from || deathPos >= until {
			continue
		}

		variable := state.scopes.Get(identifier)

		if variable != nil {
			variable.Register().Free()
			// fmt.Println(variable, "died at", state.tokens[:deathPos+1])
		}

		delete(identifiers, identifier)
	}
}
