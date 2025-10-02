init() {
	stack := asm.sp as *uint
	argc = [stack]
	argv = stack + 8
}