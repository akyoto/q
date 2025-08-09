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
	name := left.String(f.File.Bytes)
	leftValue, exists := f.Block().FindIdentifier(name)

	if !exists {
		return errors.New(&UnknownIdentifier{Name: name}, f.File, left.Token.Position)
	}

	right := node.Expression.Children[1]
	rightValue, err := f.evaluate(right)

	if err != nil {
		return err
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
		Source: ssa.Source(node.Expression.Source()),
	})

	f.Block().Identify(name, operation)
	return nil
}