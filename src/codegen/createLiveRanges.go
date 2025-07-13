package codegen

import "git.urbach.dev/cli/q/src/ssa"

// createLiveRanges adds the `value` to the list of live values in its live range.
func (f *Function) createLiveRanges(step *step) {
	users := step.Value.Users()

	if len(users) == 0 {
		return
	}

	liveStart := step.Index
	_, isParam := step.Value.(*ssa.Parameter)

	if isParam {
		liveStart = 0
	}

	liveEnd := f.ValueToStep[users[len(users)-1]].Index

	for i := liveStart; i < liveEnd; i++ {
		f.Steps[i].Live = append(f.Steps[i].Live, step)
	}
}