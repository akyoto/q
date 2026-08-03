id() -> uint {
	tid := new(uint)
	libc.pthread_threadid_np(0, tid)
	return [tid]
}

extern {
	libc {
		pthread_threadid_np(thread *any|nil, tid *uint) -> int
	}
}