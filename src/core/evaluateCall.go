package core

import (
	"unsafe"

	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// evaluateCall converts a call expression to an SSA value.
func (f *Function) evaluateCall(expr *expression.Expression) (ssa.Value, error) {
	identifier := expr.Children[0]

	if identifier.Token.Kind.IsBuiltin() {
		switch identifier.Token.Kind {
		case token.New:
			right := (*expression.TypeExpression)(unsafe.Pointer(expr.Children[1]))
			typ, err := f.Env.TypeFromTokens(right.Tokens, f.File)

			if err != nil {
				return nil, err
			}

			malloc := f.Env.Function("mem", "alloc")
			returnType := &types.Pointer{To: typ}

			size := f.Append(&ssa.Int{
				Int: typ.Size(),
			})

			call := f.Append(&ssa.Call{
				Func: &ssa.Function{
					FunctionRef: malloc,
					Typ: &types.Function{
						Output: []types.Type{returnType},
					},
				},
				Arguments: []ssa.Value{size},
				Source:    expr.Source(),
			})

			f.Dependencies.Add(malloc)
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
		}
	}

	funcValue, err := f.evaluate(identifier)

	if err != nil {
		return nil, err
	}

	ssaFunc, isFunction := funcValue.(*ssa.Function)

	if !isFunction {
		return nil, errors.New(InvalidCallExpression, f.File, identifier.Source().StartPos)
	}

	fn := ssaFunc.FunctionRef.(*Function)
	inputExpressions := expr.Children[1:]
	args, err := f.decompose(inputExpressions, fn.Input, false)

	if err != nil {
		return nil, err
	}

	if fn.IsExtern() {
		v := f.Append(&ssa.CallExtern{Call: ssa.Call{
			Func:      ssaFunc,
			Arguments: args,
			Source:    expr.Source(),
		}})

		return v, nil
	}

	if f == f.Env.Init && fn == f.Env.Main {
		f.runAll("init")
	}

	v := f.Append(&ssa.Call{
		Func:      ssaFunc,
		Arguments: args,
		Source:    expr.Source(),
	})

	if f == f.Env.Init && fn == f.Env.Main {
		f.runAll("exit")
	}

	return v, nil
}