package core

import (
	"iter"
)

// parseGlobals parses the tokens of global variables.
func (env *Environment) parseGlobals(globals iter.Seq[*Global]) error {
	for global := range globals {
		typ, err := env.TypeFromTokens(global.Tokens[1:], global.File)

		if err != nil {
			return err
		}

		global.Typ = typ
	}

	return nil
}