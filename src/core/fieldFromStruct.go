package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// fieldFromStruct extracts a field from a struct.
func (f *Function) fieldFromStruct(leftValue *ssa.Struct, left *expression.Expression, right *expression.Expression) (ssa.Value, error) {
	fieldName := right.Token.StringFrom(f.File.Bytes)
	structType := types.Unwrap(leftValue.Typ).(*types.Struct)
	field := structType.FieldByName(fieldName)

	if field == nil {
		return nil, errors.New(&UnknownStructField{StructName: leftValue.Typ.Name(), FieldName: fieldName}, f.File, right.Source())
	}

	arg := leftValue.Arguments[field.Index]

	if arg == nil {
		return nil, errors.New(&UndefinedStructField{Identifier: left.SourceString(f.File.Bytes), FieldName: fieldName}, f.File, right.Source())
	}

	return arg, nil
}