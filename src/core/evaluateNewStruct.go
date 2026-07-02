package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// evaluateNewStruct allocates a new struct with a field initializer list.
func (f *Function) evaluateNewStruct(expr *expression.Expression) (ssa.Value, error) {
	object, err := f.evaluateNew(expr.Children[0])

	if err != nil {
		return nil, err
	}

	switch typ := types.Unwrap(object.Type()).(type) {
	case *types.Pointer:
		structType, isStructType := typ.To.(*types.Struct)

		if !isStructType {
			return nil, errors.New(&NotDataStruct{TypeName: typ.To.Name()}, f.File, expr.Children[0].Children[1].Source())
		}

		for _, definition := range expr.Children[1:] {
			if isTrailing(definition, expr.Children) {
				continue
			}

			field, rightValue, err := f.extractField(structType, definition)

			if err != nil {
				return nil, err
			}

			memory := f.structField(object, field)
			err = f.store(memory, rightValue)

			if err != nil {
				return nil, err
			}
		}

	case *types.Struct:
		return nil, errors.New(&NotImplemented{Subject: "struct initialization for arrays"}, f.File, expr.Source())

	default:
		panic("invalid struct initialization type")
	}

	return object, nil
}