package core

import (
	"fmt"

	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
)

// evaluateDot converts a dot expression to an SSA value.
func (f *Function) evaluateDot(expr *expression.Expression) (ssa.Value, error) {
	left := expr.Children[0]
	right := expr.Children[1]
	leftText := left.String(f.File.Bytes)
	rightText := right.String(f.File.Bytes)
	leftValue, err := f.evaluate(left)

	if err != nil {
		return nil, err
	}

	switch leftValue := leftValue.(type) {
	case *ssa.Package:
		pkg := f.All.Packages[leftText]

		if !pkg.IsExtern && f != f.All.Init {
			imp, exists := f.File.Imports[leftText]

			if !exists {
				return nil, errors.New(&UnknownIdentifier{Name: leftText}, f.File, left.Token.Position)
			}

			imp.Used.Add(1)
		}

		function, exists := pkg.Functions[rightText]

		if !exists {
			return nil, errors.New(&UnknownIdentifier{Name: fmt.Sprintf("%s.%s", pkg.Name, rightText)}, f.File, left.Token.Position)
		}

		if function.IsExtern() {
			f.Assembler.Libraries.Append(function.Package, function.Name)
		} else {
			f.Dependencies.Add(function)
		}

		v := &ssa.Function{
			Package:  function.Package,
			Name:     function.Name,
			Typ:      function.Type,
			IsExtern: function.IsExtern(),
			Source:   ssa.Source(expr.Source()),
		}

		return v, nil

	case *ssa.Struct:
		field := leftValue.Typ.FieldByName(rightText)

		if field == nil {
			panic("unknown field")
		}

		return f.Append(leftValue.Arguments[field.Index]), nil

	default:
		panic("not implemented")
	}
}