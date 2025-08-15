package compiler

import (
	"iter"

	"git.urbach.dev/cli/q/src/core"
	"git.urbach.dev/cli/q/src/types"
)

// parseFieldTypes parses the tokens of the struct field types.
func parseFieldTypes(structs iter.Seq[*types.Struct], env *core.Environment) {
	for structure := range structs {
		offset := 0

		for i, field := range structure.Fields {
			field.Type = core.ParseType(field.Tokens[1:], structure.File.Bytes, env)
			field.Index = uint8(i)
			field.Offset = uint8(offset)
			offset += field.Type.Size()
		}
	}
}