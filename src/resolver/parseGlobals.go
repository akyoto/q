package resolver

import (
	"iter"

	"git.urbach.dev/cli/q/src/core"
)

// parseGlobals parses the tokens of global variables.
func parseGlobals(env *core.Environment, globals iter.Seq[*core.Global]) error {
	for global := range globals {
		typ, err := env.TypeFromTokens(global.Tokens[1:], global.File)

		if err != nil {
			return err
		}

		global.Typ = typ
	}

	return nil
}