package core

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// decomposeSlice decomposes a slices to its pointer, type and length.
func (f *Function) decomposeSlice(addressValue ssa.Value) (ssa.Value, types.Type, ssa.Value) {
	addressType := types.Unwrap(addressValue.Type())

	switch addressType.(type) {
	case *types.Struct:
		switch value := addressValue.(type) {
		case *ssa.Struct:
			pointer := value.Arguments[0]
			length := value.Arguments[1]
			return pointer, pointer.Type(), length

		case *ssa.Call:
			pointer := f.Append(&ssa.Field{
				Tuple: value,
				Index: 0,
			})

			length := f.Append(&ssa.Field{
				Tuple: value,
				Index: 1,
			})

			return pointer, pointer.Type(), length

		default:
			panic("not implemented")
		}

	case *types.Pointer:
		return addressValue, addressType, nil

	default:
		panic("not implemented")
	}
}