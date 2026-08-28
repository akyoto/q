package core

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// updateStruct creates a new struct with the field replaced and binds the name and field identifier to the new values.
func (f *Function) updateStruct(structure *ssa.Struct, field *types.Field, name string, value ssa.Value) {
	args := make(ssa.Arguments, len(structure.Arguments))
	copy(args, structure.Arguments)
	args[field.Index] = value

	newStruct := &ssa.Struct{
		Typ:       structure.Typ,
		Arguments: args,
		Source:    structure.Source,
	}

	f.Block().Identify(name, newStruct)
	f.Block().Identify(name+"."+field.Name, value)
}