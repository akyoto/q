package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/ssa"
)

// CheckDeadCode checks for dead values.
func (f *Function) CheckDeadCode() error {
	for _, instr := range f.Values {
		if !instr.IsConst() {
			continue
		}

		if len(instr.(ssa.HasLiveness).Users()) > 0 {
			continue
		}

		source := instr.(ssa.HasSource)
		return errors.New(&UnusedValue{Value: string(f.File.Bytes[source.Start():source.End()])}, f.File, source.Start())
	}

	return nil
}