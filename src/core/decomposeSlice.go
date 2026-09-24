package core

import (
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// decomposeSlice decomposes a slices to its pointer, type and length.
func (f *Function) decomposeSlice(addressValue ssa.Value) (ssa.Value, types.Type, ssa.Value, error) {
	switch addressType := types.Unwrap(addressValue.Type()).(type) {
	case *types.Struct:
		if addressType.IsArray() {
			pointerType := f.Env.Pointer(addressType.Fields[0].Type)
			memory := addressValue.(*ssa.Memory)
			address := memory.Address

			if memory.Index != nil {
				integer, isInt := memory.Index.(*ssa.Int)

				if isInt && integer.Int == 0 {
					return address, pointerType, nil, nil
				}

				index := memory.Index

				if memory.Scale {
					size := memory.Typ.Size()

					if size != 1 {
						index = f.Append(&ssa.BinaryOp{Op: token.Mul, Left: memory.Index, Right: f.Append(&ssa.Int{Int: size})})
					}
				}

				address = f.Append(&ssa.BinaryOp{
					Op:    token.Add,
					Left:  address,
					Right: index,
				})
			}

			return address, pointerType, nil, nil
		}

		structure, isStructure := addressValue.(*ssa.Struct)

		if !isStructure {
			panic("not implemented")
		}

		pointer := structure.Arguments[0]
		length := structure.Arguments[1]
		return pointer, pointer.Type(), length, nil
	case *types.Pointer:
		structure, isStructure := addressType.To.(*types.Struct)

		if isStructure && structure.IsArray() {
			return addressValue, f.Env.Pointer(structure.Fields[0].Type), nil, nil
		}

		return addressValue, addressType, nil, nil
	default:
		panic("not implemented")
	}
}