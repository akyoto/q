package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/types"
)

// defineMulti creates SSA values from expressions and composes structs from their individual fields.
func (f *Function) defineMulti(left *expression.Expression, right *expression.Expression, isAssign bool) error {
	rightValue, err := f.evaluateRight(right)

	if err != nil {
		return err
	}

	fn := rightValue.(*ssa.Call).Func
	leaves := make([]*expression.Expression, 0, 2)

	for leaf := range left.Leaves() {
		leaves = append(leaves, leaf)
	}

	if len(leaves) != len(fn.Typ.Output) {
		return errors.New(&DefinitionCountMismatch{Function: fn.String(), Count: len(leaves), ExpectedCount: len(fn.Typ.Output)}, f.File, left.Source().StartPos)
	}

	protected := make([]ssa.Value, 0, len(leaves))
	count := 0

	for i, identifier := range leaves {
		name := identifier.String(f.File.Bytes)

		if name == "_" {
			count++
			continue
		}

		_, err := f.validateLeft(identifier, right, name, fn.Typ.Output[i], isAssign)

		if err != nil {
			return err
		}

		structure, isStructType := types.Unwrap(fn.Typ.Output[i]).(*types.Struct)

		if !isStructType {
			value := &ssa.FromTuple{
				Tuple:  rightValue,
				Index:  count,
				Source: identifier.Source(),
			}

			f.Block().Append(value)
			f.Block().Identify(name, value)

			if value.Type() == types.Error {
				f.Block().Protect(value, protected)
				protected = nil
			} else {
				protected = append(protected, value)
			}

			count++
			continue
		}

		composite := &ssa.Struct{
			Typ:       fn.Typ.Output[i],
			Arguments: make(ssa.Arguments, 0, len(structure.Fields)),
			Source:    identifier.Source(),
		}

		for _, field := range structure.Fields {
			fieldValue := &ssa.FromTuple{
				Tuple:     rightValue,
				Index:     count,
				Structure: composite,
				Source:    identifier.Source(),
			}

			f.Block().Append(fieldValue)
			f.Block().Identify(name+"."+field.Name, fieldValue)
			composite.Arguments = append(composite.Arguments, fieldValue)
			protected = append(protected, fieldValue)
			count++
		}

		f.Block().Identify(name, composite)
		protected = append(protected, composite)
	}

	return nil
}