package core

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// decomposeSlice decomposes a slices to its pointer, type and length.
func (f *Function) decomposeSlice(addressValue ssa.Value) (ssa.Value, types.Type, ssa.Value, error) {
	switch addressType := types.Unwrap(addressValue.Type()).(type) {
	case *types.Struct:
		if addressType.IsArray() {
			pointerType := &types.Pointer{To: addressType.Fields[0].Type}
			return addressValue.(*ssa.Memory).Address, pointerType, nil, nil
		}

		structure, isStructure := addressValue.(*ssa.Struct)

		if !isStructure {
			panic("not implemented")
		}

		pointer := structure.Arguments[0]
		length := structure.Arguments[1]
		return pointer, pointer.Type(), length, nil
	case *types.Pointer:
		return addressValue, addressType, nil, nil
	default:
		panic("not implemented")
	}
}