package core

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// makeSlice creates a new slice.
func (f *Function) makeSlice(address ssa.Value, from ssa.Value, to ssa.Value, source ssa.Source) *ssa.Struct {
	newPointer := f.Append(&ssa.BinaryOp{
		Op:    token.Add,
		Left:  address,
		Right: from,
	})

	newLength := f.Append(&ssa.BinaryOp{
		Op:    token.Sub,
		Left:  to,
		Right: from,
	})

	return f.makeStruct(types.String, []ssa.Value{newPointer, newLength}, source)
}