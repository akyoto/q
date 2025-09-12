import exec
import mem

create(func any) -> int {
	stack := mem.mmap(0, 4096, mem.read|mem.write, mem.private|mem.anonymous, -1, 0) as *uint64
	stack[512 - 1] = end as uint64
	stack[512 - 2] = func as uint64
	return exec.clone(vm|fs|files|signalHandlers|parent|thread|ioContext, stack + 4096 - 16, 0, 0, 0)
}

end() {
	syscall(_exit, 0)
}