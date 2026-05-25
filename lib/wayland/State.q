State {
	socket int
	id uint32
	wl_compositor uint32
	wl_registry uint32
	wl_shm uint32
	xdg_wm_base uint32
}

newId(state *State) -> uint32 {
	state.id += 1
	return state.id
}