package core

import (
	"git.urbach.dev/cli/q/src/expression"
)

// Assignment compiles an assignment.
func (f *Function) Assignment(expr *expression.Expression) error {
	return f.Definition(expr)
}