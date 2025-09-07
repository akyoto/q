package core

import (
	"slices"

	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
)

// evaluateIdentifier converts an identifier to an SSA value.
func (f *Function) evaluateIdentifier(expr *expression.Expression) (ssa.Value, error) {
	name := expr.Token.String(f.File.Bytes)

	switch name {
	case "false":
		v := f.Append(&ssa.Bool{
			Bool:   false,
			Source: expr.Source(),
		})

		return v, nil

	case "true":
		v := f.Append(&ssa.Bool{
			Bool:   true,
			Source: expr.Source(),
		})

		return v, nil
	}

	value, exists := f.Block().FindIdentifier(name)

	if exists {
		for _, p := range f.Block().Protected {
			if slices.Contains(p, value) {
				return nil, errors.New(&ErrorNotChecked{Identifier: name}, f.File, expr.Token.Position)
			}
		}

		return value, nil
	}

	if name != f.File.Package {
		_, exists := f.Env.Packages[name]

		if exists {
			return &ssa.Package{Name: name}, nil
		}
	}

	return f.evaluatePackageMember(f.Env.Packages[f.File.Package], name, expr)
}