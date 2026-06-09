fork() -> int {
	return clone(sigchld, 0, 0, 0, 0)
}

clone(flags uint, stack *any|nil, parent *int32|nil, child *int32|nil, tls uint) -> int {
	return syscall(_clone, flags, stack, parent, child, tls)
}

const {
	sigchld = 17
}