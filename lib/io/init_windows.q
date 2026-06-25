init() {
	kernel32.SetConsoleCP(utf8)
	kernel32.SetConsoleOutputCP(utf8)
	stdin = kernel32.GetStdHandle(-10)
	stdout = kernel32.GetStdHandle(-11)
}

global {
	stdin uint
	stdout uint
}

const {
	utf8 = 65001
}

extern {
	kernel32 {
		GetStdHandle(device int64) -> (handle uint)
		SetConsoleCP(cp uint)
		SetConsoleOutputCP(cp uint)
	}
}