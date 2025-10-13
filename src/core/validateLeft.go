package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// validateLeft validates the left side in a definition or an assignment.
// For a definition it expects that the name does not exist yet.
// For an assignment it expects that the name exists and the type matches.
func (f *Function) validateLeft(left *expression.Expression, right *expression.Expression, name string, rightType types.Type, isAssign bool) (ssa.Value, error) {
	leftValue, exists := f.Block().FindIdentifier(name)

	if isAssign {
		if !exists {
			name := left.Token.StringFrom(f.File.Bytes)
			pkg := f.Env.Packages[f.File.Package]
			global, isGlobal := pkg.Globals[name]

			if isGlobal {
				v := f.Append(&ssa.Data{
					Label:  f.File.Package + "." + global.Name,
					Typ:    f.Env.Pointer(global.Typ),
					Source: left.Source(),
				})

				return v, nil
			}

			return nil, errors.New(&UnknownIdentifier{Name: name}, f.File, left.Source())
		}

		phi, isPhi := leftValue.(*ssa.Phi)

		if isPhi && phi.IsPartiallyUndefined() {
			return nil, errors.New(&PartiallyUnknownIdentifier{Name: name}, f.File, left.Source())
		}

		if !types.Is(rightType, leftValue.Type()) {
			return nil, errors.New(&TypeMismatch{Encountered: rightType.Name(), Expected: leftValue.Type().Name()}, f.File, right.Source())
		}

		resource, leftIsResource := leftValue.Type().(*types.Resource)

		if leftIsResource {
			return nil, errors.New(&ResourceNotConsumed{TypeName: resource.Name()}, f.File, left.Source())
		}
	} else if exists {
		return nil, errors.New(&VariableAlreadyExists{Name: name}, f.File, left.Source())
	}

	return leftValue, nil
}