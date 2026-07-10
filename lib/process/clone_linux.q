clone3(args *CloneArgs, size uint) -> int {
	return syscall(_clone3, args, size)
}

CloneArgs {
	flags uint64
	pidfd uint64
	child_tid uint64
	parent_tid uint64
	exit_signal uint64
	stack uint64
	stack_size uint64
	tls uint64
	set_tid uint64
	set_tid_size uint64
	cgroup uint64
}