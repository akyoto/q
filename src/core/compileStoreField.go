package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/errors"
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
			Source: node.Expression.Source(),
		})
	}

	switch pointer := addressValue.Type().(type) {
	case *types.Pointer:
		field := pointer.To.(*types.Struct).FieldByName(fieldName)

		if field == nil {
			return errors.New(&UnknownStructField{StructName: pointer.To.Name(), FieldName: fieldName}, f.File, left.Children[1].Source().StartPos)
		}

		if !types.Is(rightValue.Type(), field.Type) {
			return errors.New(&TypeMismatch{Encountered: rightValue.Type().Name(), Expected: field.Type.Name()}, f.File, right.Source().StartPos)
		}

		offset := f.Append(&ssa.Int{Int: int(field.Offset)})

		f.Append(&ssa.Store{
			Memory: ssa.Memory{
				Address: addressValue,
				Index:   offset,
				Scale:   false,
				Typ:     field.Type,
			},
			Value:  rightValue,
			Source: node.Expression.Source(),
		})

	case *types.Struct:
		field := pointer.FieldByName(fieldName)
		addressValue.(*ssa.Struct).Arguments[field.Index] = rightValue
		f.Block().Identify(address.String(f.File.Bytes)+"."+fieldName, rightValue)
	}

	return nil
}