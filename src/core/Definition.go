package core

import (
	"git.urbach.dev/cli/q/src/expression"
)

// Definition compiles a define instruction.
func (f *Function) Definition(expr *expression.Expression) error {
	left := expr.Children[0]
	right := expr.Children[1]
	name := left.String(f.File.Bytes)
	value, err := f.eval(right)
	f.Identifiers[name] = value
	return err
}