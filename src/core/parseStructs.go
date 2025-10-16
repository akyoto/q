package core

import (
	"iter"

	"git.urbach.dev/cli/q/src/types"
)

// parseStructs parses the tokens of the struct field types.
func (env *Environment) parseStructs(structs iter.Seq[*types.Struct]) error {
	processed := map[*types.Struct]recursionState{}

	for structure := range structs {
		err := env.parseStruct(structure, processed)

		if err != nil {
			return err
		}
	}

	return nil
}