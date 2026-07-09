import mem
import process

create(func any) -> (tid int) {
	stack := mem.mmap(0, STACK_SIZE, mem.read|mem.write, mem.private|mem.anonymous, -1, 0)
	tls := stack + STACK_SIZE - 32
	entry := tls - 16
	[entry as *uint64] = func as uint64
	[entry + 8 as *uint64] = end as uint64
	flags := CLONE_THREAD | CLONE_VM | CLONE_FS | CLONE_FILES | CLONE_IO | CLONE_SIGHAND | CLONE_SETTLS
	return process.clone(flags, entry, 0, 0, tls as uint)
}

end() {
	syscall(_exit, 0)
}