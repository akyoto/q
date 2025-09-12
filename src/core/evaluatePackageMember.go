package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// evaluatePackageMember converts a pkg.something expression to an SSA value.
func (f *Function) evaluatePackageMember(pkg *Package, rightText string, expr *expression.Expression) (ssa.Value, error) {
	constant, exists := pkg.Constants[rightText]

	if exists {
		number, err := toNumber(constant.Value.Token, constant.File)

		if err != nil {
			return nil, err
		}

		v := f.Append(&ssa.Int{
			Int:    number,
			Source: expr.Source(),
		})

		return v, nil
	}

	variants, exists := pkg.Functions[rightText]

	if !exists {
		if pkg.Name != f.File.Package {
			rightText = pkg.Name + "." + rightText
		}

		return nil, errors.New(&UnknownIdentifier{Name: rightText}, f.File, expr.Source().StartPos)
	}

	if expr.Parent.Token.Kind == token.Call && expr.Parent.Children[0] == expr {
		inputExpressions := expr.Parent.Children[1:]
		fn, err := f.selectFunction(variants, inputExpressions, expr)

		if err != nil {
			return nil, err
		}

		if fn == nil {
			return nil, errors.New(&NoMatchingFunction{Function: pkg.Name + "." + rightText}, f.File, expr.Source().StartPos)
		}

		if fn.IsExtern() {
			f.Assembler.Libraries.Append(fn.Package(), fn.Name())
		} else {
			f.Dependencies.Add(fn)
		}

		v := &ssa.Function{
			FunctionRef: fn,
			Typ:         fn.Type,
			Source:      expr.Source(),
		}

		return v, nil
	}

	v := f.Append(&ssa.Function{
		FunctionRef: variants,
		Typ:         variants.Type,
		Source:      expr.Source(),
	})

	f.Dependencies.Add(variants)
	return v, nil
}