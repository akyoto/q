fork() -> int {
	return clone(sigchld, 0, 0, 0, 0)
}

clone(flags uint, stack *any, parent *int, child *int, tls uint) -> int {
	return syscall(_clone, flags, stack, parent, child, tls)
}

const {
	sigchld = 17
}