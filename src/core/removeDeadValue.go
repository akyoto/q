package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/ssa"
	"git.urbach.dev/cli/q/src/token"
)

// removeDeadValue checks if the value is dead and removes it.
func (f *Function) removeDeadValue(block *ssa.Block, i int, value ssa.Value, folded map[ssa.Value]struct{}) error {
	if !value.IsPure() {
		return nil
	}

	if len(value.Users()) > 0 {
		return nil
	}

	_, isFolded := folded[value]

	if isFolded {
		block.RemoveAt(i)
		return nil
	}

	phi, isPhi := value.(*ssa.Phi)

	if isPhi {
		if phi.IsPartiallyUndefined() {
			block.RemoveAt(i)
			return nil
		}

		value = phi.FirstDefined()
	}

	structure, isFieldOfStruct := f.valueToStruct[value]

	if isFieldOfStruct {
		partiallyUnused := false

		for _, field := range structure.Arguments {
			if len(field.Users()) > 0 {
				partiallyUnused = true
				break
			}
		}

		if partiallyUnused {
			block.RemoveAt(i)
			return nil
		}
	}

	source := value.(errors.Source)
	return errors.New(&UnusedValue{Value: source.StringFrom(f.File.Bytes)}, f.File, token.NewSource(source.Start(), source.End()))
}