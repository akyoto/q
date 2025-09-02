init() {
	main.main()
	exit(0)
}

crash() {
	exit(1)
}

exit(code int) {
	syscall(_exit, code)
}