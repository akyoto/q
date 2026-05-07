package core

import (
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// evaluateNewStruct allocates a new struct with a field initializer list.
func (f *Function) evaluateNewStruct(expr *expression.Expression) (ssa.Value, error) {
	pointer, err := f.evaluateNew(expr.Children[0])

	if err != nil {
		return nil, err
	}

	structType := types.Unwrap(pointer.Type()).(*types.Pointer).To.(*types.Struct)

	for _, definition := range expr.Children[1:] {
		field, rightValue, err := f.extractField(structType, definition)

		if err != nil {
			return nil, err
		}

		offset := f.Append(&ssa.Int{Int: int(field.Offset)})

		memory := &ssa.Memory{
			Address: pointer,
			Index:   offset,
			Scale:   false,
			Typ:     field.Type,
		}

		err = f.store(memory, rightValue)

		if err != nil {
			return nil, err
		}
	}

	return pointer, nil
}