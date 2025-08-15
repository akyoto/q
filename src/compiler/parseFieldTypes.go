package compiler

import (
	"iter"

	"git.urbach.dev/cli/q/src/types"
)

// parseFieldTypes parses the tokens of the struct field types.
func parseFieldTypes(structs iter.Seq[*types.Struct]) {
	for structure := range structs {
		offset := 0

		for i, field := range structure.Fields {
			field.Type = types.Parse(field.Tokens[1:], structure.File.Bytes)
			field.Index = uint8(i)
			field.Offset = uint8(offset)
			offset += field.Type.Size()
		}
	}
}