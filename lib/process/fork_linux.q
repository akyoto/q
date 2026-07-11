const {
	SIGCHLD = 17
}

fork() -> int {
	args := new(CloneArgs) {
		exit_signal: SIGCHLD
	}

	return syscall(_clone3, args, CLONE_ARGS_SIZE)
}