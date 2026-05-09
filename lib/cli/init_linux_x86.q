init() {
	stack := asm.sp + 8 as *uint
	argc = [stack]
	argv = stack + 8
	envp = argv + argc * 8 + 8
}