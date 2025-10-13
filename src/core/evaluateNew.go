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
		size        ssa.Value
		sliceType   types.Type
	)

	if isSlice {
		numElements, err = f.evaluateRight(expr.Children[2])

		if err != nil {
			return nil, err
		}

		sliceType = f.Env.Slice(typ)
		elementSize := f.Append(&ssa.Int{Int: typ.Size()})

		size = f.Append(&ssa.BinaryOp{
			Op:    token.Mul,
			Left:  elementSize,
			Right: numElements,
		})
	} else {
		size = f.Append(&ssa.Int{Int: typ.Size()})
	}

	malloc := f.Env.Function("mem", "alloc")
	f.Dependencies.Add(malloc)

	call := f.Append(&ssa.Call{
		Func: &ssa.Function{
			FunctionRef: malloc,
			Typ: &types.Function{
				Output: []types.Type{&types.Pointer{To: typ}},
			},
		},
		Arguments: []ssa.Value{size},
		Source:    expr.Source(),
	})

	if isSlice {
		structure := &ssa.Struct{
			Typ:       sliceType,
			Arguments: ssa.Arguments{call, numElements},
			Source:    expr.Source(),
		}

		return structure, nil
	}

	return call, nil
}