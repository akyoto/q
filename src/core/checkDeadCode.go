package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/ssa"
)

// checkDeadCode checks for dead values.
func (f *Function) checkDeadCode() error {
	for _, value := range f.Values {
		if !value.IsConst() {
			continue
		}

		if len(value.Users()) > 0 {
			continue
		}

		source := value.(ssa.HasSource)
		return errors.New(&UnusedValue{Value: source.StringFrom(f.File.Bytes)}, f.File, source.Start())
	}

	for _, value := range f.Identifiers {
		structure, isStruct := value.(*ssa.Struct)

		if isStruct {
			used := false

			for _, field := range structure.Arguments {
				if len(field.Users()) > 0 {
					used = true
					break
				}
			}

			if used {
				continue
			}

			source := value.(ssa.HasSource)
			return errors.New(&UnusedValue{Value: source.StringFrom(f.File.Bytes)}, f.File, source.Start())
		}

		if len(value.Users()) > 0 {
			continue
		}

		source := value.(ssa.HasSource)
		return errors.New(&UnusedValue{Value: source.StringFrom(f.File.Bytes)}, f.File, source.Start())
	}

	return nil
}