import mem
import process

create(func any) -> (tid int) {
	stack := mem.mmap(0, STACK_SIZE, mem.read|mem.write, mem.private|mem.anonymous, -1, 0)
	tls := stack + STACK_SIZE - 32
	entry := tls - 32
	[entry as *uint64] = func as uint64
	[entry + 16 as *uint64] = end as uint64
	args := entry - 88 as *process.CloneArgs
	args.flags = CLONE_THREAD | CLONE_VM | CLONE_FS | CLONE_FILES | CLONE_IO | CLONE_SIGHAND | CLONE_SETTLS
	args.stack = stack as uint64
	args.stack_size = STACK_SIZE - 32 - 32
	args.tls = tls as uint64
	return process.clone3(args, 88)
}

end() {
	syscall(_exit, 0)
}