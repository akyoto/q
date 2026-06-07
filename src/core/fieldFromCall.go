package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// fieldFromCall extracts a field from a function call.
func (f *Function) fieldFromCall(call *ssa.Call, left *expression.Expression, right *expression.Expression, expr *expression.Expression) (ssa.Value, error) {
	fieldName := right.Token.StringFrom(f.File.Bytes)
	leftUnwrapped := types.Unwrap(call.Type())
	structure, isStruct := leftUnwrapped.(*types.Struct)

	if isStruct {
		field := structure.FieldByName(fieldName)

		if field == nil {
			return nil, errors.New(&UnknownStructField{StructName: structure.Name(), FieldName: fieldName}, f.File, right.Source())
		}

		value := f.Append(&ssa.Field{
			Tuple:  call,
			Index:  int(field.Index),
			Source: expr.Source(),
		})

		return value, nil
	}

	pointer, isPointer := leftUnwrapped.(*types.Pointer)

	if isPointer {
		leftUnwrapped = pointer.To
	}

	_, isStructPointer := leftUnwrapped.(*types.Struct)

	if isStructPointer {
		return f.fieldFromMemory(call, left, right, expr)
	}

	return nil, errors.New(&NotDataStruct{TypeName: leftUnwrapped.Name()}, f.File, call.Source)
}