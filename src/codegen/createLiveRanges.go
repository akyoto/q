package codegen

// createLiveRanges adds the `value` to the list of live values in its live range.
func (f *Function) createLiveRanges(step *step) {
	for _, user := range step.Value.Users() {
		userStep := f.ValueToStep[user]
		index := userStep.Block.Index(user)
		f.markAlive(step, userStep.Block.Instructions[:index], userStep.Block)
	}
}