package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// compileAssign compiles an assignment.
func (f *Function) compileAssign(node *ast.Assign) error {
	left := node.Expression.Children[0]

	if left.Token.Kind == token.Array {
		return f.compileStoreArray(node)
	}

	if left.Token.Kind == token.Dot {
		return f.compileStoreField(node)
	}

	name := left.String(f.File.Bytes)
	leftValue, exists := f.Block().FindIdentifier(name)

	if !exists {
		return errors.New(&UnknownIdentifier{Name: name}, f.File, left.Source().StartPos)
	}

	right := node.Expression.Children[1]
	rightValue, err := f.evaluate(right)

	if err != nil {
		return err
	}

	if f.IsIdentified(rightValue) {
		rightValue = f.copy(rightValue, right.Source())
	}

	if node.Expression.Token.Kind == token.Assign {
		f.Block().Identify(name, rightValue)
		return nil
	}

	operator := removeAssign(node.Expression.Token.Kind)

	operation := f.Append(&ssa.BinaryOp{
		Op:     operator,
		Left:   leftValue,
		Right:  rightValue,
		Source: node.Expression.Source(),
	})

	f.Block().Identify(name, operation)
	return nil
}