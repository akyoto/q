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

	call, isCall := rightValue.(*ssa.Call)

	if !isCall {
		return errors.New(ExpectedFunctionCall, f.File, right.Source())
	}

	fn := call.Func
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
			value := &ssa.Field{
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

		composite := f.makeStructFromTuple(rightValue, fn.Typ.Output[i], structure, name, identifier.Source())

		for _, element := range composite.Arguments {
			element.(*ssa.Field).Index += count
			protected = append(protected, element)
		}

		count += len(composite.Arguments)
		f.Block().Identify(name, composite)
		protected = append(protected, composite)
	}

	return nil
}