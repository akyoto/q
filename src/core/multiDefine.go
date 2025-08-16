package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/expression"
	"git.urbach.dev/cli/q/src/ssa"
)

// multiDefine defines multiple variables at once.
func (f *Function) multiDefine(left *expression.Expression, right *expression.Expression) error {
	value, err := f.evaluateCall(right)

	if err != nil {
		return err
	}

	fn := value.(*ssa.Call).Func
	output := fn.Typ.Output
	count := 0

	for leaf := range left.Leaves() {
		name := leaf.String(f.File.Bytes)
		fromTuple := f.Append(&ssa.FromTuple{Tuple: value, Index: count})
		f.Block().Identify(name, fromTuple)
		count++
	}

	if count != len(output) {
		return errors.New(&DefinitionCountMismatch{Function: fn.String(), Count: count, ExpectedCount: len(output)}, f.File, left.Source().StartPos)
	}

	return nil
}