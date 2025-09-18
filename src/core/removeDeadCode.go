package core

import (
	"git.urbach.dev/cli/q/src/ssa"
)

// removeDeadCode checks for dead values and also removes the ones that are partially allowed.
func (f *Function) removeDeadCode(folded map[ssa.Value]struct{}) error {
	for _, block := range f.Blocks {
		for i, value := range block.Instructions {
			err := f.removeDeadValue(block, i, value, folded)

			if err != nil {
				return err
			}
		}

		block.RemoveNilValues()
	}

	return nil
}