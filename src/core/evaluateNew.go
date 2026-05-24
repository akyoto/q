package core

import (
	"unsafe"

	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// evaluateNew converts a new call to an SSA value.
func (f *Function) evaluateNew(expr *expression.Expression) (ssa.Value, error) {
	right := (*expression.TypeExpression)(unsafe.Pointer(expr.Children[1]))
	typ, err := f.Env.TypeFromTokens(right.Tokens, f.File)

	if err != nil {
		return nil, err
	}

	var (
		isSlice     = len(expr.Children) == 3
		numElements ssa.Value
		sizeInBytes ssa.Value
		sliceType   types.Type
		mallocType  types.Type
	)

	if isSlice {
		numElements, err = f.evaluateRight(expr.Children[2])

		if err != nil {
			return nil, err
		}

		sliceType = f.Env.Slice(typ)
		elementSize := typ.Size()

		if elementSize == 1 {
			sizeInBytes = numElements
		} else {
			elementSizeValue := f.Append(&ssa.Int{Int: elementSize})

			sizeInBytes = f.Append(&ssa.BinaryOp{
				Op:    token.Mul,
				Left:  numElements,
				Right: elementSizeValue,
			})
		}

		mallocType = &types.Pointer{To: typ}
	} else {
		sizeInBytes = f.Append(&ssa.Int{Int: typ.Size()})
		mallocType = &types.Resource{Of: &types.Pointer{To: typ}}
	}

	malloc := f.Env.Function("mem", "alloc")
	f.Dependencies.Add(malloc)

	call := f.Append(&ssa.Call{
		Func: &ssa.Function{
			FunctionRef: malloc,
			Typ: &types.Function{
				Output: []types.Type{mallocType},
			},
		},
		Arguments: []ssa.Value{sizeInBytes},
		Source:    expr.Source(),
	})

	if isSlice {
		structure := &ssa.Struct{
			Typ:       &types.Resource{Of: sliceType},
			Arguments: ssa.Arguments{call, numElements},
			Source:    expr.Source(),
		}

		return structure, nil
	}

	return call, nil
}