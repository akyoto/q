import mem
import process

create(func any) -> (tid int) {
	stack := mem.mmap(0, 4096, mem.read|mem.write, mem.private|mem.anonymous, -1, 0) as *uint64
	stack[512 - 1] = end as uint64
	stack[512 - 2] = func as uint64
	flags := CLONE_THREAD | CLONE_VM | CLONE_FS | CLONE_FILES | CLONE_IO | CLONE_SIGHAND | CLONE_SETTLS
	return process.clone(flags, stack + 4096 - 16, 0, 0, 0)
}

end() {
	syscall(_exit, 0)
}