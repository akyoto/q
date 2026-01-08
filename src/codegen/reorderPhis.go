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
			start = int(step.Index)
		case !isPhi && start != -1:
			end := step.Index
			phis := f.Steps[start:end]

			slices.SortStableFunc(phis, func(a *Step, b *Step) int {
				aPhi := a.Value.(*ssa.Phi)
				bPhi := b.Value.(*ssa.Phi)

				for i := range aPhi.Arguments {
					aValue := aPhi.Arguments[i]
					bValue := bPhi.Arguments[i]
					aIndex := f.ValueToStep[aValue].Index
					bIndex := f.ValueToStep[bValue].Index
					cmp := aIndex - bIndex

					if cmp != 0 {
						return int(cmp)
					}
				}

				return 0
			})

			for i, phi := range phis {
				phi.Index = Index(start + i)
			}

			start = -1
		}
	}
}