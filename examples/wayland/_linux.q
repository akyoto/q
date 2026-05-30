import c
import fs
import mem

createShm(state *wayland.State) -> error {
	path := c.string("wl_shm")
	fd, err := fs.memfd_create(path.ptr, 0)
	delete(path)

	if err != 0 {
		return err
	}

	state.shm_fd = fd
	width := 1280
	height := 720
	state.shm_size = width * height * 4
	err := fs.ftruncate(state.shm_fd, state.shm_size as uint64)

	if err != 0 {
		return err
	}

	state.shm_data = mem.mmap(0, state.shm_size as uint64, mem.read|mem.write, mem.shared, state.shm_fd, 0)
	return 0
}

deleteShm(state *wayland.State) {
	if state.shm_data != 0 {
		mem.munmap(state.shm_data, state.shm_size as uint64)
	}

	if state.shm_fd != 0 {
		fs.close(state.shm_fd)
	}
}