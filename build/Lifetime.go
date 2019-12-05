package build

import (
	"github.com/akyoto/q/build/token"
)

// IdentifierLifeTimeMap returns a map of variable names
// mapped to the position they were last used.
func IdentifierLifeTimeMap(tokens []token.Token) map[string]token.Position {
	identifiers := map[string]token.Position{}

	for i := len(tokens) - 1; i >= 0; i-- {
		t := tokens[i]

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

// KillVariables frees the registers of all variables that die in the given token range.
func (state *State) KillVariables(from int, until int) {
	state.scopes.Each(func(variable *Variable) {
		if variable.AliveUntil < from || variable.AliveUntil >= until {
			return
		}

		variable.Register().Free()
		// fmt.Println(variable, "died at", state.tokens[:variable.AliveUntil+1])
		delete(state.identifierLifeTime, variable.Name)
	})
}

// InstructionEndPosition returns the token position of the next instruction.
func (state *State) InstructionEndPosition() token.Position {
	instr := state.instructions[state.instrCursor]
	return instr.Position + len(instr.Tokens)
}
