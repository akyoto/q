init() {
	// kernel32.SetConsoleCP(cp.utf8)
	// kernel32.SetConsoleOutputCP(cp.utf8)
	main.main()
	exit()
}

exit() {
	// kernel32.ExitProcess(0)
}

crash() {
	// kernel32.ExitProcess(1)
}

// const {
// 	cp {
// 		utf8 65001
// 	}
// }

// extern {
// 	kernel32 {
// 		SetConsoleCP(cp uint)
// 		SetConsoleOutputCP(cp uint)
// 		ExitProcess(code uint)
// 	}
// }