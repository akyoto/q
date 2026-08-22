package core

import (
	"slices"

	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
)

// evaluatePackageMember converts a pkg.something expression to an SSA value.
func (f *Function) evaluatePackageMember(pkg *Package, rightText string, expr *expression.Expression) (ssa.Value, error) {
	variants, exists := pkg.Functions[rightText]

	if exists {
		v := &ssa.Function{
			FunctionRef: variants,
			Typ:         variants.Type,
			Source:      expr.Source(),
		}

		return v, nil
	}

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

	enum, exists := pkg.Enums[rightText]

	if exists {
		value := &ssa.Enum{
			Typ:    enum,
			Source: expr.Source(),
		}

		return value, nil
	}

	global, exists := pkg.Globals[rightText]

	if exists {
		f.Globals.Add(global)

		v := f.Append(&ssa.Global{
			Label:       pkg.Name + "." + global.Name,
			Typ:         f.Env.Pointer(global.Typ),
			ThreadLocal: global.ThreadLocal,
			Source:      expr.Source(),
		})

		return v, nil
	}

	if pkg.Name != f.File.Package {
		rightText = pkg.Name + "." + rightText
	}

	return nil, errors.New(&UnknownIdentifier{Name: rightText}, f.File, expr.Source())
}