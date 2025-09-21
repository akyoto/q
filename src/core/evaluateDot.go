package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// evaluateDot converts a dot expression to an SSA value.
func (f *Function) evaluateDot(expr *expression.Expression) (ssa.Value, error) {
	if len(expr.Children) != 2 {
		return nil, errors.New(InvalidExpression, f.File, expr.Source().StartPos)
	}

	right := expr.Children[1]
	left := expr.Children[0]
	leftValue, err := f.evaluate(left)

	if err != nil {
		return nil, err
	}

	pkgValue, isPackage := leftValue.(*ssa.Package)

	if isPackage {
		pkg := f.Env.Packages[pkgValue.Name]

		if !pkg.IsExtern && f != f.Env.Init {
			imp, exists := f.File.Imports[pkgValue.Name]

			if !exists {
				return nil, errors.New(&UnknownIdentifier{Name: pkgValue.Name}, f.File, left.Token.Position)
			}

			imp.Used.Add(1)
		}

		if right.Token.Kind != token.Identifier {
			return nil, errors.New(ExpectedPackageMember, f.File, right.Source().StartPos)
		}

		rightText := right.Token.String(f.File.Bytes)
		return f.evaluatePackageMember(pkg, rightText, expr)
	}

	if right.Token.Kind != token.Identifier {
		return nil, errors.New(ExpectedStructField, f.File, right.Source().StartPos)
	}

	rightText := right.Token.String(f.File.Bytes)

	switch leftValue := leftValue.(type) {
	case *ssa.Struct:
		field := types.Unwrap(leftValue.Typ).(*types.Struct).FieldByName(rightText)

		if field == nil {
			return nil, errors.New(&UnknownStructField{StructName: leftValue.Typ.Name(), FieldName: rightText}, f.File, right.Token.Position)
		}

		value := leftValue.Arguments[field.Index]

		if value == nil {
			return nil, errors.New(&UndefinedStructField{Identifier: left.SourceString(f.File.Bytes), FieldName: rightText}, f.File, right.Token.Position)
		}

		return value, nil

	case *ssa.Call:
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
	}

	leftUnwrapped := types.Unwrap(leftValue.Type())
	pointer, isPointer := leftUnwrapped.(*types.Pointer)

	if isPointer {
		leftUnwrapped = pointer.To
	}

	structure, isStructPointer := leftUnwrapped.(*types.Struct)

	if !isStructPointer {
		return nil, errors.New(&NotDataStruct{TypeName: leftUnwrapped.Name()}, f.File, left.Source().StartPos)
	}

	field := structure.FieldByName(rightText)
	memory := f.structField(leftValue, field)
	load := f.Append(&ssa.Load{Memory: memory})
	return load, nil
}