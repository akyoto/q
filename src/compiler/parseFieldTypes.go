package compiler

import (
	"iter"

	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/fs"
	"git.urbach.dev/cli/q/src/types"
)

// parseFieldTypes parses the tokens of the struct field types.
func parseFieldTypes(structs iter.Seq[*types.Struct], env *core.Environment) error {
	for structure := range structs {
		offset := 0

		for i, field := range structure.Fields {
			file := structure.File.(*fs.File)
			typ, err := env.TypeFromTokens(field.Tokens[1:], file)

			if err != nil {
				return err
			}

			field.Type = typ
			field.Index = uint8(i)
			field.Offset = uint8(offset)
			offset += field.Type.Size()
		}
	}

	return nil
}