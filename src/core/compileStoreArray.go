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
	index := left.Children[1]
	addressValue, err := f.evaluate(address)

	if err != nil {
		return err
	}

	addressType := addressValue.Type()

	if addressType == types.String {
		addressValue = addressValue.(*ssa.Struct).Arguments[0]
		addressType = addressValue.Type()
	}

	pointer, isPointer := addressType.(*types.Pointer)

	if !isPointer {
		return errors.New(&TypeNotIndexable{TypeName: addressType.Name()}, f.File, address.Source().StartPos)
	}

	indexValue, err := f.evaluate(index)

	if err != nil {
		return err
	}

	right := node.Expression.Children[1]
	rightValue, err := f.evaluate(right)

	if err != nil {
		return err
	}

	f.Append(&ssa.Store{
		Address: addressValue,
		Index:   indexValue,
		Value:   rightValue,
		Length:  uint8(pointer.To.Size()),
		Source:  ssa.Source(node.Expression.Source()),
	})

	return nil
}