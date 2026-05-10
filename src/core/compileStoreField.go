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
	right := node.Expression.Children[1]
	rightValue, err := f.evaluateRight(right)

	if err != nil {
		return err
	}

	left := node.Expression.Children[0]
	address := left.Children[0]
	fieldExpr := left.Children[1]
	fieldName := fieldExpr.String(f.File.Bytes)
	addressValue, err := f.evaluate(address)

	if err != nil {
		return err
	}

	if node.Expression.Token.Kind != token.Assign {
		leftValue, err := f.evaluateRight(left)

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

	switch pointer := types.Unwrap(addressValue.Type()).(type) {
	case *types.Pointer:
		field := pointer.To.(*types.Struct).FieldByName(fieldName)

		if field == nil {
			return errors.New(&UnknownStructField{StructName: pointer.To.Name(), FieldName: fieldName}, f.File, fieldExpr.Source())
		}

		if !types.Is(rightValue.Type(), field.Type) {
			return errors.New(&TypeMismatch{Encountered: rightValue.Type().Name(), Expected: field.Type.Name()}, f.File, right.Source())
		}

		memory := f.structField(addressValue, field)
		return f.store(memory, rightValue)

	case *types.Struct:
		field := pointer.FieldByName(fieldName)

		if field == nil {
			return errors.New(&UnknownStructField{StructName: pointer.Name(), FieldName: fieldName}, f.File, fieldExpr.Source())
		}

		structure, isStruct := addressValue.(*ssa.Struct)

		if isStruct {
			structure.Arguments[field.Index] = rightValue
			f.Block().Identify(address.SourceString(f.File.Bytes)+"."+fieldName, rightValue)
			return nil
		}

		memory := f.structField(addressValue, field)
		return f.store(memory, rightValue)

	default:
		panic("unknown memory store")
	}
}