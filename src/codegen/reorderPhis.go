package codegen

import (
	"slices"

	"git.urbach.dev/cli/q/src/ssa"
)

// reorderPhis makes the order of phis consistent.
func (f *Function) reorderPhis() {
	start := -1

	for _, step := range f.Steps {
		_, isPhi := step.Value.(*ssa.Phi)

		switch {
		case isPhi && start == -1:
			start = step.Index
		case !isPhi && start != -1:
			end := step.Index
			phis := f.Steps[start:end]

			slices.SortStableFunc(phis, func(a *Step, b *Step) int {
				aIndex := f.ValueToStep[a.Value.(*ssa.Phi).Arguments[0]].Index
				bIndex := f.ValueToStep[b.Value.(*ssa.Phi).Arguments[0]].Index
				return aIndex - bIndex
			})

			for i, phi := range phis {
				phi.Index = start + i
			}

			start = -1
		}
	}
}