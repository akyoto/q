exit(code int) {
	syscall(93, code)
}