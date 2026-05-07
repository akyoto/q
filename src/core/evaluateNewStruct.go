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

		memory := f.structField(pointer, field)
		err = f.store(memory, rightValue)

		if err != nil {
			return nil, err
		}
	}

	return pointer, nil
}