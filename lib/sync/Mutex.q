lock(state *uint32) {
	loop {
		if cas(state, 0, 1) {
			return
		}

		wait(state, 1)
	}
}

unlock(state *uint32) {
	[state] = 0
	wake(state, 1)
}