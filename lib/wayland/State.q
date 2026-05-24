State {
	socket int
	id uint32
	registry uint32
}

newId(state *State) -> uint32 {
	state.id += 1
	return state.id
}