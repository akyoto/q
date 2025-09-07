fork() -> int {
	return clone(sigchld, 0, 0, 0, 0)
}

clone(flags uint, stack *any|nil, parent *int|nil, child *int|nil, tls uint) -> int {
	return syscall(_clone, flags, stack, parent, child, tls)
}

const {
	sigchld = 17
}