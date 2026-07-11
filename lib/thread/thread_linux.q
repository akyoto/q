import mem
import process

create(func ()) -> (tid int) {
	stack := mem.mmap(0, STACK_SIZE, mem.read|mem.write, mem.private|mem.anonymous, -1, 0)
	tls := stack + STACK_SIZE - TLS_SIZE
	args := tls - process.CLONE_ARGS_SIZE as *process.CloneArgs
	args.flags = CLONE_THREAD | CLONE_VM | CLONE_FS | CLONE_FILES | CLONE_IO | CLONE_SIGHAND | CLONE_SETTLS
	args.stack = stack as uint64
	args.stack_size = STACK_SIZE - TLS_SIZE
	args.tls = tls as uint64
	tid := syscall(process._clone3, args, process.CLONE_ARGS_SIZE) as int

	if tid == 0 {
		func()
		syscall(_exit, 0)
	}

	return tid
}