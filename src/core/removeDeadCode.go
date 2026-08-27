package core

import (
	"slices"

	"git.urbach.dev/cli/q/src/ssa"
)

// removeDeadCode checks for dead values and also removes the ones that are partially allowed.
func (f *Function) removeDeadCode(folded map[ssa.Value]struct{}) error {
	for {
		count := f.CountValues()

		for _, block := range slices.Backward(f.Blocks) {
			var errors []error

			for i, value := range slices.Backward(block.Instructions) {
				err := f.removeDeadValue(block, i, value, folded)

				if err != nil {
					errors = append(errors, err)
				}
			}

			if len(errors) > 0 {
				return errors[len(errors)-1]
			}

			block.RemoveNilValues()
		}

		if f.CountValues() == count {
			return nil
		}
	}
}