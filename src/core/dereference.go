package core

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// dereference loads from memory if the value was a memory address, otherwise returns the raw value.
func (f *Function) dereference(value ssa.Value) ssa.Value {
	switch value := value.(type) {
	case *ssa.Global:
		zero := f.Append(&ssa.Int{Int: 0})

		memory := &ssa.Memory{
			Address: value,
			Index:   zero,
			Typ:     value.Typ.(*types.Pointer).To,
		}

		switch typ := value.Typ.(*types.Pointer).To.(type) {
		case *types.Struct:
			structure := f.loadFields(memory, typ, value.Source)
			return structure

		default:
			load := f.Append(&ssa.Load{
				Memory: memory,
				Source: value.Source,
			})

			return load
		}

	case *ssa.Memory:
		switch typ := value.Typ.(type) {
		case *types.Struct:
			structure := f.loadFields(value, typ, value.Source)
			return structure

		default:
			load := f.Append(&ssa.Load{
				Memory: value,
				Source: value.Source,
			})

			return load
		}

	default:
		return value
	}
}