fork() -> int {
	pid := syscall(_fork)

	if pid == id() {
		return 0
	}

	return pid
}