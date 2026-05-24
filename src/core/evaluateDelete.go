package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// evaluateDelete converts a delete call to an SSA value.
func (f *Function) evaluateDelete(expr *expression.Expression) (ssa.Value, error) {
	value, err := f.evaluateRight(expr.Children[1])

	if err != nil {
		return nil, err
	}

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
			Source:    expr.Source(),
		})

		return call, nil

	case *types.Struct:
		var (
			structure   = value.(*ssa.Struct)
			ptr         = structure.Arguments[0]
			length      = structure.Arguments[1]
			elementSize = types.Unwrap(ptr.Type()).(*types.Pointer).To.Size()
			sizeInBytes ssa.Value
		)

		if elementSize == 1 {
			sizeInBytes = length
		} else {
			elementSizeValue := f.Append(&ssa.Int{Int: elementSize})

			sizeInBytes = f.Append(&ssa.BinaryOp{
				Op:    token.Mul,
				Left:  length,
				Right: elementSizeValue,
			})
		}

		call := f.Append(&ssa.Call{
			Func: &ssa.Function{
				FunctionRef: free,
				Typ:         free.Type,
			},
			Arguments: []ssa.Value{ptr, sizeInBytes},
			Source:    expr.Source(),
		})

		block := f.Block()
		block.Unidentify(ptr)
		block.Unidentify(length)
		return call, nil

	default:
		return nil, errors.New(&TypeMismatch{Encountered: valueType.Name(), Expected: types.AnyPointer.Name()}, f.File, expr.Children[1].Source())
	}
}