package core

import (
	"git.urbach.dev/cli/q/src/ast"
)

// compileDefinition compiles a define instruction.
func (f *Function) compileDefinition(node *ast.Define) error {
	left := node.Expression.Children[0]
	right := node.Expression.Children[1]

	if left.IsLeaf() {
		return f.define(left, right, false)
	}

	rightValue, err := f.evaluate(right)

	if err != nil {
		return err
	}

	return f.compose(left, rightValue)
}