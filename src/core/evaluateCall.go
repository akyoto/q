package core

import (
	"strings"

	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
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
		_, isFunctionPointer := funcValue.Type().(*types.Function)

		if isFunctionPointer {
			return f.evaluateCallPointer(funcValue, expr.Source())
		}

		return nil, errors.New(&TypeMismatch{Expected: "function", Encountered: funcValue.Type().Name()}, f.File, identifier.Source())
	}

	values, err := f.evaluateAll(expr.Children[1:])

	if err != nil {
		return nil, err
	}

	args, err := f.decompose(values)

	if err != nil {
		return nil, err
	}

	variants := ssaFunc.FunctionRef.(*Function)
	fn, err := f.selectFunction(variants, values, identifier)

	if err != nil {
		return nil, err
	}

	if fn == nil {
		typeNames := make([]string, len(values))

		for i := range values {
			typeNames[i] = values[i].Type().Name()
		}

		rawName, _, _ := strings.Cut(variants.FullName, "[")
		return nil, errors.New(&NoMatchingFunction{Function: rawName, Types: strings.Join(typeNames, ",")}, f.File, identifier.Source())
	}

	ssaFunc.FunctionRef = fn
	ssaFunc.Typ = fn.Type

	for i, value := range values {
		given := value.Type()
		expected := fn.Input[i].Typ
		_, givenIsResource := given.(*types.Resource)
		_, expectedIsResource := expected.(*types.Resource)

		if givenIsResource && expectedIsResource {
			f.Block().Unidentify(value)
		}

		if types.Is(given, expected) {
			continue
		}

		typeMismatch := &TypeMismatch{
			Encountered:   given.Name(),
			Expected:      expected.Name(),
			ParameterName: fn.Input[i].Name,
		}

		return nil, errors.New(typeMismatch, f.File, expr.Children[1+i].Source())
	}

	if fn.IsExtern() {
		v := f.Append(&ssa.CallExtern{Call: ssa.Call{
			Func:      ssaFunc,
			Arguments: args,
			Source:    expr.Source(),
		}})

		f.Assembler.Libraries.Append(fn.Package(), fn.Name())
		return v, nil
	}

	if f == f.Env.Init && fn == f.Env.Main {
		f.runAll("init")
	}

	v := f.call(ssaFunc, args, expr.Source())

	if f == f.Env.Init && fn == f.Env.Main {
		f.runAll("exit")
	}

	return v, nil
}