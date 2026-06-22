package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// decomposeSlice decomposes a slices to its pointer, type and length.
func (f *Function) decomposeSlice(addressValue ssa.Value, source token.Source) (ssa.Value, types.Type, ssa.Value, error) {
	addressType := types.Unwrap(addressValue.Type())

	switch addressType.(type) {
	case *types.Struct:
		switch value := addressValue.(type) {
		case *ssa.Struct:
			pointer := value.Arguments[0]
			length := value.Arguments[1]
			return pointer, pointer.Type(), length, nil

		case *ssa.Call:
			resource, isResource := addressValue.Type().(*types.Resource)

			if isResource {
				return nil, nil, nil, errors.New(&ResourceNotConsumed{TypeName: resource.Name()}, f.File, source)
			}

			pointer := f.Append(&ssa.Field{
				Tuple: value,
				Index: 0,
			})

			length := f.Append(&ssa.Field{
				Tuple: value,
				Index: 1,
			})

			return pointer, pointer.Type(), length, nil

		default:
			panic("not implemented")
		}

	case *types.Pointer:
		return addressValue, addressType, nil, nil

	default:
		panic("not implemented")
	}
}