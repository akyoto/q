exit(code int) {
	syscall(_exit, code)
}