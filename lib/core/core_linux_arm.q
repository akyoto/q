init() {
	main.main()
	exit()
}

exit() {
	syscall(93, 0)
}

crash() {
	syscall(93, 1)
}