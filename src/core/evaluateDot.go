package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// evaluateDot converts a dot expression to an SSA value.
func (f *Function) evaluateDot(expr *expression.Expression) (ssa.Value, error) {
	if len(expr.Children) != 2 {
		return nil, errors.New(InvalidExpression, f.File, expr.Source().StartPos)
	}

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
		pkg := f.Env.Packages[leftText]

		if !pkg.IsExtern && f != f.Env.Init {
			imp, exists := f.File.Imports[leftText]

			if !exists {
				return nil, errors.New(&UnknownIdentifier{Name: leftText}, f.File, left.Token.Position)
			}

			imp.Used.Add(1)
		}

		return f.evaluatePackageMember(pkg, rightText, expr)

	case *ssa.Struct:
		field := types.Unwrap(leftValue.Typ).(*types.Struct).FieldByName(rightText)

		if field == nil {
			return nil, errors.New(&UnknownStructField{StructName: leftValue.Typ.Name(), FieldName: rightText}, f.File, right.Token.Position)
		}

		value := leftValue.Arguments[field.Index]

		if value == nil {
			return nil, errors.New(&UndefinedStructField{Identifier: leftText, FieldName: rightText}, f.File, right.Token.Position)
		}

		return value, nil

	default:
		leftUnwrapped := types.Unwrap(leftValue.Type())
		structure, isStruct := leftUnwrapped.(*types.Struct)

		if isStruct {
			field := structure.FieldByName(rightText)

			value := f.Append(&ssa.FromTuple{
				Tuple:  leftValue,
				Index:  int(field.Index),
				Source: left.Source(),
			})

			return value, nil
		}

		pointer, isPointer := leftUnwrapped.(*types.Pointer)

		if !isPointer {
			return nil, errors.New(&NotDataStruct{TypeName: leftValue.Type().Name()}, f.File, left.Source().StartPos)
		}

		structure, isStructPointer := pointer.To.(*types.Struct)

		if !isStructPointer {
			return nil, errors.New(&NotDataStruct{TypeName: pointer.To.Name()}, f.File, left.Source().StartPos)
		}

		field := structure.FieldByName(rightText)
		offset := f.Append(&ssa.Int{Int: int(field.Offset)})

		load := f.Append(&ssa.Load{
			Typ:     field.Type,
			Address: leftValue,
			Index:   offset,
		})

		return load, nil
	}
}