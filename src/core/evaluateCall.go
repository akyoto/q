package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
)

// evaluateCall converts a call expression to an SSA value.
func (f *Function) evaluateCall(expr *expression.Expression) (ssa.Value, error) {
	identifier := expr.Children[0]

	if identifier.Token.Kind.IsBuiltin() {
		return f.evaluateBuiltin(expr)
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