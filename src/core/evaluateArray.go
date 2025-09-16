package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// evaluateArray converts a array indexing expression to an SSA value.
func (f *Function) evaluateArray(expr *expression.Expression) (ssa.Value, error) {
	address := expr.Children[0]
	addressValue, err := f.evaluate(address)

	if err != nil {
		return nil, err
	}

	addressType := types.Unwrap(addressValue.Type())
	addressStruct, addressIsStruct := addressValue.(*ssa.Struct)
	var length ssa.Value

	if addressIsStruct {
		length = addressStruct.Arguments[1]
		addressValue = addressStruct.Arguments[0]
		addressType = addressValue.Type()
	}

	pointer, isPointer := addressType.(*types.Pointer)

	if !isPointer {
		return nil, errors.New(&TypeNotIndexable{TypeName: addressType.Name()}, f.File, address.Source().StartPos)
	}

	structure, isPointerToStruct := pointer.To.(*types.Struct)

	if isPointerToStruct {
		fields := make([]ssa.Value, 0, len(structure.Fields))

		for _, field := range structure.Fields {
			fieldValue := f.Append(&ssa.Load{
				Memory: ssa.Memory{
					Address: addressValue,
					Index:   f.Append(&ssa.Int{Int: int(field.Offset)}),
					Scale:   false,
					Typ:     field.Type,
				},
				Source: expr.Source(),
			})

			fields = append(fields, fieldValue)
		}

		value := &ssa.Struct{Typ: structure, Arguments: fields}
		return value, nil
	}

	var indexValue ssa.Value

	if len(expr.Children) > 1 {
		index := expr.Children[1]

		if index.Token.Kind == token.Range {
			return f.evaluateSlice(expr, index, addressValue, length)
		} else {
			indexValue, err = f.evaluate(index)

			if err != nil {
				return nil, err
			}

			if !types.Is(indexValue.Type(), types.AnyInt) {
				return nil, errors.New(&TypeMismatch{Encountered: indexValue.Type().Name(), Expected: types.AnyInt.Name()}, f.File, index.Source().StartPos)
			}
		}
	} else {
		indexValue = f.Append(&ssa.Int{Int: 0})
	}

	v := f.Append(&ssa.Load{
		Memory: ssa.Memory{
			Address: addressValue,
			Index:   indexValue,
			Scale:   true,
			Typ:     pointer.To,
		},
		Source: expr.Source(),
	})

	return v, nil
}