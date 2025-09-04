package codegen

// createLiveRanges adds the `value` to the list of live values in its live range.
func (f *Function) createLiveRanges(live *Step) {
	for _, user := range live.Value.Users() {
		use := f.ValueToStep[user]
		f.markAlive(live, use.Block, use, true)
	}
}