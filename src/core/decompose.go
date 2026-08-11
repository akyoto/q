package core

import (
	"git.urbach.dev/cli/q/src/ssa"
)

// decompose splits structs into their individual fields.
func (f *Function) decompose(values []ssa.Value) ([]ssa.Value, error) {
	args := make([]ssa.Value, 0, len(values))

	for _, value := range values {
		structure, isStruct := value.(*ssa.Struct)

		if isStruct {
			args = f.decomposeStruct(args, structure)
			continue
		}

		args = append(args, value)
	}

	return args, nil
}