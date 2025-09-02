init() {
	main.main()
	exit(0)
}

crash() {
	exit(1)
}

exit(code int) {
	kernel32.ExitProcess(code)
}

extern {
	kernel32 {
		ExitProcess(code uint)
	}
}