package core

import (
	"fmt"

	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// eval converts an expression to an SSA value.
func (f *Function) eval(expr *expression.Expression) (ssa.Value, error) {
	if expr.IsLeaf() {
		switch expr.Token.Kind {
		case token.Identifier:
			name := expr.Token.String(f.File.Bytes)
			value, exists := f.Identifiers[name]

			if exists {
				return value, nil
			}

			_, exists = f.All.Packages[name]

			if exists {
				return &ssa.Package{Name: name}, nil
			}

			function := f.All.Function(f.File.Package, name)

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
			number, err := f.toNumber(expr.Token)

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

	switch expr.Token.Kind {
	case token.Call:
		if expr.Children[0].Token.Kind == token.Identifier && expr.Children[0].String(f.File.Bytes) == "syscall" {
			args, err := f.decompose(expr.Children[1:], nil)

			if err != nil {
				return nil, err
			}

			syscall := &ssa.Syscall{
				Arguments: args,
				Source:    ssa.Source(expr.Source()),
			}

			return f.Append(syscall), nil
		}

		funcValue, err := f.eval(expr.Children[0])

		if err != nil {
			return nil, err
		}

		ssaFunc := funcValue.(*ssa.Function)
		pkg := f.All.Packages[ssaFunc.Package]
		fn := pkg.Functions[ssaFunc.Name]
		inputExpressions := expr.Children[1:]

		if len(inputExpressions) != len(fn.Input) {
			return nil, errors.New(&ParameterCountMismatch{Function: fn.FullName, Count: len(inputExpressions), ExpectedCount: len(fn.Input)}, f.File, expr.Source().StartPos)
		}

		args, err := f.decompose(inputExpressions, fn.Input)

		if err != nil {
			return nil, err
		}

		if fn.IsExtern() {
			v := f.Append(&ssa.CallExtern{Call: ssa.Call{
				Func:      ssaFunc,
				Arguments: args,
				Source:    ssa.Source(expr.Source()),
			}})

			return v, nil
		}

		v := f.Append(&ssa.Call{
			Func:      ssaFunc,
			Arguments: args,
			Source:    ssa.Source(expr.Source()),
		})

		return v, nil

	case token.Dot:
		left := expr.Children[0]
		right := expr.Children[1]
		leftText := left.String(f.File.Bytes)
		rightText := right.String(f.File.Bytes)
		leftValue, err := f.eval(left)

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

				imp.Used = true
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

	default:
		if expr.Token.Kind.IsOperator() {
			left := expr.Children[0]
			right := expr.Children[1]

			leftValue, err := f.eval(left)

			if err != nil {
				return nil, err
			}

			rightValue, err := f.eval(right)

			if err != nil {
				return nil, err
			}

			v := f.Append(&ssa.BinaryOp{
				Left:   leftValue,
				Right:  rightValue,
				Op:     expr.Token.Kind,
				Source: ssa.Source(expr.Source()),
			})

			return v, nil
		}

		panic("not implemented")
	}
}