package core

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// storeFields stores all struct fields into memory.
func (f *Function) storeFields(memory *ssa.Memory, typ *types.Struct, fields []ssa.Value) {
	for i, field := range typ.Fields {
		f.Append(&ssa.Store{
			Memory: f.structField(memory, field),
			Value:  fields[i],
		})
	}
}