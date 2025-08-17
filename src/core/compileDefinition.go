package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/errors"
)

// compileDefinition compiles a define instruction.
func (f *Function) compileDefinition(node *ast.Define) error {
	left := node.Expression.Children[0]
	right := node.Expression.Children[1]

	if !left.IsLeaf() {
		return f.multiDefine(left, right)
	}

	name := left.String(f.File.Bytes)
	_, exists := f.Block().FindIdentifier(name)

	if exists {
		return errors.New(&VariableAlreadyExists{Name: name}, f.File, left.Source().StartPos)
	}

	value, err := f.evaluate(right)
	f.Block().Identify(name, value)
	return err
}