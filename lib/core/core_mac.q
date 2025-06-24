init() {
	main.main()
	exit()
}

exit() {
	syscall(0x2000001, 0)
}

crash() {
	syscall(0x2000001, 1)
}