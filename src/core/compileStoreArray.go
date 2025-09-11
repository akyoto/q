package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// compileStoreArray compiles an assignment to an element in an array.
func (f *Function) compileStoreArray(node *ast.Assign) error {
	left := node.Expression.Children[0]
	address := left.Children[0]

	addressValue, err := f.evaluate(address)

	if err != nil {
		return err
	}

	addressType := types.Unwrap(addressValue.Type())
	addressStruct, addressIsStruct := addressValue.(*ssa.Struct)

	if addressIsStruct {
		addressValue = addressStruct.Arguments[0]
		addressType = addressValue.Type()
	}

	pointer, isPointer := addressType.(*types.Pointer)

	if !isPointer {
		return errors.New(&TypeNotIndexable{TypeName: addressType.Name()}, f.File, address.Source().StartPos)
	}

	right := node.Expression.Children[1]
	rightValue, err := f.evaluate(right)

	if err != nil {
		return err
	}

	if !types.Is(rightValue.Type(), pointer.To) {
		return errors.New(&TypeMismatch{Encountered: rightValue.Type().Name(), Expected: pointer.To.Name()}, f.File, right.Source().StartPos)
	}

	var indexValue ssa.Value

	if len(left.Children) > 1 {
		index := left.Children[1]
		indexValue, err = f.evaluate(index)

		if err != nil {
			return err
		}
	} else {
		indexValue = f.Append(&ssa.Int{Int: 0})
	}

	f.Append(&ssa.Store{
		Memory: ssa.Memory{
			Address: addressValue,
			Index:   indexValue,
			Scale:   true,
			Typ:     pointer.To,
		},
		Value:  rightValue,
		Source: node.Expression.Source(),
	})

	return nil
}