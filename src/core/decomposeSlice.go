package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// decomposeSlice decomposes a slices to its pointer, type and length.
func (f *Function) decomposeSlice(addressValue ssa.Value) (ssa.Value, types.Type, ssa.Value, error) {
	switch addressType := types.Unwrap(addressValue.Type()).(type) {
	case *types.Struct:
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
		return nil, nil, nil, errors.New(&TypeNotIndexable{TypeName: addressType.Name()}, f.File, addressValue.(errors.Source))
	}
}