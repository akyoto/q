package core

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// convertToSlice evaluates a static array as a slice value.
func (f *Function) convertToSlice(address ssa.Value, index ssa.Value, scale bool, typ *types.Struct, source ssa.Source) ssa.Value {
	element := typ.Fields[0].Type

	if index != nil {
		if scale {
			index = f.multiplySize(index, typ.Size())
		}

		address = f.Append(&ssa.BinaryOp{
			Op:    token.Add,
			Left:  address,
			Right: index,
		})
	}

	ptr := f.Append(&ssa.Copy{
		Value: address,
		Typ:   f.Env.Pointer(element),
		Read:  true,
	})

	length := f.Append(&ssa.Int{
		Int: len(typ.Fields),
	})

	return f.makeStruct(f.Env.Slice(element), []ssa.Value{ptr, length}, source)
}