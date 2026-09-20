package core

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// inputRegisterCount returns how many registers an input occupies.
func inputRegisterCount(input *ssa.Parameter) int {
	structType, isStructType := types.Unwrap(input.Typ).(*types.Struct)

	if !isStructType {
		return 1
	}

	if structType.Size() > 16 {
		registers := 0

		for _, field := range structType.Fields {
			nestedType, isNested := types.Unwrap(field.Type).(*types.Struct)

			if isNested {
				registers += inputRegisterCount(&ssa.Parameter{Typ: nestedType})
				continue
			}

			registers++
		}

		return registers
	}

	size := 0
	registers := 1

	for _, field := range structType.Fields {
		fieldSize := field.Type.Size()

		if size+fieldSize > 8 {
			registers++
			size = 0
		}

		size += fieldSize
	}

	return registers
}