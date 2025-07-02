package core

import (
	"git.urbach.dev/cli/q/src/errors"
)

// CheckDeadCode checks for dead values.
func (f *Function) CheckDeadCode() error {
	for instr := range f.Values {
		if instr.IsConst() && instr.CountUsers() == 0 {
			return errors.New(&UnusedValue{Value: instr.String()}, f.File, instr.Start())
		}
	}

	return nil
}