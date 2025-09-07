package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/token"
)

// compileAssign compiles an assignment.
func (f *Function) compileAssign(node *ast.Assign) error {
	left := node.Expression.Children[0]
	right := node.Expression.Children[1]

	if left.Token.Kind == token.Array {
		return f.compileStoreArray(node)
	}

	if left.Token.Kind == token.Dot {
		return f.compileStoreField(node)
	}

	if left.IsLeaf() {
		return f.define(left, right, true)
	}

	return f.defineMulti(left, right, true)
}