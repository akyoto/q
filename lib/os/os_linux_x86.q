exit(code int) {
	syscall(60, code)
}