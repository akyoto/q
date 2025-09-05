fork() -> int {
	pid := syscall(_fork)

	if pid == getpid() {
		return 0
	}

	return pid
}

getpid() -> int {
	return syscall(_getpid)
}