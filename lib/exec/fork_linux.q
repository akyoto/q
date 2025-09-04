fork() -> int {
	return clone(17, 0, 0, 0, 0)
}

clone(flags uint, stack *any, parent *int, child *int, tls uint) -> int {
	return syscall(_clone, flags, stack, parent, child, tls)
}