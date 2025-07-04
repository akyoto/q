import os

init() {
	utf8 := 65001
	kernel32.SetConsoleCP(utf8)
	kernel32.SetConsoleOutputCP(utf8)
	main.main()
	os.exit(0)
}

crash() {
	os.exit(1)
}

extern {
	kernel32 {
		SetConsoleCP(cp uint)
		SetConsoleOutputCP(cp uint)
	}
}