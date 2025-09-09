init() {
	main.main()
	exit(0)
}

crash() {
	exit(1)
}

exit(code byte) {
	kernel32.ExitProcess(code as uint)
}

extern {
	kernel32 {
		ExitProcess(code uint)
	}
}