package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/ssa"
)

// compileStore compiles an assignment to memory.
func (f *Function) compileStore(node *ast.Assign) error {
	left := node.Expression.Children[0]
	address := left.Children[0]
	index := left.Children[1]
	addressValue, err := f.evaluate(address)

	if err != nil {
		return err
	}

	indexValue, err := f.evaluate(index)

	if err != nil {
		return err
	}

	right := node.Expression.Children[1]
	rightValue, err := f.evaluate(right)

	if err != nil {
		return err
	}

	f.Append(&ssa.Store{
		Address: addressValue,
		Index:   indexValue,
		Value:   rightValue,
		Source:  ssa.Source(node.Expression.Source()),
	})

	return nil
}