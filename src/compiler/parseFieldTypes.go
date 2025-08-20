package compiler

import (
	"iter"

	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/types"
)

// parseFieldTypes parses the tokens of the struct field types.
func parseFieldTypes(structs iter.Seq[*types.Struct], env *core.Environment) error {
	for structure := range structs {
		offset := 0

		for i, field := range structure.Fields {
			field.Type = core.ParseType(field.Tokens[1:], structure.File.Bytes, env)

			if field.Type == nil {
				return errors.New(&UnknownType{Name: field.Tokens[1:].String(structure.File.Bytes)}, structure.File, field.Tokens[1].Position)
			}

			field.Index = uint8(i)
			field.Offset = uint8(offset)
			offset += field.Type.Size()
		}
	}

	return nil
}