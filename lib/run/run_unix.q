init() {
	stack = asm.sp
	main.main()
	exit(0)
}

crash() {
	exit(1)
}

exit(code byte) {
	syscall(_exit, code)
}

global {
	stack *uint
}