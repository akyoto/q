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
	var length ssa.Value

	if addressType == types.String {
		length = addressValue.(*ssa.Struct).Arguments[1]
		addressValue = addressValue.(*ssa.Struct).Arguments[0]
		addressType = addressValue.Type()
	}

	pointer, isPointer := addressType.(*types.Pointer)

	if !isPointer {
		return nil, errors.New(&TypeNotIndexable{TypeName: addressType.Name()}, f.File, address.Source().StartPos)
	}

	_, isPointerToStruct := pointer.To.(*types.Struct)

	if isPointerToStruct {
		return nil, errors.New(&NotImplemented{Subject: "struct pointer dereferencing"}, f.File, address.Source().StartPos)
	}

	var indexValue ssa.Value

	if len(expr.Children) > 1 {
		index := expr.Children[1]

		if index.Token.Kind == token.Range {
			switch len(index.Children) {
			case 1:
				from, err := f.evaluate(index.Children[0])

				if err != nil {
					return nil, err
				}

				newPointer := f.Append(&ssa.BinaryOp{
					Op:    token.Add,
					Left:  addressValue,
					Right: from,
				})

				newLength := f.Append(&ssa.BinaryOp{
					Op:    token.Sub,
					Left:  length,
					Right: from,
				})

				slice := &ssa.Struct{
					Typ:       types.String,
					Arguments: []ssa.Value{newPointer, newLength},
					Source:    expr.Source(),
				}

				return slice, nil

			case 2:
				if index.Children[0].Token.Kind == token.Invalid {
					newLength, err := f.evaluate(index.Children[1])

					if err != nil {
						return nil, err
					}

					slice := &ssa.Struct{
						Typ:       types.String,
						Arguments: []ssa.Value{addressValue, newLength},
						Source:    expr.Source(),
					}

					return slice, nil
				}

				from, err := f.evaluate(index.Children[0])

				if err != nil {
					return nil, err
				}

				to, err := f.evaluate(index.Children[1])

				if err != nil {
					return nil, err
				}

				newPointer := f.Append(&ssa.BinaryOp{
					Op:    token.Add,
					Left:  addressValue,
					Right: from,
				})

				newLength := f.Append(&ssa.BinaryOp{
					Op:    token.Sub,
					Left:  to,
					Right: from,
				})

				slice := &ssa.Struct{
					Typ:       types.String,
					Arguments: []ssa.Value{newPointer, newLength},
					Source:    expr.Source(),
				}

				return slice, nil
			}
		} else {
			indexValue, err = f.evaluate(index)

			if err != nil {
				return nil, err
			}
		}
	} else {
		indexValue = f.Append(&ssa.Int{Int: 0})
	}

	v := f.Append(&ssa.Load{
		Typ:     pointer.To,
		Address: addressValue,
		Index:   indexValue,
		Source:  expr.Source(),
	})

	return v, nil
}