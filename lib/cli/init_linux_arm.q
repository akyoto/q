init() {
	stack := asm.sp as *uint
	argc = [stack]
	argv = stack + 8
	envp = argv + argc * 8 + 8
}