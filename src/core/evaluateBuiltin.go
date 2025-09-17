package core

import (
	"unsafe"

	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// evaluateBuiltin converts a call to a builtin function to an SSA value.
func (f *Function) evaluateBuiltin(expr *expression.Expression) (ssa.Value, error) {
	switch expr.Children[0].Token.Kind {
	case token.New:
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
			numElements, err = f.evaluate(expr.Children[2])

			if err != nil {
				return nil, err
			}

			sliceType = types.Slice(typ, "[]"+typ.Name())
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

	case token.Delete:
		value, err := f.evaluate(expr.Children[1])

		if err != nil {
			return nil, err
		}

		free := f.Env.Function("mem", "free")
		f.Dependencies.Add(free)
		f.Block().Unidentify(value)

		switch valueType := value.Type().(type) {
		case *types.Pointer:
			typ := valueType.To

			size := f.Append(&ssa.Int{
				Int: typ.Size(),
			})

			call := f.Append(&ssa.Call{
				Func: &ssa.Function{
					FunctionRef: free,
					Typ:         &types.Function{},
				},
				Arguments: []ssa.Value{value, size},
				Source:    expr.Source(),
			})

			return call, nil

		default:
			return nil, errors.New(&TypeMismatch{Encountered: valueType.Name(), Expected: types.AnyPointer.Name()}, f.File, expr.Children[1].Source().StartPos)
		}

	case token.Syscall:
		args, err := f.decompose(expr.Children[1:], nil, false)

		if err != nil {
			return nil, err
		}

		syscall := &ssa.Syscall{
			Arguments: args,
			Source:    expr.Source(),
		}

		return f.Append(syscall), nil

	default:
		panic("not implemented")
	}
}