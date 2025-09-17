package core

import (
	"strings"

	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// compileInputs registers every input as an identifier.
func (f *Function) compileInputs() {
	offset := 0

	for i, input := range f.Input {
		structType, isStructType := types.Unwrap(input.Typ).(*types.Struct)

		if isStructType {
			if strings.HasPrefix(input.Name, "_") {
				offset += len(structType.Fields) - 1
				continue
			}

			structure := &ssa.Struct{
				Typ:    structType,
				Source: input.Source,
			}

			offset = f.composeStruct(structure, structType, input, i, offset)
			f.Block().Identify(input.Name, structure)
			continue
		}

		if strings.HasPrefix(input.Name, "_") {
			continue
		}

		input.Index = uint8(offset + i)
		f.Block().Identify(input.Name, input)
		f.Append(input)
	}
}