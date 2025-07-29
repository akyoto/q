package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// evaluateLeaf converts a leaf expression to an SSA value.
func (f *Function) evaluateLeaf(expr *expression.Expression) (ssa.Value, error) {
	switch expr.Token.Kind {
	case token.Identifier:
		name := expr.Token.String(f.File.Bytes)
		value, exists := f.Block().FindIdentifier(name)

		if exists {
			return value, nil
		}

		constant, exists := f.Env.Packages[f.Package].Constants[name]

		if exists {
			number, err := toNumber(constant.Value.Token, constant.File)

			if err != nil {
				return nil, err
			}

			v := f.Append(&ssa.Int{
				Int:    number,
				Source: ssa.Source(expr.Source()),
			})

			return v, nil
		}

		_, exists = f.Env.Packages[name]

		if exists {
			return &ssa.Package{Name: name}, nil
		}

		function := f.Env.Function(f.File.Package, name)

		if function != nil {
			f.Dependencies.Add(function)

			v := &ssa.Function{
				Package:  function.Package,
				Name:     function.Name,
				Typ:      function.Type,
				IsExtern: function.IsExtern(),
				Source:   ssa.Source(expr.Source()),
			}

			return v, nil
		}

		return nil, errors.New(&UnknownIdentifier{Name: name}, f.File, expr.Token.Position)

	case token.Number, token.Rune:
		number, err := toNumber(expr.Token, f.File)

		if err != nil {
			return nil, err
		}

		v := f.Append(&ssa.Int{
			Int:    number,
			Source: ssa.Source(expr.Source()),
		})

		return v, nil

	case token.String:
		data := expr.Token.Bytes(f.File.Bytes)
		data = unescape(data)

		v := &ssa.Struct{
			Typ:    types.String,
			Source: ssa.Source(expr.Source()),
		}

		length := f.Append(&ssa.Int{
			Int:       len(data),
			Structure: v,
			Source:    ssa.Source(expr.Source()),
		})

		pointer := f.Append(&ssa.Bytes{
			Bytes:     data,
			Structure: v,
			Source:    ssa.Source(expr.Source()),
		})

		v.Arguments = []ssa.Value{pointer, length}
		return v, nil
	}

	return nil, errors.New(InvalidExpression, f.File, expr.Token.Position)
}