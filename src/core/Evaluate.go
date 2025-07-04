package core

import (
	"fmt"
	"slices"

	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// Evaluate converts an expression to an SSA value.
func (f *Function) Evaluate(expr *expression.Expression) (ssa.Value, error) {
	if expr.IsLeaf() {
		switch expr.Token.Kind {
		case token.Identifier:
			name := expr.Token.String(f.File.Bytes)
			value, exists := f.Identifiers[name]

			if !exists {
				function, exists := f.All.Functions[f.File.Package+"."+name]

				if !exists {
					return nil, errors.New(&UnknownIdentifier{Name: name}, f.File, expr.Token.Position)
				}

				f.Dependencies.Add(function)

				v := f.Append(&ssa.Function{
					UniqueName: function.UniqueName,
					Typ:        function.Type,
					IsExtern:   function.IsExtern(),
					Source:     ssa.Source(expr.Source),
				})

				return v, nil
			}

			return value, nil

		case token.Number:
			number, err := f.ToNumber(expr.Token)

			if err != nil {
				return nil, err
			}

			v := f.Append(&ssa.Int{
				Int:    number,
				Source: ssa.Source(expr.Source),
			})

			return v, nil

		case token.String:
			data := expr.Token.Bytes(f.File.Bytes)
			data = Unescape(data)

			length := f.Append(&ssa.Int{
				Int:    len(data),
				Source: ssa.Source(expr.Source),
			})

			pointer := f.Append(&ssa.Bytes{
				Bytes:  data,
				Source: ssa.Source(expr.Source),
			})

			v := f.Append(&ssa.Struct{
				Arguments: []ssa.Value{pointer, length},
				Typ:       types.String,
				Source:    ssa.Source(expr.Source),
			})

			return v, nil
		}

		return nil, errors.New(InvalidExpression, f.File, expr.Token.Position)
	}

	switch expr.Token.Kind {
	case token.Call:
		children := expr.Children
		isSyscall := false

		if children[0].Token.Kind == token.Identifier {
			funcName := children[0].String(f.File.Bytes)

			if funcName == "syscall" {
				children = children[1:]
				isSyscall = true
			}
		}

		args := make([]ssa.Value, len(children))

		for i, child := range slices.Backward(children) {
			value, err := f.Evaluate(child)

			if err != nil {
				return nil, err
			}

			args[i] = value
		}

		if isSyscall {
			syscall := &ssa.Syscall{
				Arguments: args,
				Source:    ssa.Source(expr.Source),
			}

			return f.Append(syscall), nil
		}

		name := args[0].(*ssa.Function).UniqueName
		fn := f.All.Functions[name]
		parameters := args[1:]

		if len(parameters) != len(fn.Input) {
			return nil, errors.New(&ParameterCountMismatch{Function: name, Count: len(parameters), ExpectedCount: len(fn.Input)}, f.File, expr.Source[0].Position)
		}

		for i, param := range slices.Backward(parameters) {
			if !types.Is(param.Type(), fn.Input[i].Typ) {
				_, isPointer := fn.Input[i].Typ.(*types.Pointer)

				if isPointer {
					number, isInt := param.(*ssa.Int)

					if isInt && number.Int == 0 {
						continue
					}
				}

				// Temporary hack to allow int64 -> uint32 conversion
				if types.Is(param.Type(), types.AnyInt) && types.Is(fn.Input[i].Typ, types.AnyInt) {
					continue
				}

				return nil, errors.New(&TypeMismatch{
					Encountered:   param.Type().Name(),
					Expected:      fn.Input[i].Typ.Name(),
					ParameterName: fn.Input[i].Name,
				}, f.File, param.(ssa.HasSource).Start())
			}
		}

		if fn.IsExtern() {
			v := f.Append(&ssa.CallExtern{Call: ssa.Call{
				Arguments: args,
				Source:    ssa.Source(expr.Source),
			}})

			return v, nil
		}

		v := f.Append(&ssa.Call{
			Arguments: args,
			Source:    ssa.Source(expr.Source),
		})

		return v, nil

	case token.Dot:
		left := expr.Children[0]
		right := expr.Children[1]
		leftText := left.String(f.File.Bytes)
		rightText := right.String(f.File.Bytes)
		fullName := fmt.Sprintf("%s.%s", leftText, rightText)
		identifier, exists := f.Identifiers[fullName]

		if exists {
			return identifier, nil
		}

		// identifier, exists := f.Identifiers[leftText]

		// if exists {
		// 	structType := identifier.Type().(*types.Struct)
		// 	field := structType.FieldByName(rightText)

		// 	if field == nil {
		// 		return nil, errors.New(&UnknownStructField{StructName: structType.Name(), FieldName: rightText}, f.File, right.Token.Position)
		// 	}

		// 	v := f.Append(&ssa.Field{
		// 		Object: identifier,
		// 		Field:  field,
		// 		Source: ssa.Source(expr.Source),
		// 	})

		// 	return v, nil
		// }

		function, exists := f.All.Functions[fullName]

		if exists {
			if function.IsExtern() {
				f.Assembler.Libraries = f.Assembler.Libraries.Append(function.Package, function.Name)
			} else {
				f.Dependencies.Add(function)
			}

			v := f.Append(&ssa.Function{
				UniqueName: function.UniqueName,
				Typ:        function.Type,
				IsExtern:   function.IsExtern(),
				Source:     ssa.Source(expr.Source),
			})

			return v, nil
		}

		return nil, errors.New(&UnknownIdentifier{Name: fullName}, f.File, left.Token.Position)

	default:
		if expr.Token.IsOperator() {
			left := expr.Children[0]
			right := expr.Children[1]

			leftValue, err := f.Evaluate(left)

			if err != nil {
				return nil, err
			}

			rightValue, err := f.Evaluate(right)

			if err != nil {
				return nil, err
			}

			v := f.Append(&ssa.BinaryOp{
				Left:   leftValue,
				Right:  rightValue,
				Op:     expr.Token.Kind,
				Source: ssa.Source(expr.Source),
			})

			return v, nil
		}

		panic("not implemented")
	}
}