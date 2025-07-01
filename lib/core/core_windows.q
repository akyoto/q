init() {
	utf8 := 65001
	kernel32.SetConsoleCP(utf8)
	kernel32.SetConsoleOutputCP(utf8)
	main.main()
	exit()
}

exit() {
	kernel32.ExitProcess(0)
}

crash() {
	kernel32.ExitProcess(1)
}

extern {
	kernel32 {
		SetConsoleCP(cp uint)
		SetConsoleOutputCP(cp uint)
		ExitProcess(code uint)
	}
}