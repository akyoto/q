init() {
	stack := asm.sp + 8 as *uint
	argc = [stack]
	argv = stack + 8
}