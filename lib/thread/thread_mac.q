create(func ()) -> (tid int) {
	thread := new(uint64)
	System.pthread_create(thread, 0, func, 0)
	return 0
}

extern {
	System {
		pthread_create(thread *uint64, attr *any|nil, func any, arg *any|nil) -> int
	}
}