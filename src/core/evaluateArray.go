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
	addressValue, err := f.evaluateRight(address)

	if err != nil {
		return nil, err
	}

	addressValue, addressType, length := f.decomposeSlice(addressValue)
	pointer, isPointer := types.Unwrap(addressType).(*types.Pointer)

	if !isPointer {
		return nil, errors.New(&TypeNotIndexable{TypeName: addressType.Name()}, f.File, address.Source())
	}

	var indexValue ssa.Value

	if len(expr.Children) > 1 {
		index := expr.Children[1]

		if index.Token.Kind == token.Range {
			return f.evaluateSlice(expr, index, addressValue, length)
		} else {
			indexValue, err = f.evaluateRight(index)

			if err != nil {
				return nil, err
			}

			if !types.Is(indexValue.Type(), types.AnyInt) {
				return nil, errors.New(&TypeMismatch{Encountered: indexValue.Type().Name(), Expected: types.AnyInt.Name()}, f.File, index.Source())
			}
		}
	} else {
		indexValue = f.Append(&ssa.Int{Int: 0})
	}

	memory := &ssa.Memory{
		Address: addressValue,
		Index:   indexValue,
		Scale:   true,
		Typ:     pointer.To,
		Source:  expr.Source(),
	}

	return memory, nil
}