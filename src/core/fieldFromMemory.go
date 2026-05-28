package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// fieldFromMemory extracts a field from memory.
func (f *Function) fieldFromMemory(leftValue ssa.Value, left *expression.Expression, right *expression.Expression, expr *expression.Expression) (ssa.Value, error) {
	fieldName := right.Token.StringFrom(f.File.Bytes)
	leftUnwrapped := types.Unwrap(leftValue.Type())
	pointer, isPointer := leftUnwrapped.(*types.Pointer)

	if isPointer {
		leftUnwrapped = pointer.To
	}

	structure, isStructPointer := leftUnwrapped.(*types.Struct)

	if !isStructPointer {
		return nil, errors.New(&NotDataStruct{TypeName: leftUnwrapped.Name()}, f.File, left.Source())
	}

	field := structure.FieldByName(fieldName)

	if field == nil {
		return nil, errors.New(&UnknownStructField{StructName: structure.Name(), FieldName: fieldName}, f.File, right.Source())
	}

	memory := f.structField(leftValue, field)
	_, memoryIsStruct := memory.Typ.(*types.Struct)

	if !memoryIsStruct {
		load := f.Append(&ssa.Load{
			Memory: memory,
			Source: expr.Source(),
		})

		return load, nil
	}

	return memory, nil
}