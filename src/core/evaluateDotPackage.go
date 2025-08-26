package core

import (
	"fmt"

	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
)

// evaluateDotPackage converts a pkg.something expression to an SSA value.
func (f *Function) evaluateDotPackage(pkg *Package, rightText string, expr *expression.Expression) (ssa.Value, error) {
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
			rightText = fmt.Sprintf("%s.%s", pkg.Name, rightText)
		}

		return nil, errors.New(&UnknownIdentifier{Name: rightText}, f.File, expr.Source().StartPos)
	}

	inputExpressions := expr.Parent.Children[1:]
	fn, err := f.selectFunction(variants, inputExpressions, expr)

	if err != nil {
		return nil, err
	}

	if fn == nil {
		return nil, errors.New(&NoMatchingFunction{Function: fmt.Sprintf("%s.%s", pkg.Name, rightText)}, f.File, expr.Source().StartPos)
	}

	if fn.IsExtern() {
		f.Assembler.Libraries.Append(fn.Package, fn.Name)
	} else {
		f.Dependencies.Add(fn)
	}

	v := &ssa.Function{
		Package:  fn.Package,
		Name:     fn.Name,
		Typ:      fn.Type,
		IsExtern: fn.IsExtern(),
		Source:   expr.Source(),
	}

	return v, nil
}