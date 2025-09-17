init() {
	kernel32.SetConsoleCP(utf8)
	kernel32.SetConsoleOutputCP(utf8)
}

extern {
	kernel32 {
		SetConsoleCP(cp uint)
		SetConsoleOutputCP(cp uint)
	}
}

const {
	utf8 = 65001
}