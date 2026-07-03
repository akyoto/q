init() {
	stack := asm.sp + 8
	argc = [stack as *uint]
	argv = stack + 8 as **byte
	envp = argv + argc * 8 + 8 as **byte
}