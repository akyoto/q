package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// delete frees an allocated value.
func (f *Function) delete(value ssa.Value) (ssa.Value, error) {
	free := f.Env.Function("mem", "free")
	f.Dependencies.Add(free)
	f.Block().Unidentify(value)

	switch valueType := types.Unwrap(value.Type()).(type) {
	case *types.Pointer:
		typ := valueType.To

		size := f.Append(&ssa.Int{
			Int: typ.Size(),
		})

		call := f.Append(&ssa.Call{
			Func: &ssa.Function{
				FunctionRef: free,
				Typ:         free.Type,
			},
			Arguments: []ssa.Value{value, size},
		})

		return call, nil

	case *types.Struct:
		structure := value.(*ssa.Struct)
		ptr := structure.Arguments[0]
		numElements := structure.Arguments[1]
		elementSize := types.Unwrap(ptr.Type()).(*types.Pointer).To.Size()
		sizeInBytes := f.multiplySize(numElements, elementSize)

		call := f.Append(&ssa.Call{
			Func: &ssa.Function{
				FunctionRef: free,
				Typ:         free.Type,
			},
			Arguments: []ssa.Value{ptr, sizeInBytes},
		})

		block := f.Block()
		block.Unidentify(ptr)
		block.Unidentify(numElements)
		return call, nil

	default:
		return nil, errors.New(&TypeMismatch{Encountered: valueType.Name(), Expected: types.AnyPointer.Name()}, f.File, value.(errors.Source))
	}
}