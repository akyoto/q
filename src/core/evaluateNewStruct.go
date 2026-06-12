package core

import (
	"git.urbach.dev/cli/q/src/errors"
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

	typ := types.Unwrap(pointer.Type()).(*types.Pointer).To
	structType, isStructType := typ.(*types.Struct)

	if !isStructType {
		return nil, errors.New(&NotDataStruct{TypeName: typ.Name()}, f.File, expr.Children[0].Children[1].Source())
	}

	for _, definition := range expr.Children[1:] {
		if isTrailing(definition, expr.Children) {
			continue
		}

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