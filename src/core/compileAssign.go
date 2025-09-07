package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/errors"
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

	right := node.Expression.Children[1]

	if left.IsLeaf() {
		return f.define(left, right, true)
	}

	return errors.New(&NotImplemented{Subject: "multi-assignment"}, f.File, node.Expression.Source().StartPos)
}