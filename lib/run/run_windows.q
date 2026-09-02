import io

init() {
	main.main()
	exit(0)
}

crash(message string) {
	io.writeLine(message)
	exit(1)
}

exit(code uint8) {
	kernel32.ExitProcess(code as uint)
}

extern {
	kernel32 {
		ExitProcess(code uint)
	}
}