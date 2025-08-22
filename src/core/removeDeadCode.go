package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/ssa"
)

// removeDeadCode checks for dead values and also removes the ones that are partially allowed.
func (f *Function) removeDeadCode() error {
	for _, block := range f.Blocks {
		for i, value := range block.Instructions {
			if !value.IsConst() {
				continue
			}

			if len(value.Users()) > 0 {
				continue
			}

			structField, isFieldOfStruct := value.(ssa.StructField)

			if isFieldOfStruct && structField.Struct() != nil {
				partiallyUnused := false

				for _, field := range structField.Struct().Arguments {
					if len(field.Users()) > 0 {
						partiallyUnused = true
						break
					}
				}

				if partiallyUnused {
					block.RemoveAt(i)
					continue
				}
			}

			source := value.(ssa.HasSource)
			return errors.New(&UnusedValue{Value: source.StringFrom(f.File.Bytes)}, f.File, source.Start())
		}

		block.RemoveNilValues()
	}

	return nil
}