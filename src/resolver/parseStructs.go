package resolver

import (
	"iter"

	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/types"
)

// parseStructs parses the tokens of the struct field types.
func parseStructs(env *core.Environment, structs iter.Seq[*types.Struct]) error {
	processed := map[*types.Struct]recursionState{types.String: finished}

	for structure := range structs {
		err := parseStruct(env, structure, processed)

		if err != nil {
			return err
		}
	}

	return nil
}