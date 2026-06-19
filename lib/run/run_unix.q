init() {
	main.main()
	exit(0)
}

crash() {
	exit(1)
}

exit(code uint8) {
	syscall(_exit, code)
}