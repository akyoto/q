const {
	SIGCHLD = 17
}

fork() -> int {
	args := new(CloneArgs) {
		exit_signal: SIGCHLD
	}

	return clone3(args, 88)
}