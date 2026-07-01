package core

import (
	"unsafe"

	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// evaluateNew converts a new call to an SSA value.
func (f *Function) evaluateNew(expr *expression.Expression) (ssa.Value, error) {
	if len(expr.Children) == 1 {
		return nil, errors.NewAt(MissingType, f.File, expr.Children[0].Token.End()+1)
	}

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
		sizeInBytes = f.multiplySize(numElements, elementSize)
		mallocType = &types.Pointer{To: typ}
	} else {
		sizeInBytes = f.Append(&ssa.Int{Int: typ.Size()})
		mallocType = &types.Resource{Of: &types.Pointer{To: typ}}
	}

	malloc := f.Env.Function("mem", "alloc")
	f.Dependencies.Add(malloc)

	fn := &ssa.Function{
		FunctionRef: malloc,
		Typ: &types.Function{
			Output: []types.Type{mallocType},
		},
	}

	args := []ssa.Value{sizeInBytes}
	call := f.call(fn, args, expr.Source())

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