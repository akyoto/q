init() {
	stack := asm.sp
	argc = [stack as *uint]
	argv = stack + 8 as **byte
	envp = argv + argc * 8 + 8 as **byte
}