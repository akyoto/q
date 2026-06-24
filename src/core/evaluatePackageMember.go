package core

import (
	"slices"

	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// evaluatePackageMember converts a pkg.something expression to an SSA value.
func (f *Function) evaluatePackageMember(pkg *Package, rightText string, expr *expression.Expression) (ssa.Value, error) {
	constant, exists := pkg.Constants[rightText]

	if exists {
		if slices.Contains(f.constantStack, constant) {
			return nil, errors.New(&CycleDetected{A: f.constantStack[len(f.constantStack)-1].Name, B: constant.Name}, f.File, expr.Source())
		}

		f.constantStack = append(f.constantStack, constant)
		tmp := f.File
		f.File = constant.File
		v, err := f.evaluateRight(constant.Value)
		f.File = tmp
		f.constantStack = f.constantStack[:len(f.constantStack)-1]
		return v, err
	}

	variants, exists := pkg.Functions[rightText]

	if !exists {
		global, exists := pkg.Globals[rightText]

		if exists {
			v := f.Append(&ssa.Global{
				Label:  pkg.Name + "." + global.Name,
				Typ:    f.Env.Pointer(global.Typ),
				Source: expr.Source(),
			})

			return v, nil
		}

		if pkg.Name != f.File.Package {
			rightText = pkg.Name + "." + rightText
		}

		return nil, errors.New(&UnknownIdentifier{Name: rightText}, f.File, expr.Source())
	}

	if expr.Parent.Token.Kind == token.Call && expr.Parent.Children[0] == expr {
		inputExpressions := expr.Parent.Children[1:]
		fn, err := f.selectFunction(variants, inputExpressions, expr)

		if err != nil {
			return nil, err
		}

		if fn == nil {
			return nil, errors.New(&NoMatchingFunction{Function: pkg.Name + "." + rightText}, f.File, expr.Source())
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