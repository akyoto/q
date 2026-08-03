id() -> uint {
	tid := new(uint)
	System.pthread_threadid_np(0, tid)
	return [tid]
}

extern {
	System {
		pthread_threadid_np(thread *any|nil, tid *uint) -> int
	}
}