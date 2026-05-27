package core

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// loadFields loads all fields from memory into a struct.
func (f *Function) loadFields(memory *ssa.Memory, typ *types.Struct, source ssa.Source) *ssa.Struct {
	fields := make([]ssa.Value, len(typ.Fields))

	for i, field := range typ.Fields {
		fields[i] = f.Append(&ssa.Load{
			Memory: f.structField(memory, field),
			Source: source,
		})
	}

	return f.makeStruct(typ, fields, source)
}