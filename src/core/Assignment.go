package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
)

// Assignment compiles an assignment.
func (f *Function) Assignment(expr *expression.Expression) error {
	left := expr.Children[0]
	name := left.String(f.File.Bytes)
	_, exists := f.Block().FindIdentifier(name)

	if !exists {
		return errors.New(&UnknownIdentifier{Name: name}, f.File, left.Token.Position)
	}

	right := expr.Children[1]
	value, err := f.eval(right)
	f.Block().Identify(name, value)
	return err
}