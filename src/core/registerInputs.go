package core

import (
	"fmt"

	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// registerInputs registers every input as an identifier.
func (f *Function) registerInputs() {
	offset := 0

	for i, input := range f.Input {
		if input.Name == "_" {
			continue
		}

		structType, isStructType := input.Typ.(*types.Struct)

		if isStructType {
			structure := &ssa.Struct{Typ: structType}

			for _, field := range structType.Fields {
				param := &ssa.Parameter{
					Index:  uint8(offset + i),
					Name:   fmt.Sprintf("%s.%s", input.Name, field.Name),
					Typ:    field.Type,
					Tokens: input.Tokens,
				}

				f.Append(param)
				structure.Arguments = append(structure.Arguments, param)
				offset++
			}

			offset--
			f.Block().Identify(input.Name, structure)
			continue
		}

		input.Index = uint8(offset + i)
		f.Block().Identify(input.Name, input)
		f.Append(input)
	}
}