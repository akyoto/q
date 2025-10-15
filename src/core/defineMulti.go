package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// defineMulti creates SSA values from expressions and composes structs from their individual fields.
func (f *Function) defineMulti(left *expression.Expression, right *expression.Expression, isAssign bool) error {
	if left.Token.Kind != token.Separator {
		return errors.New(InvalidLeftExpression, f.File, left.Source())
	}

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
		return errors.New(&DefinitionCountMismatch{Function: fn.String(), Count: len(leaves), ExpectedCount: len(fn.Typ.Output)}, f.File, left.Source())
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

		fields := make(ssa.Arguments, 0, len(structure.Fields))

		for _, field := range structure.Fields {
			fieldValue := &ssa.FromTuple{
				Tuple:  rightValue,
				Index:  count,
				Source: identifier.Source(),
			}

			f.Block().Append(fieldValue)
			f.Block().Identify(name+"."+field.Name, fieldValue)
			fields = append(fields, fieldValue)
			protected = append(protected, fieldValue)
			count++
		}

		composite := f.makeStruct(fn.Typ.Output[i], identifier.Source(), fields)
		f.Block().Identify(name, composite)
		protected = append(protected, composite)
	}

	return nil
}