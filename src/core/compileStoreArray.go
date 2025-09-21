package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/ssa"
)

// compileStoreArray compiles an assignment to an element in an array.
func (f *Function) compileStoreArray(node *ast.Assign) error {
	right := node.Expression.Children[1]
	rightValue, err := f.evaluateRight(right)

	if err != nil {
		return err
	}

	left := node.Expression.Children[0]
	leftValue, err := f.evaluateArray(left)

	if err != nil {
		return err
	}

	memory, isMemory := leftValue.(*ssa.Memory)

	if !isMemory {
		panic("not a memory address")
	}

	return f.store(memory, rightValue)
}