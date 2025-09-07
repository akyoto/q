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

	return f.defineMulti(left, right, false)
}