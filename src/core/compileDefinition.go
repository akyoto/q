package core

import "git.urbach.dev/cli/q/src/ast"

// compileDefinition compiles a define instruction.
func (f *Function) compileDefinition(node *ast.Define) error {
	left := node.Expression.Children[0]
	right := node.Expression.Children[1]

	if !left.IsLeaf() {
		return f.multiDefine(left, right)
	}

	name := left.String(f.File.Bytes)
	value, err := f.evaluate(right)
	f.Block().Identify(name, value)
	return err
}