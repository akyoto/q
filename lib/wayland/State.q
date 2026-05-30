State {
	socket int
	id uint32

	wl_compositor uint32
	wl_registry uint32
	wl_shm uint32
	wl_shm_pool uint32
	xdg_wm_base uint32

	wl_compositor_name uint32
	wl_shm_name uint32
	xdg_wm_base_name uint32

	shm_fd !int
	shm_size uint32
	shm_data *byte
}

newId(state *State) -> uint32 {
	state.id += 1
	return state.id
}