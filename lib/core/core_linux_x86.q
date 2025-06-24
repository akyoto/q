init() {
	main.main()
	exit()
}

exit() {
	syscall(60, 0)
}

crash() {
	syscall(60, 1)
}