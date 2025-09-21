package core

import (
	"git.urbach.dev/cli/q/src/ast"
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
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

	if !types.Is(rightValue.Type(), memory.Typ) {
		return errors.New(&TypeMismatch{Encountered: rightValue.Type().Name(), Expected: memory.Typ.Name()}, f.File, right.Source().StartPos)
	}

	structure, isStructType := types.Unwrap(rightValue.Type()).(*types.Struct)

	if !isStructType {
		f.Append(&ssa.Store{
			Memory: memory,
			Value:  rightValue,
			Source: node.Expression.Source(),
		})

		return nil
	}

	composite, isStruct := rightValue.(*ssa.Struct)

	if isStruct {
		for i, field := range structure.Fields {
			fieldValue := composite.Arguments[i]

			f.Append(&ssa.Store{
				Memory: f.structField(memory, field),
				Value:  fieldValue,
				Source: node.Expression.Source(),
			})
		}

		return nil
	}

	fieldValues := make([]ssa.Value, len(structure.Fields))

	for i := range structure.Fields {
		fieldValue := &ssa.FromTuple{
			Tuple:     rightValue,
			Index:     i,
			Structure: composite,
			Source:    left.Source(),
		}

		f.Block().Append(fieldValue)
		fieldValues[i] = fieldValue
	}

	for i, field := range structure.Fields {
		f.Append(&ssa.Store{
			Memory: f.structField(memory, field),
			Value:  fieldValues[i],
			Source: node.Expression.Source(),
		})
	}

	return nil
}