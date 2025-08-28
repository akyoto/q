package codegen

import "git.urbach.dev/cli/q/src/ssa"

// createLiveRanges adds the `value` to the list of live values in its live range.
func (f *Function) createLiveRanges(step *Step) {
	for _, user := range step.Value.Users() {
		userStep := f.ValueToStep[user]
		index := 0
		_, isPhi := user.(*ssa.Phi)

		if !isPhi {
			index = userStep.Block.Index(user)
		}

		f.markAlive(step, userStep.Block.Instructions[:index], userStep.Block, userStep)
	}
}