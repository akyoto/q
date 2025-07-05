exit(code int) {
	syscall(0x1, code)
}