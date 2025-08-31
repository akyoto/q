package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// evaluateCall converts a call expression to an SSA value.
func (f *Function) evaluateCall(expr *expression.Expression) (ssa.Value, error) {
	identifier := expr.Children[0]

	if identifier.Token.Kind == token.Identifier {
		switch identifier.String(f.File.Bytes) {
		case "new":
			typ := ParseType([]token.Token{expr.Children[1].Token}, f.File.Bytes, f.Env)
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

		case "syscall":
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

	// pkg := f.Env.Packages[ssaFunc.Package]
	// funcName := ssaFunc.Name
	// funcName, _, _ = strings.Cut(funcName, "[")
	// generic := pkg.Functions[funcName]
	// var fn *Function

	// for variant := range generic.Variants {
	// 	if variant.Name == ssaFunc.Name {
	// 		fn = variant
	// 		break
	// 	}
	// }
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

	v := f.Append(&ssa.Call{
		Func:      ssaFunc,
		Arguments: args,
		Source:    expr.Source(),
	})

	return v, nil
}