exit(code int) {
	syscall(0x2000001, code)
}