package core

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// makeStruct creates a struct and saves a reference to it for every field.
// This is later used in dead code removal to know if a value is part of a partially used struct.
func (f *Function) makeStruct(typ types.Type, fields []ssa.Value, source ssa.Source) *ssa.Struct {
	structure := &ssa.Struct{
		Typ:       typ,
		Arguments: fields,
		Source:    source,
	}

	if f.valueToStruct == nil {
		f.valueToStruct = map[ssa.Value]*ssa.Struct{}
	}

	for _, value := range fields {
		f.valueToStruct[value] = structure
	}

	return structure
}