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
			return value, nil
		}

		_, exists = f.Env.Packages[name]

		if exists {
			return &ssa.Package{Name: name}, nil
		}

		return f.evaluateDotPackage(f.Env.Packages[f.File.Package], name, expr)

	case token.Number, token.Rune:
		number, err := toNumber(expr.Token, f.File)

		if err != nil {
			return nil, err
		}

		v := f.Append(&ssa.Int{
			Int:    number,
			Source: expr.Source(),
		})

		return v, nil

	case token.String:
		data := expr.Token.Bytes(f.File.Bytes)
		data = unescape(data)

		v := &ssa.Struct{
			Typ:    types.String,
			Source: expr.Source(),
		}

		length := f.Append(&ssa.Int{
			Int:       len(data),
			Structure: v,
			Source:    expr.Source(),
		})

		pointer := f.Append(&ssa.Bytes{
			Bytes:     data,
			Structure: v,
			Source:    expr.Source(),
		})

		v.Arguments = []ssa.Value{pointer, length}
		return v, nil
	}

	return nil, errors.New(InvalidExpression, f.File, expr.Token.Position)
}