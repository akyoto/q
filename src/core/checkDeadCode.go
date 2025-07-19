package core

import (
	"git.urbach.dev/cli/q/src/errors"
	"git.urbach.dev/cli/q/src/ssa"
)

// checkDeadCode checks for dead values.
func (f *Function) checkDeadCode() error {
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
				used := false

				for _, field := range structField.Struct().Arguments {
					if len(field.Users()) > 0 {
						used = true
						break
					}
				}

				if used {
					block.Instructions[i] = nil
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