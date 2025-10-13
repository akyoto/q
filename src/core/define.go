package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// define gives a value an identifier.
func (f *Function) define(left *expression.Expression, right *expression.Expression, isAssign bool) error {
	rightValue, err := f.evaluateRight(right)

	if err != nil {
		return err
	}

	name := left.String(f.File.Bytes)

	if name == "_" {
		return nil
	}

	leftValue, err := f.validateLeft(left, right, name, rightValue.Type(), isAssign)

	if err != nil {
		return err
	}

	data, isData := leftValue.(*ssa.Data)

	if isData {
		zero := f.Append(&ssa.Int{Int: 0})

		f.Append(&ssa.Store{
			Memory: &ssa.Memory{
				Typ:     data.Typ.(*types.Pointer).To,
				Address: data,
				Index:   zero,
				Source:  data.Source,
			},
			Value: rightValue,
		})

		return nil
	}

	call, isCall := rightValue.(*ssa.Call)

	if isCall && len(call.Func.Typ.Output) != 1 {
		return errors.New(&DefinitionCountMismatch{Function: call.Func.String(), Count: 1, ExpectedCount: len(call.Func.Typ.Output)}, f.File, left.Source())
	}

	// If the value we got was a value that is stored in a variable,
	// it must have been returned from the optimizer as a cached value.
	// We want to assure that every named variable creates a copy of
	// another named variable instead of using the cached value itself
	// because it could lead to incorrect optimizations.
	if f.IsIdentified(rightValue) {
		_, isResource := rightValue.Type().(*types.Resource)

		if isResource {
			f.Block().Unidentify(rightValue)
		} else {
			rightValue = f.copy(rightValue, left.Source())
		}
	}

	root := left.Parent

	if isAssign && root.Token.Kind != token.Assign {
		operator := removeAssign(root.Token.Kind)

		operation := f.Append(&ssa.BinaryOp{
			Op:     operator,
			Left:   leftValue,
			Right:  rightValue,
			Source: root.Source(),
		})

		f.Block().Identify(name, operation)
		return nil
	}

	if !isCall {
		structure, isStructType := rightValue.(*ssa.Struct)

		if isStructType {
			for i, field := range types.Unwrap(structure.Typ).(*types.Struct).Fields {
				f.Block().Identify(name+"."+field.Name, structure.Arguments[i])
			}
		}

		f.Block().Identify(name, rightValue)
		return nil
	}

	structure, isStructType := types.Unwrap(rightValue.Type()).(*types.Struct)

	if !isStructType {
		f.Block().Identify(name, rightValue)
		return nil
	}

	composite := &ssa.Struct{
		Typ:       rightValue.Type(),
		Arguments: make(ssa.Arguments, 0, len(structure.Fields)),
		Source:    left.Source(),
	}

	for i, field := range structure.Fields {
		fieldValue := &ssa.FromTuple{
			Tuple:     rightValue,
			Index:     i,
			Structure: composite,
			Source:    left.Source(),
		}

		f.Block().Append(fieldValue)
		f.Block().Identify(name+"."+field.Name, fieldValue)
		composite.Arguments = append(composite.Arguments, fieldValue)
	}

	f.Block().Identify(name, composite)
	return nil
}