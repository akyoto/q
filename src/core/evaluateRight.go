package core

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// evaluateRight creates a load from memory if the value
// was a memory address, otherwise returns the raw value.
func (f *Function) evaluateRight(expr *expression.Expression) (ssa.Value, error) {
	value, err := f.evaluate(expr)

	if err != nil {
		return nil, err
	}

	data, isData := value.(*ssa.Data)

	if isData {
		zero := f.Append(&ssa.Int{Int: 0})

		load := f.Append(&ssa.Load{
			Memory: &ssa.Memory{
				Address: data,
				Index:   zero,
				Typ:     data.Typ.(*types.Pointer).To,
			},
			Source: data.Source,
		})

		return load, nil
	}

	memory, isMemory := value.(*ssa.Memory)

	if !isMemory {
		return value, nil
	}

	switch typ := memory.Typ.(type) {
	case *types.Struct:
		structure := f.loadFields(memory, typ, expr.Source())
		return structure, nil

	default:
		load := f.Append(&ssa.Load{
			Memory: memory,
			Source: memory.Source,
		})

		return load, nil
	}
}