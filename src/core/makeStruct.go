package core

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// makeStruct creates a struct and saves a reference to it for every field.
// This is later used in dead code removal to know if a value is part of a partially used struct.
func (f *Function) makeStruct(typ types.Type, source ssa.Source, fields []ssa.Value) *ssa.Struct {
	structure := &ssa.Struct{
		Typ:       typ,
		Source:    source,
		Arguments: fields,
	}

	if f.valueToStruct == nil {
		f.valueToStruct = map[ssa.Value]*ssa.Struct{}
	}

	for _, value := range fields {
		f.valueToStruct[value] = structure
	}

	return structure
}