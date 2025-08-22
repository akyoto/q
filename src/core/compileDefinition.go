package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
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

	rightValue, err := f.evaluate(right)

	if err != nil {
		return err
	}

	// If the value we got was a value that is stored in a variable,
	// it must have been returned from the optimizer as a cached value.
	// We want to assure that every named variable creates a copy of
	// another named variable instead of using the cached value itself
	// because it could lead to incorrect optimizations.
	if f.IsIdentified(rightValue) {
		rightValue = f.copy(rightValue, ssa.Source(right.Source()))
	}

	_, isCall := rightValue.(*ssa.Call)

	if isCall {
		structure, isStructType := rightValue.Type().(*types.Struct)

		if isStructType {
			composite := &ssa.Struct{
				Typ:       structure,
				Arguments: make(ssa.Arguments, 0, len(structure.Fields)),
				Source:    ssa.Source(left.Source()),
			}

			for i := range structure.Fields {
				field := &ssa.FromTuple{
					Tuple:     rightValue,
					Index:     i,
					Structure: composite,
					Source:    ssa.Source(left.Source()),
				}

				f.Block().Append(field)
				composite.Arguments = append(composite.Arguments, field)
			}

			f.Block().Identify(name, composite)
			return nil
		}
	}

	f.Block().Identify(name, rightValue)
	return nil
}