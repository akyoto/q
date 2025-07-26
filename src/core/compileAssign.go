package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/errors"
)

// compileAssign compiles an assignment.
func (f *Function) compileAssign(node *ast.Assign) error {
	left := node.Expression.Children[0]
	name := left.String(f.File.Bytes)
	_, exists := f.Block().FindIdentifier(name)

	if !exists {
		return errors.New(&UnknownIdentifier{Name: name}, f.File, left.Token.Position)
	}

	right := node.Expression.Children[1]
	value, err := f.evaluate(right)
	f.Block().Identify(name, value)
	return err
}