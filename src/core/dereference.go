package core

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// dereference loads from memory if the value was a memory address, otherwise returns the raw value.
func (f *Function) dereference(value ssa.Value) ssa.Value {
	var memory *ssa.Memory

	switch v := value.(type) {
	case *ssa.Global:
		typ := v.Typ.(*types.Pointer).To
		structure, isStructure := typ.(*types.Struct)

		if isStructure && structure.IsArray() {
			return f.convertToSlice(v, nil, false, structure, v.Source)
		}

		zero := f.Append(&ssa.Int{Int: 0})

		memory = &ssa.Memory{
			Address: v,
			Index:   zero,
			Typ:     typ,
			Source:  v.Source,
		}
	case *ssa.Memory:
		memory = v
	default:
		return value
	}

	typ, isStruct := memory.Typ.(*types.Struct)

	if isStruct {
		if typ.IsArray() {
			return f.convertToSlice(memory.Address, memory.Index, memory.Scale, typ, memory.Source)
		}

		return f.loadFields(memory, typ, memory.Source)
	}

	return f.Append(&ssa.Load{
		Memory: memory,
		Source: memory.Source,
	})
}