package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/ssa"
)

// compileDefinition compiles a define instruction.
func (f *Function) compileDefinition(node *ast.Define) error {
	left := node.Expression.Children[0]
	right := node.Expression.Children[1]
	rightValue, err := f.evaluate(right)

	if err != nil {
		return err
	}

	if left.IsLeaf() {
		return f.define(left, rightValue)
	}

	count := 0

	for leaf := range left.Leaves() {
		value := f.Append(&ssa.FromTuple{
			Tuple:  rightValue,
			Index:  count,
			Source: leaf.Source(),
		})

		err = f.define(leaf, value)

		if err != nil {
			return err
		}

		count++
	}

	fn := rightValue.(*ssa.Call).Func

	if count != len(fn.Typ.Output) {
		return errors.New(&DefinitionCountMismatch{Function: fn.String(), Count: count, ExpectedCount: len(fn.Typ.Output)}, f.File, left.Source().StartPos)
	}

	return nil
}