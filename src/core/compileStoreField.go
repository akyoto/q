package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// compileStoreField compiles an assignment to a struct field.
func (f *Function) compileStoreField(node *ast.Assign) error {
	left := node.Expression.Children[0]
	address := left.Children[0]
	fieldName := left.Children[1].String(f.File.Bytes)
	addressValue, err := f.evaluate(address)

	if err != nil {
		return err
	}

	field := addressValue.Type().(*types.Pointer).To.(*types.Struct).FieldByName(fieldName)
	offset := f.Append(&ssa.Int{Int: int(field.Offset)})
	right := node.Expression.Children[1]
	rightValue, err := f.evaluate(right)

	if err != nil {
		return err
	}

	if node.Expression.Token.Kind != token.Assign {
		leftValue, err := f.evaluate(left)

		if err != nil {
			return err
		}

		rightValue = f.Append(&ssa.BinaryOp{
			Op:     removeAssign(node.Expression.Token.Kind),
			Left:   leftValue,
			Right:  rightValue,
			Source: ssa.Source(node.Expression.Source()),
		})
	}

	f.Append(&ssa.Store{
		Address: addressValue,
		Index:   offset,
		Value:   rightValue,
		Length:  uint8(field.Type.Size()),
		Source:  ssa.Source(node.Expression.Source()),
	})

	return nil
}