package resolver

import (
	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/types"
)

// parseStruct parses the field tokens of a single struct.
func parseStruct(env *core.Environment, structure *types.Struct, processed map[*types.Struct]recursionState) error {
	if processed[structure] == finished {
		return nil
	}

	processed[structure] = started
	offset := 0

	for i, field := range structure.Fields {
		file := structure.File.(*fs.File)
		typ, err := env.TypeFromTokens(field.Tokens[1:], file)

		if err != nil {
			return err
		}

		dependency, isStruct := typ.(*types.Struct)

		if isStruct {
			switch processed[dependency] {
			case notStarted:
				err := parseStruct(env, dependency, processed)

				if err != nil {
					return err
				}
			case started:
				return errors.New(&core.CycleDetected{A: structure.Name(), B: dependency.Name()}, file, field.Tokens[1:])
			case finished:
			}
		}

		field.Type = typ
		field.Index = uint8(i)
		field.Offset = uint8(offset)
		offset += field.Type.Size()
	}

	processed[structure] = finished
	return nil
}